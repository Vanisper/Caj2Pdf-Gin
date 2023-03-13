package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"
)

func replaceFileExt(path, newExt string) string {
	dir, file := filepath.Split(path)
	ext := filepath.Ext(file)
	if ext == "" {
		return path + "." + newExt
	}
	oldExt := strings.TrimPrefix(ext, ".")
	newFile := strings.TrimSuffix(file, oldExt) + newExt
	return filepath.Join(dir, newFile)
}

// Upload 上传文件
func Upload(c *gin.Context) {
	// 从请求中获取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	// 将文件保存到本地
	if err := c.SaveUploadedFile(file, "./assets/upload/"+file.Filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully", file.Filename))
	exe, _ := filepath.Abs(".\\lib\\caj2pdf\\caj2pdf.exe")
	inputFile, _ := filepath.Abs("./assets/upload/" + file.Filename)
	outputFile := replaceFileExt(inputFile, "pdf")
	cmd := exec.Command(exe, "convert", inputFile, "-o", outputFile)
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	log.Println(string(output))
}
