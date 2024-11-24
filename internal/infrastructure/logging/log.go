package logging

import (
	"github.com/evalphobia/logrus_sentry"
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"path/filepath"
)

var log = logrus.New()

func InitLogger() {
	// Set log format to JSON for structured logging
	log.SetFormatter(&logrus.JSONFormatter{})

	// Setup file logging with lumberjack.Logger
	logFile := &lumberjack.Logger{
		Filename:   filepath.Join("storage", "Logs", "app.log"),
		MaxSize:    10, // Megabytes
		MaxBackups: 3,
		MaxAge:     28, // Days
		Compress:   true,
	}

	// Open the file for writing. Create it if it doesn't exist.
	file, err := os.OpenFile(logFile.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.WithError(err).Error("Failed to open log file")
		return
	}
	defer file.Close()

	// Use a MultiWriter to write to both stdout and file
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	// Only log the info severity or above
	log.SetLevel(logrus.InfoLevel)

	// Initialize Sentry asynchronously
	go initSentry()
}

func initSentry() {
	defer func() {
		if r := recover(); r != nil {
			log.Warn("Recovered from panic in Sentry initialization: ", r)
		}
	}()

	err := sentry.Init(sentry.ClientOptions{
		Dsn: viper.GetString("SENTRY_DSN"),
		// Adjust other Sentry options as needed
	})
	if err != nil {
		log.Warn("Sentry initialization failed: ", err)
		return
	}

	// Add Sentry hook
	hook, err := logrus_sentry.NewSentryHook(viper.GetString("SENTRY_DSN"), []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	})
	if err != nil {
		log.Warn("Failed to create Sentry hook: ", err)
		return
	}

	log.Hooks.Add(hook)
	log.Info("Sentry initialized successfully")
}
