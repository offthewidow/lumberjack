package log

import "time"

type Logger struct{
  cfg Config
  ctx *entry
}

func NewLogger(cfg Config) *Logger {
  l := &Logger{
    cfg: cfg,
    ctx: newEntry(),
  }
  l.ctx.logger = l
  return l
}

func NewNopLogger() *Logger {
  return NewLogger(configZeroValue)
}

func NotNil(l *Logger) *Logger {
  if l == nil {
    return NewNopLogger()
  }
  return l
}

func (l *Logger) Context() *entry {
  return l.ctx
}

func (l *Logger) Fork() *Logger {
  fork := NewLogger(l.cfg)

  if len(l.ctx.buf) != 0 {
    ctx := fork.ctx
    for k, f := range l.ctx.fields {
      ctx.fields[k] = &field{ f.start, f.end }
    }
    ctx.buf = append(ctx.buf, l.ctx.buf...)
  }

  return fork
}

func (l *Logger) Log(lvl Level, msg string) *entry {
  cfg := l.cfg
  if cfg.Writer == nil || lvl > cfg.Level {
    return nil
  }

  e := entryPool.Get().(*entry)
  e.logger = l
  e.buf = append(append(now().AppendFormat(append(e.buf, "\x1b[90m"...), time.RFC3339), "\x1b[0m "...), formatLevel(lvl, cfg.Pretty)...)

  e.str("msg", false, msg, shouldQuote(msg))

  if len(l.ctx.buf) != 0 {
    size := len(e.buf)
    for k, f := range l.ctx.fields {
      e.fields[k] = &field{ f.start + size, f.end + size }
    }
    e.buf = append(e.buf, l.ctx.buf...)
  }

  return e
}

func (l *Logger) Fatal(msg string) *entry {
  e := l.Log(LevelFatal, msg)
  e.exit = true
  return e
}

func (l *Logger) Error(msg string) *entry {
  return l.Log(LevelError, msg)
}

func (l *Logger) Warn(msg string) *entry {
  return l.Log(LevelWarn, msg)
}

func (l *Logger) Info(msg string) *entry {
  return l.Log(LevelInfo, msg)
}

func (l *Logger) Debug(msg string) *entry {
  return l.Log(LevelDebug, msg)
}