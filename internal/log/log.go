package log

import (
	"fmt"

	"github.com/AnnonaOrg/osenv"
	"github.com/sirupsen/logrus"
)

func init() {
	logLevelStr := osenv.GetLogLevel()
	logLevel, err := logrus.ParseLevel(logLevelStr)
	if err != nil {
		fmt.Printf("日志级别(%s)解析失败❌: %v\n", logLevelStr, err)
		return
	}
	logrus.SetLevel(logLevel)
	// 日志输出文件位置
	// logrus.SetReportCaller(true)
	// Set the formatter to include file and line information
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05 MST",
		// CallerPrettyfier: func(f *runtime.Frame) (string, string) {
		// 	return "", fmt.Sprintf("%s:%d", f.File, f.Line)
		// },
	})

	fmt.Println("日志配置完成✅")
	return
}
