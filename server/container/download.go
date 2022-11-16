package container

import (
	"net/http"
	"path/filepath"

	"github.com/fangaoxs/synk/utils"
	"github.com/gin-gonic/gin"
)

// DownloadContainer 下载 uploads 目录中对应 filename 的文件
func DownloadContainer() gin.HandlerFunc {
	return func(c *gin.Context) {
		filename := c.Param("filename")
		if filename == "" {
			c.Status(http.StatusNotFound)
			return
		}

		// 下载的文件为 uploads/filename
		file := filepath.Join(utils.GetUploadsDir(), filename)
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Transfer-Encoding", "binary")
		c.Header("Content-Disposition", "attachment; filename="+file)
		c.Header("Content-Type", "application/octet-stream")
		c.File(file) // 可下载文件
	}
}
