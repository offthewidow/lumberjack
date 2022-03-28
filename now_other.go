//go:build !windows

package log

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