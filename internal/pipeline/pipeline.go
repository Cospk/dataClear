package pipeline

import (
	"gitee/Job/data_clear/internal/processor"
	"gitee/Job/data_clear/internal/utils"
)

// Processor 接口定义
type Processor interface {
	Process(data utils.Data) (utils.Data, error)
}

// Pipeline 结构
type Pipeline struct {
	Processors []Processor
}

// NewPipeline 创建一个新的 Pipeline
func NewPipeline() *Pipeline {
	return &Pipeline{}
}

// AddProcessor 添加一个处理器到管道中
func (p *Pipeline) AddProcessor(processor processor.Processor) {
	// 将处理器添加到管道的处理器列表中
	p.Processors = append(p.Processors, processor)
}

// Process 处理数据通过管道中的所有处理器
func (p *Pipeline) Process(data utils.Data) (utils.Data, error) {
	var err error
	for _, processor1 := range p.Processors {
		data, err = processor1.Process(data)
		if err != nil {
			return nil, err
		}
		if data == nil {
			return nil, nil // 数据被过滤掉
		}
	}
	return data, nil
}
