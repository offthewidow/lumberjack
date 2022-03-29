package lumberjack

import (
  "runtime"
  "strconv"
  "strings"
)

const maxStackSize = 32

var goroot = runtime.GOROOT()

func captureStackTrace(skip int) string {
  var pcs [maxStackSize]uintptr

  n := runtime.Callers(skip+1, pcs[:])
  frames := runtime.CallersFrames(pcs[:n])

  var (
    b strings.Builder
    first = true
  )

  b.Grow(1 + 24 * (n - 1) + 1) // allocate 24 initial bytes per frame + 2 more bytes for the square brackets
  b.WriteRune('[')

  for {
    frame, more := frames.Next()
    if !more {
      break
    }

    file := frame.File
    if strings.HasPrefix(file, goroot) {
      continue
    }

    if first {
      first = false
    } else {
      b.WriteRune(' ')
    }

    if i := strings.LastIndexByte(frame.Function, '/'); i != -1 {
      b.WriteString(frame.Function[:i])
      b.WriteRune('/')
    }

    if i := strings.LastIndexByte(file, '/'); i != -1 {
      file = file[strings.LastIndexByte(file[:i], '/')+1:]
    }

    b.WriteString(file)
    b.WriteRune(':')
    b.WriteString(strconv.Itoa(frame.Line))
  }

  b.WriteRune(']')

  return b.String()
}