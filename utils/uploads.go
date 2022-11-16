package utils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fangaoxs/synk/env"
)

// GetUploadsDir 获取存放下载文件目录的路径
func GetUploadsDir() string {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(exe)
	return filepath.Join(dir, env.ResourcesDir)
}

// 在当前执行文件同级目录创建存放下载的文件的目录，rwx: 111
// 1 获取当前执行文件的路径
// 2 获取当前执行文件所在目录
// 3 在当前执行文件所在目录下创建所需目录
func init() {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal("get path of current execution file failed:", err)
	}
	dir := filepath.Dir(exe)
	if err = os.MkdirAll(filepath.Join(dir, env.ResourcesDir), os.ModePerm); err != nil {
		log.Fatal("create directory of "+env.ResourcesDir+" failed", err)
	}
}
