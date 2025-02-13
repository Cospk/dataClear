package processor

import (
	"gitee/Job/data_clear/internal/config"
	"gitee/Job/data_clear/internal/utils"
)

// Fill 填充处理器
type Fill struct {
	Field string
	Value interface{}
}

// NewFill 创建一个新的 Fill 处理器
func NewFill(cfg config.FillConfig) *Fill {
	return &Fill{
		Field: cfg.Field,
		Value: cfg.Value,
	}
}

// Process 处理数据：将指定字段的值填充为指定的值
func (f *Fill) Process(data utils.Data) (utils.Data, error) {
	data[f.Field] = f.Value
	return data, nil
}
