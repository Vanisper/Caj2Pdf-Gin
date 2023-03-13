package controllers

import (
	"Caj2PdfServer/configs"
	"Caj2PdfServer/middlewares"
	"Caj2PdfServer/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

type OutPut struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	Message string `json:"message"`
}

func checkUpload(md5 string, path string) bool {
	dirs, _ := utils.ListSubdirs(path, true)
	return utils.Contains(dirs, md5)
}

// Upload 上传文件
func Upload(c *gin.Context) {
	// 从请求中获取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	toFile, _ := utils.ConvertFileHeaderToFile(file)
	md5, _ := utils.GetFileMD5(toFile)
	//检查是否存在历史文件
	if !checkUpload(md5, configs.UploadPath) {
		// 将文件保存到本地
		outPath := path.Join(configs.UploadPath, md5, file.Filename)
		if err := c.SaveUploadedFile(file, outPath); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("文件上传失败: %s", err.Error()))
			return
		}
		middlewares.Caj2pdf(outPath, configs.TemplatesOutPut)
		c.JSON(http.StatusOK, OutPut{
			Name:    utils.ReplaceFileExt(file.Filename, "pdf"),
			Url:     fmt.Sprintf("%s%s/%s", configs.TemplatesOutPutShort, md5, utils.ReplaceFileExt(file.Filename, "pdf")),
			Message: "文件上传且转换成功",
		})
	} else {
		c.JSON(http.StatusOK, OutPut{
			Name:    utils.ReplaceFileExt(file.Filename, "pdf"),
			Url:     fmt.Sprintf("%s%s/%s", configs.TemplatesOutPutShort, md5, utils.ReplaceFileExt(file.Filename, "pdf")),
			Message: "文件已有历史记录",
		})
	}

}
