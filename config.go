package log

import "io"

var configZeroValue Config

type Config struct{
  Writer io.Writer
  Level  Level
  Pretty bool
}