//go:build windows

package log

import "time"

func now() time.Time {
  return time.Now()
}