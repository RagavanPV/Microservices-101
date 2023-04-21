package main

import (
    Logger "github.com/ragavan/go_logger"
    LoggerModel "github.com/ragavan/go_logger/model"
)

func main() {
    logger, err := Logger.Init(LoggerModel.LogOptions{Level: "info"})
    Logger.SetInstance(logger)
    logger1 := Logger.GetInstance()
    // handle err here
    if err != nil {
		return 
	}
    logger.Warn().Msg("there is no real going back")
    logger1.Warn().Msg("there is no real going front")
}
