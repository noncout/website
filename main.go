package main

import (
	"context"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"website/models"
	"website/pkg/logging"
	"website/pkg/setting"
	"website/routers"
)

func Setup() {
	setting.Setup()
	models.Setup()
	logging.Setup()
}

func main() {
	Setup()
	logging.Info("------------------Start Website------------------")
	StartWithHttpServer()
}

//func StartWithEndless () {
//Setup()
//
//	endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
//	endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
//	endless.DefaultMaxHeaderBytes = 1 << 20
//	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
//
//	server := endless.NewServer(endPoint, routers.InitRouter())
//	server.BeforeBegin = func(add string) {
//		log.Printf("Actual pid is %d", syscall.Getpid())
//	}
//
//	err := server.ListenAndServe()
//	if err != nil {
//		log.Printf("Server err: %v", err)
//	}
//}

func StartWithHttpServer() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	// 启动并运行协程
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	// 创建channel
	quit := make(chan os.Signal)
	// 接收信号
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// 关闭资源
	defer cancel()

	// shotdown 会让server立即退出,
	// 然后shutdown阻塞在这里等待所有的连接都断开或者超时
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
