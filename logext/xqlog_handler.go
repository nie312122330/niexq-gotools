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
	"github.com/timandy/routine"
)

const TimeLayout = "2006-01-02 15:04:05.000"

var LogTracdIdThreadLocal = routine.NewInheritableThreadLocal[string]()

type XqLogHandler struct {
	Level      slog.Leveler
	PrintMehod int // 0-不打印 ，1-详情,2-仅方法名称
	Stdout     bool
	out        io.Writer
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

func SlogHandlerNew(out io.Writer, level slog.Leveler, stdout bool, printMethod int) *XqLogHandler {
	h := &XqLogHandler{Level: level, out: out, Stdout: stdout, PrintMehod: printMethod}
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
	sb.WriteString(fmt.Sprintf("%s ", LogTracdIdThreadLocal.Get()))

	callerStr, funcStr := h.caller(r)
	sb.WriteString(fmt.Sprintf("%s ", callerStr))

	if h.PrintMehod > 0 {
		sb.WriteString(fmt.Sprintf("%s ", funcStr))
	}

	sb.WriteString(r.Message + " ")

	r.Attrs(func(a slog.Attr) bool {
		sb.WriteString(a.String())
		return true
	})
	aa := context.Background().Value("aaa")
	fmt.Println(aa)

	sb.WriteString("\n")

	printData := []byte(sb.String())
	_, err := h.out.Write(printData)
	if h.Stdout {
		os.Stdout.Write(printData)
	}
	return err
}

func (h *XqLogHandler) caller(r slog.Record) (caller, funcStr string) {
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
		funcName := ""
		if h.PrintMehod > 0 {
			if h.PrintMehod == 1 {
				funcName = ec.Func.Name()
			} else {
				funcName = ec.Func.Name()
				funcName = funcName[strings.LastIndex(funcName, ".")+1:]
			}
		}
		return fmt.Sprintf("%s:%d", pathStr, int64(ec.Line)), funcName
	} else {
		return "unknown:0", "unknown"
	}
}
