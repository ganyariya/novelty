package logger

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	debugLogger *log.Logger
	logFile     *os.File
)

func InitLogger() error {
	// logsディレクトリを作成
	if err := os.MkdirAll("logs", 0755); err != nil {
		return err
	}
	
	// ログファイルを開く
	logPath := filepath.Join("logs", "debug.log")
	var err error
	logFile, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	
	// デバッグロガーを設定（ファイルのみに出力）
	debugLogger = log.New(logFile, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
	
	return nil
}

func CloseLogger() {
	if logFile != nil {
		logFile.Close()
	}
}

func Debug(format string, args ...interface{}) {
	if debugLogger != nil {
		debugLogger.Printf(format, args...)
	}
}

func Info(format string, args ...interface{}) {
	if debugLogger != nil {
		debugLogger.Printf("[INFO] "+format, args...)
	}
}

func Error(format string, args ...interface{}) {
	if debugLogger != nil {
		debugLogger.Printf("[ERROR] "+format, args...)
	}
}

// 開発時のみ標準出力にも出力する場合
func DebugToStdout(format string, args ...interface{}) {
	if debugLogger != nil {
		// ファイルとstdoutの両方に出力
		multiWriter := io.MultiWriter(logFile, os.Stdout)
		tempLogger := log.New(multiWriter, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
		tempLogger.Printf(format, args...)
	}
}