package config

import (
	"YellowPepper-FundsTransfers/pkg/misc/environment"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var singletonLogger *zap.Logger

func StopMonitoring() error {

	// Stop Zap
	if err := singletonLogger.Sync(); err != nil {
		return err
	}

	// Stop Sentry
	sentry.Flush(2 * time.Second)
	return nil
}

func InitMonitoring(environment environment.Environment) {

	// Setup & Init Sentry
	sentryOptions := sentry.ClientOptions{
		Dsn:         environment.GetValue("SENTRY_DSN").AsString(),
		Environment: environment.GetValue("SENTRY_ENVIRONMENT").AsString(),
		Release:     environment.GetValue("SENTRY_RELEASE").AsString(),
		Debug:       true,
	}

	if err := sentry.Init(sentryOptions); err != nil {
		log.Fatalln(fmt.Sprintf("sentry.Init: %s", err))
	}

	// Setup Zap
	loggerConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	// Init Zap with Sentry Hooks for error level logs
	var err error
	if singletonLogger, err = loggerConfig.Build(zap.Hooks(sentryHook)); err != nil {
		log.Println(err)
	}
	zap.ReplaceGlobals(singletonLogger)
}

func sentryHook(entry zapcore.Entry) error {
	if entry.Level == zapcore.ErrorLevel {
		sentry.CaptureMessage(fmt.Sprintf("%s, Line No: %d :: %s", entry.Caller.File, entry.Caller.Line, entry.Message))
		sentry.CaptureException(errors.New(entry.Message))
	}
	return nil
}
