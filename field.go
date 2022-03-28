package lumberjack

import "sync"

var fieldPool = sync.Pool{
  New: func() any {
    return new(field)
  },
}

type field struct{
  start, end int
}