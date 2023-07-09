package log

import (
	"io"
	stdlog "log"
	"net/http"
	"os"
)

// 创建日志类型指针
var log *stdlog.Logger

// 创建自定义string对象
type fileLog string

// Write 方法, 将 fileLog 转化为 io.Writer 类型
func (fl fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

// Run 初始化日志指针 log, 将写入文件地址 destination 传入
func Run(destination string) {
	log = stdlog.New(fileLog(destination), "[go] - ", stdlog.LstdFlags)
}

// RegisterHandlers 服务注册Handler
func RegisterHandlers() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			msg, err := io.ReadAll(r.Body)
			if err != nil || len(msg) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			write(string(msg))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}

// write 日志写入方法
func write(message string) {
	log.Printf("%v\n", message)
}
