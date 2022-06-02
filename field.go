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

func acquireField(start, end int) *field {
  f := fieldPool.Get().(*field)
  f.start = start
  f.end = end
  return f
}

func releaseField(f *field) {
  fieldPool.Put(f)
}