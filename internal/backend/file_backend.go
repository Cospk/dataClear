package backend

import (
	"encoding/json"
	"gitee/Job/data_clear/internal/utils"
	"os"
)

type FileBackend struct {
	FilePath string
}

func NewFileBackend(filePath string) *FileBackend {
	return &FileBackend{FilePath: filePath}
}

func (f *FileBackend) Send(data utils.Data) error {
	file, err := os.OpenFile(f.FilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将data数据转为json格式，然后写入文件中去
	encoder := json.NewEncoder(file)
	return encoder.Encode(data)
}
