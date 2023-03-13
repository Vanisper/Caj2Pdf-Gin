package md5

import (
	"crypto/md5"
	"fmt"
	"syscall/js"
	"time"
)

// GetFileMd5 第一个参数是入参的第一个  第二个参数是将js入参作为数组传进来  适用于多入参场景
func GetFileMd5(_ js.Value, args []js.Value) interface{} {
	startTime := time.Now()
	array := args[0]
	byteLength := array.Get("byteLength").Int()

	if len(args) > 1 {
		chunkSize := args[1].Int()
		if chunkSize <= byteLength {
			// 执行分段md5
			count := byteLength / chunkSize
			var itemSize []int
			if count == 0 {
				itemSize = append(itemSize, byteLength)
			} else {
				for i := 0; i < count; i++ {
					itemSize = append(itemSize, chunkSize)
				}
				if byteLength%chunkSize != 0 {
					itemSize = append(itemSize, byteLength%chunkSize)
				}
			}
			for index, value := range itemSize {
				start := 0
				if index == 0 {
					start = index * value
				} else {
					start = index * itemSize[index-1]
				}
				end := start + value
				fmt.Println(start, end)
				itemByte := array.Call("slice", start, end)
				var buffer []uint8 = make([]uint8, value)
				js.CopyBytesToGo(buffer, itemByte)
				wasmMd5Add(buffer)
			}
			md5hash := wasmMd5End()
			elapsed := time.Since(startTime)
			fmt.Printf("elapsed: %v\n", elapsed)
			fmt.Printf("md5hash: %v\n", md5hash)
			return md5hash
		}
	}

	var buffer []uint8 = make([]uint8, byteLength)
	js.CopyBytesToGo(buffer, array)
	md5hashByteArr := md5.Sum(buffer)
	md5hash := fmt.Sprintf("%x", md5hashByteArr)
	// endTime := time.Now()
	// elapsed := endTime.Sub(startTime).Seconds()
	elapsed := time.Since(startTime)
	fmt.Printf("elapsed: %v\n", elapsed)
	fmt.Printf("md5hash: %v\n", md5hash)
	return md5hash
}

var md5hash = md5.New()

func WasmMd5Add(value js.Value, args []js.Value) interface{} {
	array := args[0]
	byteLength := array.Get("byteLength").Int()
	var buffer []uint8 = make([]uint8, byteLength)
	js.CopyBytesToGo(buffer, array)
	md5hash.Write(buffer)
	return nil
}

func wasmMd5Add(buffer []uint8) interface{} {
	md5hash.Write(buffer)
	return nil
}

func WasmMd5End(value js.Value, args []js.Value) interface{} {
	finalMd5Hash := fmt.Sprintf("%x", md5hash.Sum(nil))
	md5hash.Reset() // 清空  否则下次使用时会有问题
	return finalMd5Hash
}

func wasmMd5End() string {
	finalMd5Hash := fmt.Sprintf("%x", md5hash.Sum(nil))
	md5hash.Reset() // 清空  否则下次使用时会有问题
	return finalMd5Hash
}
