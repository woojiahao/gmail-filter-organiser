package logging

import (
  "log"
)

func IfError(err error, info string) {
  if err != nil {
    log.Fatalf("%s:%v", info, err)
  }
}
