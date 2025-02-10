package main

import (
	"github.com/0xsj/conure/pkg/utils"
)

func main() {
	logger := utils.DefaultLogger
	logger.Info("News Aggregator started successfully", nil)
}