// Logger implementation go.uber.org/zap
package logger

import (
	"fmt"

	"go.uber.org/zap"
)

// Logger is an implementation of the go.uber.org/zap logger.
type logger struct {
	zap *zap.Logger
}

// New builds a Logger
func New() *logger {
	conf := zap.NewDevelopmentConfig()
	conf.Encoding = "console"
	z, _ := conf.Build()
	return &logger{
		zap: z,
	}
}

// Set prodaction
func (l *logger) SetProd() error {
	conf := zap.NewProductionConfig()
	conf.Encoding = "console"
	z, err := conf.Build()
	if err != nil {
		return fmt.Errorf("Logger.SetProd() error: %v", err)
	}
	l.zap = z
	return nil
}

// Infoln uses fmt.Sprintln to construct and log a message.
func (l *logger) Info(args ...any) {
	l.zap.Sugar().Infoln(args...)
}

// Errorln uses fmt.Sprintln to construct and log a message.
func (l *logger) Error(args ...any) {
	l.zap.Sugar().Errorln(args...)
}

// Warnln uses fmt.Sprintln to construct and log a message.
func (l *logger) Warn(args ...any) {
	l.zap.Sugar().Warnln(args...)
}

// Debugln uses fmt.Sprintln to construct and log a message.
func (l *logger) Debug(args ...any) {
	l.zap.Sugar().Debugln(args...)
	// log.Println(args...)
}

// Fatalln uses fmt.Sprintln to construct and log a message, then calls os.Exit.
func (l *logger) Fatal(args ...any) {
	l.zap.Sugar().Fatalln(args...)
}
