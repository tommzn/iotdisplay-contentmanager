package main

import (
	"context"
	"strings"

	log "github.com/tommzn/go-log"
)

// newContentManager creates a content manager with given logger.
func newContentManager(logger log.Logger) *ContentManager {
	return &ContentManager{logger: logger}
}

// GetContent will collect contents for passed device and publish it to a device related topic.
func (mgr *ContentManager) GetContent(ctx context.Context, refreshRequest ContentRefreshRequest) {

	logger := mgr.loggerWithContext(refreshRequest)
	defer logger.Flush()

	logger.Debugf("Receive content refresh request for device: %s, topic: %s", refreshRequest.ThingName, refreshRequest.Topic)

	targetTopic := getTargetTopic(refreshRequest.Topic)
	logger.Debug("Contents will be published to: ", targetTopic)
}

// loggerWithContext adds values from content refresh request to log content and returns current logger.
func (mgr *ContentManager) loggerWithContext(refreshRequest ContentRefreshRequest) log.Logger {

	contextValues := make(map[string]string)
	contextValues["topic"] = refreshRequest.Topic
	contextValues["clientid"] = refreshRequest.ThingName
	contextValues[log.LogCtxNamespace] = "iot-display-contentmanager"
	return log.AppendContextValues(mgr.logger, contextValues)
}

// getTargetTopic will remove last part of a topic, usually "/get", and returns it.
func getTargetTopic(topic string) string {
	return strings.TrimSuffix(topic, "/get")
}
