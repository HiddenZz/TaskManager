package l

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"time"
)

func InitLogging() {
	handler := slog.Handler(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(slog.New(handler))
}

func EMsg(format string, args ...any) {
	logger := slog.Default()
	if !logger.Enabled(context.Background(), slog.LevelInfo) {
		return
	}
	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])
	r := slog.NewRecord(time.Now(), slog.LevelError, fmt.Sprintf("Error: %s", fmt.Sprintf(format, args...)), pcs[0])
	_ = logger.Handler().Handle(context.Background(), r)
}

func E(err error) {
	EMsg(err.Error())
}
