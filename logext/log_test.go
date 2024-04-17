package logext

import (
	"log/slog"
	"testing"
)

func TestLog(t *testing.T) {

	rotateFile := CreateRotateFileWriter("test")
	slogger := slog.New(SlogHandlerNew(rotateFile, slog.LevelInfo, true, 2))
	slog.SetDefault(slogger)

	slog.Info("ddd", "d", "1")

}
