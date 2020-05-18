package logging

import (
  "fmt"
  "log"
)

func IfError(err error, info string) {
  if err != nil {
    log.Fatalf("%s:%v", info, err)
  }
}

func Info(info string, args ...interface{}) {
  output := fmt.Sprintf(info, args)
  log.Printf("%s\n", output)
}
