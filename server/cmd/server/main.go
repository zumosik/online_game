package main

import (
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path"
	"server/internal/server"
	"time"

	"gitlab.com/greyxor/slogor"
)

const (
	infoLvl  = "info"
	debugLvl = "debug"
	errorLvl = "error"
)

func main() {
	var logLevel, logFile string
	flag.StringVar(&logLevel, "log-lvl", "debug", "log level (debug, info, error), debug will enable pretty slog")
	flag.StringVar(&logFile, "log-file", "", "path to log file, default: logs only in stdout")
	flag.Parse()

	logger := setupLogger(logLevel, logFile)

	srv := server.New(&server.Config{
		Addr:        ":8080",
		Logger:      logger,
		MaxReadSize: 1024,
		PathToSave:  "saves/default.111",
	})

	srv.MustStart()

}

func setupLogger(env, filePath string) *slog.Logger {
	var log *slog.Logger

	var output io.Writer

	if filePath != "" && env != debugLvl { // cant be used in Debug lvl because of color codes

		f, err := os.Create(path.Join(filePath, fmt.Sprintf("log%s.log", time.Now().Format("2006-01-02_15-04-05"))))
		if err != nil {
			panic(err)
		}
		output = io.MultiWriter(f, os.Stdout)
	} else {
		output = os.Stdout
	}

	switch env {
	case debugLvl:
		log = slog.New(slogor.NewHandler(output, slogor.Options{
			TimeFormat: time.Stamp,
			Level:      slog.LevelDebug,
			ShowSource: false,
		}))
	case infoLvl:
		log = slog.New(
			slog.NewTextHandler(output, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	case errorLvl:
		log = slog.New(
			slog.NewJSONHandler(output, &slog.HandlerOptions{Level: slog.LevelError}),
		)
	}

	return log
}
