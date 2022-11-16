package container

import (
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/fangaoxs/synk/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const ()

// TextsContainer 上传文本
// 1 获取用户输入的内容
// 2 生成随机文件名
// 3 将内容写入文件并写入目录中
func TextsContainer() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json struct {
			Raw string `json:"raw"`
		}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fileName := uuid.NewString() + ".txt"
		if err := ioutil.WriteFile(filepath.Join(utils.GetUploadsDir(), fileName), []byte(json.Raw), 0644); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{"filename": fileName})
	}
}

// FilesContainer 上传文件
// 1 获取 file 流
// 2 生成随机文件名
// 3 将文件写入目录
func FilesContainer() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("raw")
		if err != nil {
			log.Fatal(err)
		}

		filename := uuid.NewString() + filepath.Ext(file.Filename)
		if err = c.SaveUploadedFile(file, filepath.Join(utils.GetUploadsDir(), filename)); err != nil {
			log.Fatal("save file failed:", err)
		}

		c.JSON(http.StatusOK, gin.H{"filename": filename})
	}
}
