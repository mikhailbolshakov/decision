package kit

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
)

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel = "panic"
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel = "fatal"
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel = "error"
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel = "warning"
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel = "info"
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel = "debug"
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel = "trace"

	FormatterText = "plain"
	FormatterJson = "json"
)

// ErrorHook allows specifying a hook for all logged errors
type ErrorHook interface {
	Error(err error)
}

// LogConfig represents logging configuration
type LogConfig struct {
	Level   string // Level logging level
	Format  string // Format (plain, json)
	Context bool   // Context if true, context params are part of logging
	Service bool   // Service if true, service params are part of logging
}

type Logger struct {
	Logrus *logrus.Logger
	Cfg    *LogConfig
	hook   ErrorHook
}

func InitLogger(cfg *LogConfig) *Logger {
	logger := &Logger{
		Cfg:    cfg,
		Logrus: logrus.New(),
	}
	logger.Init(cfg)
	return logger
}

func (l *Logger) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	}
}

func (l *Logger) Fire(entry *logrus.Entry) error {
	if l.hook != nil {
		if err, ok := entry.Data[logrus.ErrorKey]; ok {
			l.hook.Error(err.(error))
		}
	}
	return nil
}

func (l *Logger) SetErrorHook(h ErrorHook) {
	l.hook = h
}

func (l *Logger) Init(cfg *LogConfig) {
	l.Cfg = cfg
	l.SetLevel(cfg.Level)
	l.Logrus.SetOutput(os.Stdout)

	// add hook to trace errors
	l.Logrus.AddHook(l)

	if cfg.Format == FormatterJson {
		l.Logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat:   "2006-01-02T15:04:05.000-0700",
			DisableHTMLEscape: true,
		})
	} else {
		l.Logrus.SetFormatter(&Formatter{
			FixedFields:            []string{"call.svc", "call.pr", "call.cmp", "call.mth", "call.node"},
			TimestampFormat:        "2006-01-02T15:04:05.000-0700",
			HideKeysForFixedFields: true,
			NoColors:               true,
			NoFieldsColors:         true,
			NoFieldsSpace:          true,
			ShowFullLevel:          true,
		})
	}
}

func (l *Logger) GetLogger() *logrus.Logger {
	return l.Logrus
}

func (l *Logger) SetLevel(level string) {
	lv, err := logrus.ParseLevel(level)
	if err != nil {
		panic(err)
	}
	l.Logrus.SetLevel(lv)
}

type CLoggerFunc func() CLogger

// CLogger provides structured logging abilities
// !!!! Not thread safe. Don't share one CLogger instance through multiple goroutines
type CLogger interface {
	// C - adds request context to log
	//
	// don't put context when logging error, as it makes sense a context of where error happens rather than a context of where error log is invoked
	// otherwise, context will be logged twice
	C(ctx context.Context) CLogger
	// F - adds fields to log
	F(fields KV) CLogger
	// E - adds error to log
	E(err error) CLogger
	// St - adds stack to log (if err is already set)
	St() CLogger
	// Cmp - adds component
	Cmp(c string) CLogger
	// Mth - adds method
	Mth(m string) CLogger
	// Pr - adds protocol
	Pr(m string) CLogger
	// Srv - adds service code
	Srv(s string) CLogger
	// Node - adds service instance code
	Node(n string) CLogger
	Inf(args ...interface{}) CLogger
	InfF(format string, args ...interface{}) CLogger
	Err(args ...interface{}) CLogger
	ErrF(format string, args ...interface{}) CLogger
	Dbg(args ...interface{}) CLogger
	DbgF(format string, args ...interface{}) CLogger
	Trc(args ...interface{}) CLogger
	TrcF(format string, args ...interface{}) CLogger
	// TrcObj marshals all args only if loglevel = Trace, otherwise bypass
	// Note that only Exported fields of objects are logged (due to nature of json.Marshal)
	TrcObj(format string, args ...interface{}) CLogger
	Warn(args ...interface{}) CLogger
	WarnF(format string, args ...interface{}) CLogger
	Fatal(args ...interface{}) CLogger
	FatalF(format string, args ...interface{}) CLogger
	Clone() CLogger
	Printf(string, ...interface{})
	PrintfErr(f string, args ...interface{})
	// Write writer implementation
	Write(p []byte) (n int, err error)
}

func L(logger *Logger) CLogger {
	return &clogger{
		logger: logger,
		lre:    logrus.NewEntry(logger.Logrus),
	}
}

type clogger struct {
	logger *Logger
	lre    *logrus.Entry
	err    error
}

// Clone always use Clone when pass CLogger between goroutines
func (cl *clogger) Clone() CLogger {
	entry := logrus.NewEntry(cl.logger.Logrus)
	if len(cl.lre.Data) > 0 {
		marshaled, _ := json.Marshal(cl.lre.Data)
		_ = json.Unmarshal(marshaled, &entry.Data)
	}
	clone := &clogger{
		lre: entry,
		err: cl.err,
	}
	return clone
}

func (cl *clogger) C(ctx context.Context) CLogger {
	if !cl.logger.Cfg.Context {
		return cl
	}
	if r, ok := Request(ctx); ok && r != nil {
		ff := make(KV)
		if ct := r.GetCaller(); ct != "" {
			ff["ctx.cl"] = ct
		}
		if rid := r.GetRequestId(); rid != "" {
			ff["ctx.rid"] = rid
		}
		if un := r.GetUsername(); un != "" {
			ff["ctx.un"] = un
		}
		if sid := r.GetSessionId(); sid != "" {
			ff["ctx.sid"] = sid
		}
		cl.F(ff)
	}
	return cl
}

func (cl *clogger) F(fields KV) CLogger {
	cl.lre = cl.lre.WithFields(map[string]interface{}(fields))
	return cl
}

func (cl *clogger) E(err error) CLogger {

	// if err is AppErr, log error code as a separate field
	if appErr, ok := IsAppErr(err); ok {

		// put code / message as fields
		cl.lre = cl.lre.WithField("err-code", appErr.Code())
		cl.lre = cl.lre.WithField("error", appErr.Message())
		cl.lre = cl.lre.WithField("error-type", appErr.Type())

		// pass fields from err to log
		for k, v := range appErr.Fields() {
			cl.lre = cl.lre.WithField(k, v)
		}

	}
	cl.err = err
	cl.lre = cl.lre.WithError(err)
	return cl
}

func (cl *clogger) St() CLogger {
	if cl.err != nil {
		// if err is AppErr take stack from error itself, otherwise build stack right here
		if appErr, ok := IsAppErr(cl.err); ok {
			cl.lre = cl.lre.WithField("err-stack", appErr.WithStack())
		} else {
			buf := make([]byte, 1<<16)
			s := runtime.Stack(buf, false)
			cl.lre = cl.lre.WithField("err-stack", string(buf[0:s]))
		}
	}
	return cl
}

func (cl *clogger) Srv(s string) CLogger {
	if !cl.logger.Cfg.Service {
		return cl
	}
	cl.lre = cl.lre.WithField("call.svc", s)
	return cl
}

func (cl *clogger) Node(n string) CLogger {
	if !cl.logger.Cfg.Service {
		return cl
	}
	cl.lre = cl.lre.WithField("call.node", n)
	return cl
}

func (cl *clogger) Cmp(c string) CLogger {
	cl.lre = cl.lre.WithField("call.cmp", c)
	return cl
}

func (cl *clogger) Pr(c string) CLogger {
	cl.lre = cl.lre.WithField("call.pr", c)
	return cl
}

func (cl *clogger) Mth(m string) CLogger {
	cl.lre = cl.lre.WithField("call.mth", m)
	return cl
}

func (cl *clogger) Err(args ...interface{}) CLogger {
	cl.lre.Errorln(args...)
	return cl
}

func (cl *clogger) ErrF(format string, args ...interface{}) CLogger {
	cl.lre.Errorf(format, args...)
	return cl
}

func (cl *clogger) Inf(args ...interface{}) CLogger {
	cl.lre.Infoln(args...)
	return cl
}

func (cl *clogger) InfF(format string, args ...interface{}) CLogger {
	cl.lre.Infof(format, args...)
	return cl
}

func (cl *clogger) Warn(args ...interface{}) CLogger {
	cl.lre.Warningln(args...)
	return cl
}

func (cl *clogger) WarnF(format string, args ...interface{}) CLogger {
	cl.lre.Warningf(format, args...)
	return cl
}

func (cl *clogger) Dbg(args ...interface{}) CLogger {
	cl.lre.Debugln(args...)
	return cl
}

func (cl *clogger) DbgF(format string, args ...interface{}) CLogger {
	cl.lre.Debugf(format, args...)
	return cl
}

func (cl *clogger) Trc(args ...interface{}) CLogger {
	cl.lre.Traceln(args...)
	return cl
}

func (cl *clogger) TrcF(format string, args ...interface{}) CLogger {
	cl.lre.Tracef(format, args...)
	return cl
}

func (cl *clogger) TrcObj(format string, args ...interface{}) CLogger {
	if cl.logger.Cfg.Level == TraceLevel {
		var argsJs []interface{}
		for _, a := range args {
			if a != nil {
				js, _ := json.Marshal(a)
				argsJs = append(argsJs, string(js))
			}
		}
		return cl.TrcF(format, argsJs...)
	}
	return cl
}

func (cl *clogger) Fatal(args ...interface{}) CLogger {
	cl.lre.Fatalln(args...)
	return cl
}

func (cl *clogger) FatalF(format string, args ...interface{}) CLogger {
	cl.lre.Fatalf(format, args...)
	return cl
}

func (cl *clogger) Printf(f string, args ...interface{}) {
	cl.DbgF(f, args...)
}

func (cl *clogger) PrintfErr(f string, args ...interface{}) {
	cl.ErrF(f, args...)
}

func (cl *clogger) Write(p []byte) (n int, err error) {
	cl.Trc(string(p))
	return len(p), nil
}
