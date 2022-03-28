package lumberjack

import (
  "io"
  "testing"
)

func BenchmarkLogger(b *testing.B) {
  l := NewLogger(Config{
    Writer: io.Discard,
    Level:  LevelDebug,
  })

  b.ReportAllocs()
  b.ResetTimer()

  for i := 0; i < b.N; i++ {
    l.
      Debug("this is a message with a realistic length").
      Str("data", "none of your business").
      Int("int", 10).
      Str("this", "is").
      Str("very", "fast").
      Uint16("uint16", 50).
      Str("fast", "as").
      Str("fuck", "boi").
      Uint64("uint64", 444).
      Str("last", "field").
      Flush()
  }
}