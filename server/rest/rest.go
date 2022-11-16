package rest

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/fangaoxs/synk/env"
	c "github.com/fangaoxs/synk/server/container"
	"github.com/fangaoxs/synk/server/ws"
	"github.com/gin-gonic/gin"
)

func StartGin(FS embed.FS) {
	hub := ws.NewHub()
	go hub.Run()

	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	// access static files
	staticFiles, _ := fs.Sub(FS, "frontend/dist")
	router.StaticFS("/static", http.FS(staticFiles))

	// api v1 routers
	v1 := router.Group("/api/v1")
	{
		v1.POST("/texts", c.TextsContainer())                 // 上传文本
		v1.POST("/files", c.FilesContainer())                 // 上传文件
		v1.GET("/addresses", c.AddressesContainer())          // 获取局域网可用ip
		v1.GET("/qrcode", c.QrcodeContainer())                // 对文本生成二维码
		v1.GET("/downloads/:filename", c.DownloadContainer()) // 根据文件名下载文件
	}

	router.GET("/ws", c.WsController(hub))

	// default router
	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		if strings.HasPrefix(path, "/static/") {
			reader, err := staticFiles.Open("index.html")
			if err != nil {
				log.Fatal(err)
			}
			defer reader.Close()
			stat, err := reader.Stat()
			if err != nil {
				log.Fatal(err)
			}
			c.DataFromReader(http.StatusOK, stat.Size(), "text/html;charset=utf-8", reader, nil)
		} else {
			c.Status(http.StatusNotFound)
		}
	})

	router.Run(fmt.Sprintf(":%s", env.HttpListenPort))
}
