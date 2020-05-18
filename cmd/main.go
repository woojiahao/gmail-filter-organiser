package main

import (
  "github.com/woojiahao/gmail-filter-organiser/pkg/auth"
  . "github.com/woojiahao/gmail-filter-organiser/pkg/logging"
)

func main() {
  srv := auth.Connect()
  filters, err := srv.Users.Settings.Filters.List("me").Do()
  IfError(err, "Unable to retrieve filters of current user")
  for _, filter := range filters.Filter {
    Info(filter.Id)
  }
}
