package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

// ReplaceFileExt 将指定文件的后缀改变
func ReplaceFileExt(path, newExt string) string {
	dir, file := filepath.Split(path)
	ext := filepath.Ext(file)
	if ext == "" {
		return path + "." + newExt
	}
	oldExt := strings.TrimPrefix(ext, ".")
	newFile := strings.TrimSuffix(file, oldExt) + newExt
	return filepath.Join(dir, newFile)
}

// GetMD5 获取字符串或者文件(提供文件路径)的md5
func GetMD5(s string, size ...int) (string, error) {
	// 如果不传入size，则默认为1024
	chunkSize := 1024
	if len(size) > 0 {
		chunkSize = size[0]
	}
	var r io.Reader
	if file, err := os.Open(s); err == nil {
		defer func(file *os.File) {
			_ = file.Close()
		}(file)
		r = file
	} else {
		r = strings.NewReader(s)
	}

	hash := md5.New()

	buffer := make([]byte, chunkSize)
	for {
		n, err := r.Read(buffer)
		if err != nil && err != io.EOF {
			return "", err
		}
		if n == 0 {
			break
		}
		if _, err := hash.Write(buffer[:n]); err != nil {
			return "", err
		}
	}

	hashInBytes := hash.Sum(nil)[:16]
	return hex.EncodeToString(hashInBytes), nil
}

// GetFileMD5 直接计算文件md5
func GetFileMD5(file *os.File, size ...int) (string, error) {
	defer func(file *os.File, offset int64, whence int) {
		_, _ = file.Seek(offset, whence)
	}(file, 0, 0) // 重置文件指针到文件开头

	// 如果不传入size，则默认为1024
	chunkSize := 1024
	if len(size) > 0 {
		chunkSize = size[0]
	}
	hash := md5.New()
	buffer := make([]byte, chunkSize)
	var r io.Reader
	r = file
	for {
		n, err := r.Read(buffer)
		if err != nil && err != io.EOF {
			return "", err
		}
		if n == 0 {
			break
		}
		if _, err := hash.Write(buffer[:n]); err != nil {
			return "", err
		}
	}

	hashInBytes := hash.Sum(nil)[:16]
	return hex.EncodeToString(hashInBytes), nil
}

// ConvertFileHeaderToFile 文件类型转换
func ConvertFileHeaderToFile(fh *multipart.FileHeader) (*os.File, error) {
	// 创建一个管道
	r, w := io.Pipe()

	// 创建一个 goroutine 将上传的文件内容写入管道中
	go func() {
		defer func(w *io.PipeWriter) {
			_ = w.Close()
		}(w)
		f, err := fh.Open()
		if err != nil {
			_ = w.CloseWithError(err)
			return
		}
		defer func(f multipart.File) {
			_ = f.Close()
		}(f)
		_, err = io.Copy(w, f)
		if err != nil {
			_ = w.CloseWithError(err)
			return
		}
	}()

	// 将管道中的内容写入一个临时文件
	f, err := os.CreateTemp("", "upload.*")
	if err != nil {
		return nil, err
	}
	defer func(name string) {
		_ = os.Remove(name)
	}(f.Name())
	_, err = io.Copy(f, r)
	if err != nil {
		_ = f.Close()
		return nil, err
	}

	// 将临时文件的读写位置指向文件开头
	_, err = f.Seek(0, 0)
	if err != nil {
		_ = f.Close()
		return nil, err
	}

	// 返回临时文件的句柄
	return f, nil
}

// ListSubdirs 遍历某个文件夹下面第一层的文件夹
func ListSubdirs(dir string, isOnlyName ...bool) ([]string, error) {
	isName := false
	if (len(isOnlyName) > 0) && isOnlyName[0] {
		isName = isOnlyName[0]
	}
	subdirs := make([]string, 0)
	f, err := os.Open(dir)
	if err != nil {
		return subdirs, err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	files, err := f.Readdir(-1)
	if err != nil {
		return subdirs, err
	}
	for _, file := range files {
		if file.IsDir() {
			var subdir = ""
			if isName {
				subdir = file.Name()
			} else {
				subdir = filepath.Join(dir, file.Name())
			}
			subdirs = append(subdirs, subdir)
		}
	}
	return subdirs, nil
}

// GetSubdirs 获取某文件夹下面的文件夹列表
func GetSubdirs(dir string) ([]string, error) {
	subdirs := make([]string, 0)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != dir {
			subdirs = append(subdirs, path)
		}
		return nil
	})
	if err != nil {
		return subdirs, err
	}
	return subdirs, nil
}

// Contains 查看字符串数组中是否存在
func Contains(sliceOrString interface{}, x interface{}) bool {
	switch reflect.TypeOf(sliceOrString).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(sliceOrString)
		for i := 0; i < s.Len(); i++ {
			if s.Index(i).Interface() == x {
				return true
			}
		}
		return false
	case reflect.String:
		s := reflect.ValueOf(sliceOrString).String()
		for _, c := range s {
			if c == x {
				return true
			}
		}
		return false
	default:
		return false
	}
}
