package container

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

// QrcodeContainer 根据文本内容生成二维码
func QrcodeContainer() gin.HandlerFunc {
	return func(c *gin.Context) {
		content := c.Query("content")
		if content == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		// 调用接口，根据文本生成二维码
		png, err := qrcode.Encode(content, qrcode.Medium, 256)
		if err != nil {
			log.Fatal(err)
		}
		c.Data(http.StatusOK, "image/png", png)
	}
}
