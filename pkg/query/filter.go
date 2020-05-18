package query

import (
  "fmt"
  . "github.com/woojiahao/gmail-filter-organiser/pkg/logging"
  "google.golang.org/api/gmail/v1"
  "log"
  "strings"
)

func (q *Query) listFilters() []*gmail.Filter {
  filters, err := q.Srv.Users.Settings.Filters.List(user).Do()
  IfError(err, "Unable to retrieve filters from user account")
  return filters.Filter
}

func (q *Query) deleteFilter(id string) {
  err := q.Srv.Users.Settings.Filters.Delete(user, id).Do()
  IfError(err, fmt.Sprintf("Unable to delete filter %s", id))
}

func (q *Query) createFilter(from string, label string) *gmail.Filter {
  filter, err := q.Srv.Users.Settings.Filters.Create(user, &gmail.Filter{
    Action:   &gmail.FilterAction{AddLabelIds: []string{label}, RemoveLabelIds: []string{"INBOX"}},
    Criteria: &gmail.FilterCriteria{From: from},
  }).Do()
  IfError(err, fmt.Sprintf("Unable to create filter"))
  return filter
}

func hasLabel(labels []string, label string) bool {
  for _, labelToRemove := range labels {
    if labelToRemove == label {
      return true
    }
  }
  return false
}

// Label to FROM
func (q *Query) collectFilters(filters []*gmail.Filter) ([]string, map[*gmail.Label]string) {
  const filterSeparator = " OR "

  filterGrouping := make(map[string][]string)
  filtersToClean := make([]string, 0)
  skipCount := 0
  for _, filter := range filters {
    isRemovingFromInbox := !hasLabel(filter.Action.RemoveLabelIds, "INBOX")
    hasFromCriteria := filter.Criteria.From != ""

    if !isRemovingFromInbox && !hasFromCriteria {
      skipCount++
      continue
    }

    filtersToClean = append(filtersToClean, filter.Id)

    for _, label := range filter.Action.AddLabelIds {
      if filterGrouping[label] == nil {
        filterGrouping[label] = make([]string, 0)
      }
      splitFrom := strings.Split(filter.Criteria.From, filterSeparator)
      filterGrouping[label] = append(filterGrouping[label], splitFrom...)
    }
  }

  log.Printf("Skipped %d filters", skipCount)

  collectedFilters := make(map[*gmail.Label]string)
  for label, filters := range filterGrouping {
    collectedFilters[q.getLabel(label)] = strings.Join(filters, filterSeparator)
  }

  return filtersToClean, collectedFilters
}

func (q *Query) Organise() {
  filters := q.listFilters()

  log.Printf("Total number of filters %d", len(filters))
  filtersToClean, collectedFilters := q.collectFilters(filters)

  log.Printf("Size of collected filters %d", len(collectedFilters))
  for label, filter := range collectedFilters {
    log.Printf("Label: %s assigned to %s", label.Name, filter)
  }

  log.Printf("Clearing repeated filters")
  for _, filter := range filtersToClean {
    q.deleteFilter(filter)
  }

  log.Printf("Creating new filters")
  for label, filter := range collectedFilters {
    _ = q.createFilter(filter, label.Id)
  }
}
