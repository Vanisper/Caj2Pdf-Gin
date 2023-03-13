package middlewares

import (
	"Caj2PdfServer/configs"
	"Caj2PdfServer/utils"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Caj2pdf(inputFile string, out ...string) string {
	exe, _ := filepath.Abs(configs.LibCaj2PdfPath)
	outputFile := utils.ReplaceFileExt(inputFile, "pdf")
	if len(out) > 0 {
		md5, _ := utils.GetMD5(inputFile)
		oldExt := strings.TrimPrefix(filepath.Ext(inputFile), ".")
		newFile := filepath.Base(strings.TrimSuffix(inputFile, oldExt) + "pdf")
		outputPath := filepath.Join(out[0], md5)
		_, err := os.Stat(outputPath)
		if err == nil {
			fmt.Println(outputPath, "存在")
		} else if os.IsNotExist(err) {
			fmt.Println(outputPath, "不存在,即将创建")
			_ = os.MkdirAll(outputPath, os.ModePerm)
		} else {
			fmt.Println("error:", err)
		}
		outputFile = filepath.Join(outputPath, newFile)
	}
	//log.Printf(exe, "convert", inputFile, "-o", outputFile)
	cmd := exec.Command(exe, "convert", inputFile, "-o", outputFile)
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	log.Println(string(output))
	return outputFile
}
