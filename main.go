package main

import (
	"gitee/Job/data_clear/internal/backend"
	"gitee/Job/data_clear/internal/config"
	"gitee/Job/data_clear/internal/pipeline"
	"gitee/Job/data_clear/internal/processor"
	"gitee/Job/data_clear/internal/receiver"
	"log"
	"time"
)

func main() {
	// 加载配置文件
	cfg, err := config.LoadConfigByViper()
	if err != nil {
		// 配置文件读取失败，打印并直接退出
		log.Fatalf("读取配置文件出错: %v", err)
	}

	// 初始化 Receiver
	fileReceiver := receiver.NewFileReceiver(cfg.InputFilePath)
	defer fileReceiver.Close() // 确保文件最终关闭

	// 初始化 Pipeline
	pipeline := pipeline.NewPipeline()
	pipeline.AddProcessor(processor.NewFilter(cfg.FilterConfig))
	pipeline.AddProcessor(processor.NewFill(cfg.FillConfig))
	aggregator := processor.NewAggregator(cfg.AggregatorConfig)
	pipeline.AddProcessor(aggregator)

	// 初始化 Backend
	fileBackend := backend.NewFileBackend(cfg.OutputFilePath)

	// 定时器 聚合和发送数据是周期性【时间驱动】
	ticker := time.NewTicker(10 * time.Second) // 每 10 秒聚合一次
	defer ticker.Stop()

	// 处理数据
outerLoop:
	for {
		select {
		case <-ticker.C:
			// 周期性聚合
			aggregatedData := aggregator.Aggregate()
			for _, data := range aggregatedData {
				err := fileBackend.Send(data)
				if err != nil {
					log.Fatalf("向后端发送数据失败: %v", err)
				}
			}
		default:
			// 读取并处理数据
			data, err := fileReceiver.Receive()
			if err != nil {
				log.Fatalf("读取数据失败: %v", err)
			}

			// 文件读取完毕，退出循环
			if data == nil {
				break outerLoop
			}

			// 处理数据
			_, err = pipeline.Process(data)
			if err != nil {
				log.Fatalf("处理数据失败: %v", err)
			}
		}
	}

	// 数据处理完毕，等待聚合和发送数据（确保所有数据都被处理并发送到后端）
	aggregatedData := aggregator.Aggregate()
	for _, data := range aggregatedData {
		err := fileBackend.Send(data)
		if err != nil {
			log.Fatalf("向后端发送数据失败: %v", err)
		}
	}

	log.Println("数据全部清洗完成.")
}
