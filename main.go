package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	// "syscall"
	"test/app/user"
	"test/app/zoos"
	"test/routers"
	"time"
)

func main() {
	fmt.Printf("http://localhost:8080 \n")

	// 加载多个APP的路由配置
	routers.Include(shop.Routers, user.Routers)
	// 初始化路由
	r := routers.Init()

	// 定义服务器
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	// 利用 goroutine 启动监听
	go func() {
		// server.ListenAndServe() 监听
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	// quit 信道是同步信道，若没有信号进来，处于阻塞状态
	// 反之，则执行后续代码
	<-quit
	log.Println("关闭服务中 ...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// 调用 server.Shutdown() 完成优雅停止
	// 调用时传递了一个上下文对象，对象中定义了超时时间
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("服务已关闭")

}
