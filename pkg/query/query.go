package query

import "google.golang.org/api/gmail/v1"

const user = "me"

type Query struct {
  Srv *gmail.Service
}
