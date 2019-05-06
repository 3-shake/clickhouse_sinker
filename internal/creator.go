package internal

import (
	"github.com/3-shake/clickhouse_sinker/input"
	"github.com/3-shake/clickhouse_sinker/output"
	"github.com/3-shake/clickhouse_sinker/parser"
	"github.com/3-shake/clickhouse_sinker/task"
	"github.com/3-shake/clickhouse_sinker/util"
)

// GenTasks generate the tasks via config
func (config *Config) GenTasks() []*task.TaskService {
	res := make([]*task.TaskService, 0, len(config.Tasks))
	for _, taskConfig := range config.Tasks {
		kafka := config.GenInput(taskConfig)
		ck := config.GenOutput(taskConfig)
		p := parser.NewParser(taskConfig.Parser)

		taskImpl := task.NewTaskService(kafka, ck, p)

		util.IngestConfig(taskConfig, taskImpl)

		if taskImpl.FlushInterval == 0 {
			taskImpl.FlushInterval = config.Common.FlushInterval
		}

		if taskImpl.BufferSize == 0 {
			taskImpl.BufferSize = config.Common.BufferSize
		}

		res = append(res, taskImpl)
	}
	return res
}

// GenInput generate the input via config
func (config *Config) GenInput(taskCfg *Task) *input.Kafka {
	kfkCfg := config.Kafka[taskCfg.Kafka]

	inputImpl := input.NewKafka()
	util.IngestConfig(taskCfg, inputImpl)
	util.IngestConfig(kfkCfg, inputImpl)
	return inputImpl
}

// GenOutput generate the output via config
func (config *Config) GenOutput(taskCfg *Task) *output.ClickHouse {
	ckCfg := config.Clickhouse[taskCfg.Clickhouse]

	outputImpl := output.NewClickHouse()

	util.IngestConfig(ckCfg, outputImpl)
	util.IngestConfig(taskCfg, outputImpl)
	util.IngestConfig(config.Common, outputImpl)
	return outputImpl
}
