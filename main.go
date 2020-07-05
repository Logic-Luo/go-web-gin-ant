package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-web-gin-ant/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

//ListenAddr 监听服务端口
const ListenAddr = ":80"

func main() {
	router := gin.Default()

	router.GET("/index", func(c *gin.Context) {
		//根据完整文件名渲染模板，并传递参数
		c.String(http.StatusOK, "some string")
	})

	router.POST("/api/login", func(c *gin.Context) {
		c.Writer.Write( model.LoginResp{
			Status:           "ok",
			CurrentAuthority: "admin",
		})
		c.JSON(-1, model.LoginResp{
			Status:           "ok",
			CurrentAuthority: "admin",
		})
	})

	router.StaticFile("/favicon.ico", "static/favicon.png")
	dir, _ := ioutil.ReadDir("static")
	for _, file := range dir {
		name := file.Name()
		splits := strings.Split(name, "/")
		router.StaticFile(splits[len(splits)-1], "static/"+name)
	}
	staticRouter := router.Group("/static", func(ginCtx *gin.Context) {
		//如果请求的是js或css文件，response header里设置一年的cache时间
		ext := filepath.Ext(ginCtx.Request.URL.Path)
		if ext == ".css" || ext == ".js" {
			ginCtx.Header("Cache-Control", "max-age=315360000, public")
		}
		ginCtx.Next()
	})
	staticRouter.Static("/static", "static")

	router.NoRoute(func(ginCtx *gin.Context) {
		indexContent, err := ioutil.ReadFile("static/index.html")
		if err != nil {
			ginCtx.Data(http.StatusOK, "text/plain", []byte(err.Error()))
		} else {
			ginCtx.Data(http.StatusOK, "text/html; charset=utf-8", indexContent)
		}
	})

	srv := &http.Server{Addr: ListenAddr, Handler: router}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Listen:", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")
	time.Sleep(5 * time.Second)
}
