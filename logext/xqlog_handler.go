package logext

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

const TimeLayout = "2006-01-02 15:04:05.000"

type XqLogHandler struct {
	Level  slog.Leveler
	Stdout bool
	out    io.Writer
}

func CreateRotateFileWriter(name string) *rotatelogs.RotateLogs {
	logf, err := rotatelogs.New(
		fmt.Sprintf("logs/%s.%%Y%%m%%d.log", name),
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if nil != err {
		panic(err)
	}
	return logf
}

func SlogHandlerNew(out io.Writer, level slog.Leveler, stdout bool) *XqLogHandler {
	h := &XqLogHandler{Level: level, out: out, Stdout: stdout}
	return h
}

func (h *XqLogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.Level.Level()
}

func (h *XqLogHandler) WithGroup(name string) slog.Handler {
	panic("未实现 WithGroup")
}

func (h *XqLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	panic("未实现 WithAttrs")
}

func (h *XqLogHandler) Handle(ctx context.Context, r slog.Record) error {
	sb := strings.Builder{}
	if !r.Time.IsZero() {
		sb.WriteString(fmt.Sprintf("%-23s ", r.Time.Format(TimeLayout)))
	}
	sb.WriteString(fmt.Sprintf("%-5s ", r.Level.String()))

	callerStr, funcStr := caller(r)

	if len(callerStr) > 43 {
		callerStr = callerStr[0:20] + "..." + callerStr[len(callerStr)-20:]
	}
	sb.WriteString(fmt.Sprintf("%-43s ", callerStr))

	if len(funcStr) > 16 {
		funcStr = funcStr[0:6] + "..." + funcStr[len(funcStr)-7:]
	}
	sb.WriteString(fmt.Sprintf("%-16s ", funcStr))

	sb.WriteString(r.Message + "\n")

	printData := []byte(sb.String())
	_, err := h.out.Write(printData)
	if h.Stdout {
		os.Stdout.Write(printData)
	}
	return err
}

func caller(r slog.Record) (caller, funcStr string) {
	if r.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		ec, _ := fs.Next()
		idx := strings.LastIndexByte(ec.File, '/')
		pathStr := ec.File
		if idx != -1 {
			idx = strings.LastIndexByte(ec.File[:idx], '/')
			if idx != -1 {
				pathStr = ec.File[idx+1:]
			}
		}
		funcArr := strings.Split(ec.Func.Name(), ".")
		funcName := "unknown"
		if len(funcArr) > 1 {
			funcName = funcArr[1]
		}
		return fmt.Sprintf("%s:%-5d", pathStr, int64(ec.Line)), funcName
	} else {
		return "unknown:0", "unknown"
	}
}
