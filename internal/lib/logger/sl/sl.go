package sl

import (
	"log/slog"
)

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func Attr(s1, s2 string) slog.Attr {
	return slog.Attr{
		Key:   s1,
		Value: slog.StringValue(s2),
	}
}
