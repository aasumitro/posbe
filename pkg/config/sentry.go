package config

import (
	"github.com/getsentry/sentry-go"
	"log"
)

func (cfg Config) InitCrashReporting() {
	log.Println("Trying to initialize crash reporting handler . . . .")

	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              cfg.CrashReportDsnUrl,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	}); err != nil {
		log.Fatalf("Sentry initialization failed: %v\n", err)
	}

	log.Printf("Crash reporting set to %s . . . .", cfg.CrashReportDriver)
}
