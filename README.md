# gmail-filter-organiser

CLI tool to organise your filters in Gmail

## Running this tool

Ensure that Go is installed. If it isn't, visit [this link](https://golang.org/doc/install) to get 
it setup.

Ensure that you setup the Gmail API with your own API console. Follow step 1 [here](https://developers.google.com/gmail/api/quickstart/go) 
and make sure that you copy the `credentials.json` file to the project root.

Clone the repository

```bash
$ git clone https://github.com/woojiahao/gmail-filter-organiser
$ cd gmail-filter-organiser/
```

Run the tool

```bash
$ go run cmd/main.go
```

If this is the first time you are running this tool (or if the token has expired), the tool will
prompt you to visit the site to authenticate your application. Once it has been authenticated, you
will receive a token that you can copy over to the tool and press "Enter".

## Motivations to make this tool

I have been using Gmail's labels and filters system to organise my inbox for a while now and one 
issue I have faced is having too many filters being created that assign a group of emails to the
same label. This is because Gmail's web interface is severely limited in allowing users to 
create and manage their filters for their accounts. For instance, a filter can only assign a 
search criteria to a single label at a time. Therefore, trying to assign a single email to more
than one label will require multiple filters. Additionally, two filters may have different search 
criteria to the same label.

The goal of this tool is to clean these repeated filters to clean up any clutter and make 
organising filters and labels in Gmail a lot easier.

## Use case

The use case of this CLI is very minimal as it is meant to suit my own use of filters, which
revolves around three key criterion when creating the filter:

1. Skip inbox - remove the label "INBOX"
2. Apply the label 
3. Apply filter to all existing elements

In doing so, we do not add any additional conditionals to the filter (this includes TO, SUBJECT, 
HAS THE WORDS, DOESN'T HAVE, SIZE, ATTACHMENT, etc).

There may be plans to add more but for the current use case, it is not supported.
