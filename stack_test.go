package log

import "testing"

func BenchmarkCaptureStackTrace(b *testing.B) {
  b.ReportAllocs()
  b.ResetTimer()

  for i := 0; i < b.N; i++ {
    _ = captureStackTrace(0)
  }
}