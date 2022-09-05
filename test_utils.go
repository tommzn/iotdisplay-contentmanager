package main

import (
	config "github.com/tommzn/go-config"
	log "github.com/tommzn/go-log"
)

// loadConfigForTest loads test config.
func loadConfigForTest(configFile *string) config.Config {

	if configFile == nil {
		configFile = config.AsStringPtr("testconfig.yml")
	}
	configLoader := config.NewFileConfigSource(configFile)
	config, _ := configLoader.Load()
	return config
}

// loggerForTest creates a new stdout logger for testing.
func loggerForTest() log.Logger {
	return log.NewLogger(log.Debug, nil, nil)
}

func contentManagerForTest() *ContentManager {
	logger := loggerForTest()
	publisher := &AwsIotPublisher{logger: logger, iotClient: newIotClientMock()}
	return newContentManager(logger, publisher)
}
