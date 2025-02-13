package processor

import "gitee/Job/data_clear/internal/utils"

// Processor 接口定义
type Processor interface {
	Process(data utils.Data) (utils.Data, error)
}
