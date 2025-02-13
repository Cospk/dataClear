package receiver

import (
	"bufio"
	"encoding/json"
	"gitee/Job/data_clear/internal/utils"
	"os"
)

// FileReceiver 结构体，进行流式读取文件
type FileReceiver struct {
	FilePath string
	file     *os.File
	scanner  *bufio.Scanner
}

// NewFileReceiver 创建一个新的 FileReceiver
func NewFileReceiver(filePath string) *FileReceiver {
	return &FileReceiver{
		FilePath: filePath,
	}
}

// Open 打开文件并初始化 scanner
func (f *FileReceiver) Open() error {
	file, err := os.Open(f.FilePath)
	if err != nil {
		return err
	}
	f.file = file
	f.scanner = bufio.NewScanner(file)
	return nil
}

// Close 关闭文件
func (f *FileReceiver) Close() error {
	if f.file != nil {
		return f.file.Close()
	}
	return nil
}

// Receive 逐行读取文件并返回数据
func (f *FileReceiver) Receive() (utils.Data, error) {
	// 如果 scanner 为空，则打开文件并初始化 scanner
	if f.scanner == nil {
		if err := f.Open(); err != nil {
			return nil, err
		}
	}

	// 逐行读取文件
	if f.scanner.Scan() {
		line := f.scanner.Text()
		var data utils.Data
		if err := json.Unmarshal([]byte(line), &data); err != nil {
			return nil, err
		}
		return data, nil
	}

	return nil, nil // 文件读取完毕，返回 nil
}
