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

func getField(start, end int) *field {
  f := fieldPool.Get().(*field)
  f.start = start
  f.end = end
  return f
}

func putField(f *field) {
  fieldPool.Put(f)
}