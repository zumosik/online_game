package utils

import "log/slog"

func WrapErr(err error) slog.Attr {
	return slog.String("error", err.Error())
}

func Wrap(name, msg string) slog.Attr {
	if name == "" {
		name = "info"
	}
	return slog.String(name, msg)
}
