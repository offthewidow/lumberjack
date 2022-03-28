package lumberjack

import "fmt"

const (
  LevelFatal Level = iota
  LevelError
  LevelWarn
  LevelInfo
  LevelDebug
)

type Level uint8

func formatLevel(lvl Level, pretty bool) string {
  switch lvl {
    case LevelFatal:
      if pretty {
        return "\x1b[35mFTL\x1b[0m"
      }
      return "FTL"
    case LevelError:
      if pretty {
        return "\x1b[31mERR\x1b[0m"
      }
      return "ERR"
    case LevelWarn:
      if pretty {
        return "\x1b[33mWRN\x1b[0m"
      }
      return "WRN"
    case LevelInfo:
      if pretty {
        return "\x1b[32mINF\x1b[0m"
      }
      return "INF"
    case LevelDebug:
      if pretty {
        return "\x1b[34mDBG\x1b[0m"
      }
      return "DBG"
    default:
      panic(fmt.Sprintf("unknown log level %d", lvl))
  }
}