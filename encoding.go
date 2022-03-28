package log

import "strings"

func shouldQuote(s string) bool {
  return s == "" || strings.IndexRune(s, ' ') != -1
}

func appendString(dst []byte, s string, quote bool) []byte {
  if quote {
    dst = append(dst, '"')
  }
  dst = append(dst, s...)
  if quote {
    dst = append(dst, '"')
  }
  return dst
}

func appendKey(dst []byte, k string, quote bool, pretty bool) []byte {
  dst = append(dst, ' ')
  if pretty {
    return append(append(appendString(append(dst, "\x1b[90m"...), k, quote), '='), "\x1b[0m"...)
  }
  return append(appendString(dst, k, quote), '=')
}