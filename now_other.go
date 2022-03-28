//go:build !windows

package lumberjack

import (
  "syscall"
  "time"
)

func now() time.Time {
  var t syscall.Timeval
  if syscall.Gettimeofday(&t) != nil {
    return time.Now()
  }
  return time.Unix(t.Unix())
}