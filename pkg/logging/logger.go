package logging

import (
	"fmt"
	"time"

	"github.com/mohar9h/go-graphql-ws-api/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type zapLogger struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
}

var logLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

func NewZapLogger(cfg *config.Config) *zapLogger {
	logger := &zapLogger{cfg: cfg}
	logger.Init()
	return logger
}

func (l *zapLogger) getLevel() zapcore.Level {
	level, exists := logLevelMap[l.cfg.Logger.Level]
	if !exists {
		return zapcore.InfoLevel
	}
	return level
}
func (l *zapLogger) Init() {
	fileName := fmt.Sprintf("%s%s.%s", l.cfg.Logger.FilePath, time.Now().Format("2006-01-02"), "log")

	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    1,
		MaxAge:     20,
		LocalTime:  true,
		MaxBackups: 5,
		Compress:   true,
	})

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		writeSyncer,
		l.getLevel(),
	)

	l.logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
}

func (l *zapLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]any) {
	params := prepareLogKeys(extra, cat, sub)
	l.logger.Debugw(msg, params...)
}

func (l *zapLogger) Debugf(template string, args ...any) {
	l.logger.Debugf(template, args...)
}

func (l *zapLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]any) {
	params := prepareLogKeys(extra, cat, sub)
	l.logger.Infow(msg, params...)
}

func (l *zapLogger) Infof(template string, args ...any) {
	l.logger.Infof(template, args...)
}

func (l *zapLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]any) {
	params := prepareLogKeys(extra, cat, sub)
	l.logger.Warnw(msg, params...)
}

func (l *zapLogger) Warnf(template string, args ...any) {
	l.logger.Warnf(template, args...)
}

func (l *zapLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]any) {
	params := prepareLogKeys(extra, cat, sub)
	l.logger.Errorw(msg, params...)
}

func (l *zapLogger) Errorf(template string, args ...any) {
	l.logger.Errorf(template, args...)
}

func (l *zapLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]any) {
	params := prepareLogKeys(extra, cat, sub)
	l.logger.Fatalw(msg, params...)
}

func (l *zapLogger) Fatalf(template string, args ...any) {
	l.logger.Fatalf(template, args...)
}

func prepareLogKeys(extra map[ExtraKey]any, cat Category, sub SubCategory) []any {
	if extra == nil {
		extra = make(map[ExtraKey]any, 0)
	}

	extra["Category"] = cat
	extra["SubCategory"] = sub
	params := mapToZapParams(extra)
	return params
}
