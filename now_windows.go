//go:build windows

package lumberjack

import "time"

func now() time.Time {
  return time.Now()
}