package config

import (
	"fmt"
	"log"

	"github.com/getsentry/sentry-go"
)

func SentryConnection() Option {
	return func(cfg *Config) {
		if cfg.AppDebug {
			log.Println("Sentry is disabled in localhost . . . .")
			return
		}
		log.Println("Trying to initialize crash reporting handler . . . .")
		crashReportingCollectorName := "sentry"
		sampleRatio := 0.25
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:                cfg.SentryDsnURL,
			EnableTracing:      true,
			TracesSampleRate:   1.0,
			ProfilesSampleRate: sampleRatio,
			SampleRate:         sampleRatio,
			Environment: func() string {
				if cfg.AppDebug {
					return "development"
				}
				return "production"
			}(),
			Release: fmt.Sprintf("[%s]%s",
				cfg.AppName, cfg.AppVersion),
			Debug: true,
		}); err != nil {
			log.Fatalf("Sentry initialization failed: %v\n", err)
		}
		log.Printf("Crash reporting set to %s . . . .",
			crashReportingCollectorName)
	}
}
