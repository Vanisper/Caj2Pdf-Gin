package main

import (
	"Caj2PdfServer/routes"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

var serverPort = 8080
var distPath = "./templates/dist/"

// dist目录下需要开放访问的静态资源的文件夹
var assets = []string{"assets", "output", "lib", "test"}

func main() {
	router := gin.Default()
	// 直接一个星号的话 这个dist文件夹下面禁止存在子文件夹  否则会报错||而下面的写法则规避了这种问题
	router.LoadHTMLGlob(distPath + "/*.html")
	router.GET("/", func(context *gin.Context) {
		//time.Sleep(5 * time.Second)
		//context.String(http.StatusOK, "Welcome Gin Server")
		context.HTML(http.StatusOK, "index.html", gin.H{})
	})
	router.GET("/api/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	routes.Load(router)
	// 循环挂载静态资源路由
	for _, item := range assets {
		router.Static("/"+item, distPath+item)
	}
	//_ = browser.OpenURL("http://127.0.0.1:" + strconv.Itoa(serverPort))
	//err := router.Run(":" + strconv.Itoa(serverPort))
	//if err != nil {
	//	return
	//} // 监听并在 0.0.0.0:8080 上启动服务
	log.Println("运行于:", "http://127.0.0.1:"+strconv.Itoa(serverPort))
	// 优雅地关闭server: 当退出这个软件时或者命令行执行ctrl+c终止进程时，会执行以下的退出提示
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(serverPort),
		Handler: router,
	}
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("服务优雅关闭 ...")
	// 如果服务器无法正常关闭，则会在5秒钟后强制关闭
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("服务优雅关闭遇到了状况: ", err)
	}
	log.Println("服务已优雅退出")
}
