// Handles the authentication and client creation for the Gmail library
package auth

import (
  "context"
  "encoding/json"
  "fmt"
  . "github.com/woojiahao/gmail-filter-organiser/pkg/logging"
  "golang.org/x/oauth2"
  "golang.org/x/oauth2/google"
  "google.golang.org/api/gmail/v1"
  "google.golang.org/api/option"
  "io/ioutil"
  . "log"
  "os"
)

const (
  credentialsFilename = "./credentials.json"
  tokenFilename       = "./token.json"
  webTokenState       = "state-token"
)

func getToken(config *oauth2.Config) *oauth2.Token {
  token, err := getTokenFromFile()
  if err != nil {
    // If token.json does not exist in project root, load the token from the web and request authentication
    token = getTokenFromWeb(config)
    saveToken(token)
  }

  return token
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
  authURL := config.AuthCodeURL(webTokenState, oauth2.AccessTypeOffline)
  fmt.Printf("Go to the following link in your browser, then type the authorization code: \n%v\n", authURL)

  var authCode string
  _, err := fmt.Scan(&authCode)
  IfError(err, "Unable to read authorization code")

  token, err := config.Exchange(context.TODO(), authCode)
  IfError(err, "Unable to retrieve token from web")

  return token
}

func getTokenFromFile() (*oauth2.Token, error) {
  file, err := os.Open(tokenFilename)
  if err != nil {
    return nil, err
  }
  defer func() {
    _ = file.Close()
  }()

  token := &oauth2.Token{}
  err = json.NewDecoder(file).Decode(token)
  return token, err
}

func saveToken(token *oauth2.Token) {
  Printf("Saving token to %s\n", tokenFilename)
  file, err := os.OpenFile(tokenFilename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
  IfError(err, "Unable to cache oauth token")
  defer func() {
    _ = file.Close()
  }()
  _ = json.NewEncoder(file).Encode(token)
}

func Connect() *gmail.Service {
  buf, err := ioutil.ReadFile(credentialsFilename)
  IfError(
    err,
    fmt.Sprintf(
      "Unable to read credentials file. Ensure that %s is stored in the root directory of your program",
      credentialsFilename,
    ),
  )

  scopes := []string{
    gmail.GmailReadonlyScope,
    gmail.GmailLabelsScope,
    gmail.GmailSettingsBasicScope,
  }
  config, err := google.ConfigFromJSON(buf, scopes...)
  IfError(err, "Unable to parse client secret file to config")
  token := getToken(config)

  cxt := context.Background()
  srv, err := gmail.NewService(cxt, option.WithTokenSource(config.TokenSource(cxt, token)))
  IfError(err, "Unable to retrieve Gmail client")

  return srv
}
