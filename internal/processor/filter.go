package processor

import (
	"gitee/Job/data_clear/internal/config"
	"gitee/Job/data_clear/internal/utils"
	"strconv"
	"strings"
)

// Filter 过滤器
type Filter struct {
	Field     string
	Condition func(interface{}) bool
}

func NewFilter(cfg config.FilterConfig) *Filter {
	return &Filter{
		Field:     cfg.Field,
		Condition: parseCondition(cfg.Condition),
	}
}

// Process 处理数据：根据条件过滤数据
func (f *Filter) Process(data utils.Data) (utils.Data, error) {
	if value, ok := data[f.Field]; ok {
		if f.Condition(value) {
			return data, nil
		}
	}
	return nil, nil // 过滤掉不符合条件的数据
}

func parseCondition(condition string) func(interface{}) bool {
	// 解析条件并返回一个函数
	var op string
	var value interface{}

	// 解析条件字符串，提取操作符和值
	if strings.Contains(condition, ">") {
		op = ">"
		value = strings.TrimSpace(strings.TrimPrefix(condition, ">"))
	} else if strings.Contains(condition, "<") {
		op = "<"
		value = strings.TrimSpace(strings.TrimPrefix(condition, "<"))
	} else if strings.Contains(condition, "=") {
		op = "="
		value = strings.TrimSpace(strings.TrimPrefix(condition, "="))
	} else if strings.Contains(condition, "!=") {
		op = "!="
		value = strings.TrimSpace(strings.TrimPrefix(condition, "!="))
	} else {
		// 默认条件为true（即不过滤）
		return func(interface{}) bool {
			return true
		}
	}

	// 返回过滤函数
	return func(input interface{}) bool {
		switch input.(type) {
		case string:
			// 字符串比较
			switch op {
			case "==":
				return input == value
			case "!=":
				return input != value
			default:
				return false
			}
		case int, int64, float64:
			// 数值比较
			// 将所有的数字转为float64
			inputFloat, err := strconv.ParseFloat(input.(string), 64)
			if err != nil {
				return false
			}
			valueFloat, err := strconv.ParseFloat(value.(string), 64)
			if err != nil {
				return false
			}
			switch op {
			case ">":
				return inputFloat > valueFloat
			case "<":
				return inputFloat < valueFloat
			case "==":
				return inputFloat == valueFloat
			case "!=":
				return inputFloat != valueFloat
			default:
				return false
			}

		default:
			//不支持类型
			return false
		}
	}

}
