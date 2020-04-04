/**
* @Author: zy
* @Date: 2020/04/04 15:00
 */
package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"io/ioutil"
	"os"
	"path"
	_ "path"
	"runtime"
	"strings"
	"time"

	_ "github.com/lestrrat-go/file-rotatelogs" //日志切割
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
}

var writer io.Writer
var err error

// maxAge: 最长保留时间; ratationTime: 日志文件切割周期; level: 日志级别,等级高于该级别才记录;fileFlag: 是否记录到文件; consoleFlag: 是否打印到控制台
func ConfigLogger(logPath string, logFileName string, maxAge time.Duration, ratationTime time.Duration,
	level logrus.Level, fileFlag bool, consoleFlag bool) error {
	Logger.Level = level
	Logger.Formatter = &logrus.TextFormatter{
		ForceColors:      true,
		QuoteEmptyFields: true,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05"}
	if consoleFlag {
		Logger.Out = os.Stdout
	} else {
		Logger.Out = ioutil.Discard
	}
	if !fileFlag {
		return nil
	}
	baseLogpath := path.Join(logPath, logFileName)
	writer, err = rotatelogs.New(
		baseLogpath+"-%Y%m%d%H%M.log",
		rotatelogs.WithLinkName(baseLogpath),      // 生成软链接，指向新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(ratationTime), // 日志切割时间间隔
	)
	if err != nil {
		Logger.Errorf("config local file system logger error: %v", err)
		return err
	}

	Logger.AddHook(&ContextHook{})
	return nil
}

type ContextHook struct{}

func (hook ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

/**
logrus的一个很致命的问题就是没有提供文件名和行号.
解决问题的思路都是对的：通过标准库的runtime模块获取运行时信息，并从中提取文件名，行号和调用函数名。
标准库runtime模块的Caller(skip int)函数可以返回当前goroutine调用栈中的文件名，行号，函数信息等，
参数skip表示表示返回的栈帧的层次，0表示runtime.Caller的调用着。返回值包括响应栈帧层次的pc(程序计数器)，
文件名和行号信息。为了提高效率，我们先通过跟踪调用栈发现，从runtime.Caller()的调用者开始，
到记录日志的生成代码之间，大概有8到11层左右，所有我们在hook中循环第8到11层调用栈应该可以找到日志记录的生产代码。


示例中：使用runtime.Caller()依次循环调用栈的第7~11层，过滤掉sirupsen包内容，
那么第一个非siupsenr包就认为是我们的生产代码了，并返回pc以便通过runtime.FuncForPC()获取函数名称。
然后将文件名、行号和函数名组装为source字段塞到logrus.Entry中即可

*/
func (hook ContextHook) getCallerInfo() (string, string, int) {
	var (
		shortPath string
		funcName  string
	)
	for i := 3; i < 15; i++ {
		pc, fullPath, line, ok := runtime.Caller(i)
		if !ok {
			fmt.Println("error: error during runtime.Caller")
			continue
		} else {
			lastS := strings.LastIndex(fullPath, "/")
			if lastS < 0 {
				lastS = strings.LastIndex(fullPath, "\\")
			}

			shortPath = fullPath[lastS+1:]
			funcName = runtime.FuncForPC(pc).Name()
			if strings.HasPrefix(funcName, "/") {
				funcName = funcName[len("/"):]
			}
			index := strings.LastIndex(funcName, ".")
			if index > 0 {
				funcName = funcName[index+1:]
			}
			if !strings.Contains(strings.ToLower(fullPath), "github.com/sirupsen/logrus") {
				return shortPath, funcName, line
				break
			}
		}
	}
	return "", "", 0
}

func (hook ContextHook) Fire(entry *logrus.Entry) error {
	shortPath, funcName, callLine := hook.getCallerInfo()
	if shortPath != "" && callLine != 0 {
		// entry.Data["caller"] = fmt.Sprintf("%s(%s):%d", shortPath, funcName, callLine)
		entry.Message = fmt.Sprintf("%s(%s):L%d", shortPath, funcName, callLine) + " " + entry.Message
	}
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}
	writer.Write([]byte(line))
	return nil
}
