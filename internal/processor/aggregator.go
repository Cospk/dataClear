package processor

import (
	"gitee/Job/data_clear/internal/config"
	"gitee/Job/data_clear/internal/utils"
)

// Aggregator 聚合器
type Aggregator struct {
	Field        string
	GroupByField string
	Aggregations map[string]func([]float64) float64
	Buffer       map[string][]utils.Data // 按分组字段存储数据
}

// NewAggregator 创建一个新的 Aggregator
func NewAggregator(cfg config.AggregatorConfig) *Aggregator {
	aggregations := make(map[string]func([]float64) float64)
	for key, aggType := range cfg.Aggregations {
		switch aggType {
		case "sum":
			aggregations[key] = utils.Sum
		case "count":
			aggregations[key] = utils.Count
		case "min":
			aggregations[key] = utils.Min
		case "max":
			aggregations[key] = utils.Max
		}
	}

	return &Aggregator{
		Field:        cfg.Field,
		GroupByField: cfg.GroupByField,
		Aggregations: aggregations,
		Buffer:       make(map[string][]utils.Data),
	}
}

// Process 方法用于处理数据，将数据添加到缓冲区中
func (a *Aggregator) Process(data utils.Data) (utils.Data, error) {
	// 获取分组字段的值
	groupKey, ok := data[a.GroupByField].(string)
	if !ok {
		return nil, nil // 如果分组字段不存在或类型错误，跳过该数据
	}

	// 将数据添加到对应的分组
	a.Buffer[groupKey] = append(a.Buffer[groupKey], data)
	return nil, nil // 聚合操作需要周期性处理，暂时不返回数据
}

// Aggregate 方法用于对缓冲区中的数据进行聚合操作
func (a *Aggregator) Aggregate() []utils.Data {
	var results []utils.Data

	// 遍历每个分组
	for _, groupData := range a.Buffer {
		var values []float64

		// 提取待聚合的字段值
		for _, data := range groupData {
			if value, ok := data[a.Field].(float64); ok {
				values = append(values, value)
			}
		}

		// 遍历聚合函数，对提取的字段值进行聚合操作
		for key, aggFunc := range a.Aggregations {
			result := make(utils.Data)
			for k, v := range groupData[0] { // 使用分组中第一条数据的元数据
				result[k] = v
			}
			result[a.Field+"_"+key] = aggFunc(values)
			results = append(results, result)
		}
	}

	a.Buffer = nil // 清空缓冲区
	return results
}
