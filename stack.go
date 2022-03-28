package lumberjack

import (
  "runtime"
  "strconv"
  "strings"
)

var goroot = runtime.GOROOT()

func formatFrame(f runtime.Frame) string {
  var prefix string
  if i := strings.LastIndexByte(f.Function, '/'); i != -1 {
    prefix = f.Function[:i]
  }

  suffix := f.File
  if i := strings.LastIndexByte(suffix, '/'); i != -1 {
    suffix = suffix[strings.LastIndexByte(suffix[:i], '/')+1:]
  }

  line := ":" + strconv.Itoa(f.Line)
  if prefix == "" {
    return suffix + line
  }

  return prefix + "/" + suffix + line
}

func captureStackTrace(skip int) string {
  var pcs [512]uintptr

  n := runtime.Callers(skip+1, pcs[:])
  frames := runtime.CallersFrames(pcs[:n])
  frame, more := frames.Next()

  var (
    b strings.Builder
    first = true
  )

  b.WriteRune('[')

  for more {
    frame, more = frames.Next()
    if strings.HasPrefix(frame.File, goroot) {
      continue
    }

    b.WriteString(formatFrame(frame))

    if first {
      first = false
    } else {
      b.WriteRune(' ')
    }
  }

  b.WriteRune(']')

  return b.String()
}