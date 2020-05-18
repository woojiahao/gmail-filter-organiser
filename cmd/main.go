package main

import (
  "github.com/woojiahao/gmail-filter-organiser/pkg/auth"
  "github.com/woojiahao/gmail-filter-organiser/pkg/query"
)

func main() {
  srv := auth.Connect()
  q := query.Query{Srv: srv}
  q.Organise()
}
