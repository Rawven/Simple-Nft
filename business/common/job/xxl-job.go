package job

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	xxl "github.com/xxl-job/xxl-job-executor-go"
	"log"
)

var exec xxl.Executor

type XxlTaskFunc struct {
	Name string
	Task xxl.TaskFunc
}

func initTask() {
	exec := xxl.NewExecutor(
		xxl.ServerAddr(viper.GetString("xxl.server")),
		xxl.AccessToken(viper.GetString("xxl.token")),
		xxl.RegistryKey(viper.GetString("xxl.key")),
		xxl.SetLogger(&logger{}),
	)
	exec.Init()
	exec.Use(customMiddleware)
	//设置日志查看handler
	exec.LogHandler(customLogHandle)

}

func RegAndRunTask(tasks []XxlTaskFunc) error {
	initTask()
	for _, task := range tasks {
		exec.RegTask("task."+task.Name, task.Task)
	}
	return exec.Run()
}

// 自定义日志处理器
func customLogHandle(req *xxl.LogReq) *xxl.LogRes {
	return &xxl.LogRes{Code: xxl.SuccessCode, Msg: "", Content: xxl.LogResContent{
		FromLineNum: req.FromLineNum,
		ToLineNum:   2,
		LogContent:  "自定义日志handler",
		IsEnd:       true,
	}}
}

// xxl.Logger接口实现
type logger struct{}

func (l *logger) Info(format string, a ...interface{}) {
	log.Println(fmt.Sprintf("xxl-job - "+format, a...))
}

func (l *logger) Error(format string, a ...interface{}) {
	log.Println(fmt.Sprintf("xxl-job - "+format, a...))
}

// 自定义中间件
func customMiddleware(tf xxl.TaskFunc) xxl.TaskFunc {
	return func(cxt context.Context, param *xxl.RunReq) string {
		log.Println("I am a middleware start")
		res := tf(cxt, param)
		log.Println("I am a middleware end")
		return res
	}
}
