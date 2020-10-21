package logext

import (
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"time"
)

/*GetRotationLogWriter 获取循环日志记录器
 *
 * fileName 日志文件名称 如：./logs/main.log
 * maxTime 最大日志记录时间
 * rotationTime 日志轮换周期
 */
func GetRotationLogWriter(filename string, maxTime time.Duration, rotationTime time.Duration) io.Writer {
	// 生成rotateLogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	hook, err := rotateLogs.New(filename+".%Y%m%d_%H%M",
		//rotateLogs.WithLinkName(filename),
		rotateLogs.WithMaxAge(maxTime),
		rotateLogs.WithRotationTime(rotationTime),
	)
	if err != nil {
		panic(err)
	}
	return hook
}
