package query

import (
  . "github.com/woojiahao/gmail-filter-organiser/pkg/logging"
  "google.golang.org/api/gmail/v1"
)

func (q *Query) listLabels() []*gmail.Label {
  labels, err := q.Srv.Users.Labels.List(user).Do()
  IfError(err, "Unable to retrieve labels from user account")
  return labels.Labels
}

func (q *Query) getLabel(label string) *gmail.Label {
  labels := q.listLabels()
  for _, l := range labels {
    if l.Id == label {
      return l
    }
  }

  return nil
}