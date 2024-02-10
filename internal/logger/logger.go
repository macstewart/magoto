package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func init() {
	writer := zerolog.NewConsoleWriter()
	writer.Out = os.Stderr
	writer.PartsExclude = []string{"time"}
	//skip frame count 3 gets back to the actual calling method (zerolog << <this class> << caller)
	// logger = zerolog.New(writer).With().CallerWithSkipFrameCount(3).Logger()
	logger = zerolog.New(writer)
}

// use custom level enum to reduce zerolog coupling
func SetGlobalLevel(level Level) {
	switch level {
	case TRACE:
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case DEBUG:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case INFO:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case WARN:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case ERROR:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}
}

func Trace(msg string, fields ...interface{}) {
	if len(fields) > 0 {
		logger.Trace().Msgf(msg, fields...)
	} else {
		logger.Trace().Msg(msg)
	}
}

func Debug(msg string, fields ...interface{}) {
	if len(fields) > 0 {
		logger.Debug().Msgf(msg, fields...)
	} else {
		logger.Debug().Msg(msg)
	}
}

func Info(msg string, fields ...interface{}) {
	if len(fields) > 0 {
		logger.Info().Msgf(msg, fields...)
	} else {
		logger.Info().Msg(msg)
	}
}

func Warn(msg string, fields ...interface{}) {
	if len(fields) > 0 {
		logger.Warn().Msgf(msg, fields...)
	} else {
		logger.Warn().Msg(msg)
	}
}

func WarnErr(msg string, err error, fields ...interface{}) {
	if len(fields) > 0 {
		logger.Warn().Err(err).Msgf(msg, fields...)
	} else {
		logger.Warn().Err(err).Msg(msg)
	}
}

func Error(msg string, fields ...interface{}) {
	if len(fields) > 0 {
		logger.Error().Msgf(msg, fields...)
	} else {
		logger.Error().Msg(msg)
	}
}

func ErrorErr(msg string, err error, fields ...interface{}) {
	if len(fields) > 0 {
		logger.Error().Err(err).Msgf(msg, fields...)
	} else {
		logger.Error().Err(err).Msg(msg)
	}
}

func Fatal(msg string, fields ...interface{}) {
	if len(fields) > 0 {
		logger.Fatal().Msgf(msg, fields...)
	} else {
		logger.Fatal().Msg(msg)
	}
}

func FatalErr(msg string, err error, fields ...interface{}) {
	if len(fields) > 0 {
		logger.Fatal().Err(err).Msgf(msg, fields...)
	} else {
		logger.Fatal().Err(err).Msg(msg)
	}
}
