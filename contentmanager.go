package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/mitchellh/hashstructure/v2"
	log "github.com/tommzn/go-log"
)

// newContentManager creates a content manager with given logger.
func newContentManager(logger log.Logger, publisher ContentPublisher) *ContentManager {
	return &ContentManager{
		logger:    logger,
		publisher: publisher,
	}
}

// GetContent will collect contents for passed device and publish it to a device related topic.
func (mgr *ContentManager) GetContent(ctx context.Context, refreshRequest ContentRefreshRequest) {

	logger := mgr.loggerWithContext(ctx, refreshRequest)
	defer logger.Flush()

	logger.Debugf("Receive content refresh request for device: %s, topic: %s", refreshRequest.ThingName, refreshRequest.Topic)

	targetTopic := getTargetTopic(refreshRequest.Topic)
	logger.Debug("Contents will be published to: ", targetTopic)

	response := ContentResponse{
		Items: []ContentItem{
			ContentItem{
				Position: Position{
					X: 10,
					Y: 10,
				},
				Text: "Hi!",
			},
			ContentItem{
				Position: Position{
					X: 10,
					Y: 40,
				},
				Text: fmt.Sprintf("I'm %s.", refreshRequest.ThingName),
			},
		},
	}

	if err := mgr.publisher.Send(response, targetTopic); err != nil {
		logger.Error("Message publishing to topic %s failed, reason: %s: ", refreshRequest.Topic, err)
	}

	logger.Debug("Content successful published to topic: ", refreshRequest.Topic)
}

// loggerWithContext adds values from content refresh request to log content and returns current logger.
func (mgr *ContentManager) loggerWithContext(ctx context.Context, refreshRequest ContentRefreshRequest) log.Logger {
	contextValues := make(map[string]string)
	contextValues["topic"] = refreshRequest.Topic
	contextValues["clientid"] = refreshRequest.ThingName
	contextValues[log.LogCtxNamespace] = "iot-display-contentmanager"
	return log.AppendFromLambdaContext(log.AppendContextValues(mgr.logger, contextValues), ctx)
}

// getTargetTopic will remove last part of a topic, usually "/get", and returns it.
func getTargetTopic(topic string) string {
	return strings.TrimSuffix(topic, "/get")
}

// addContentHash generates a hash for all items in a content response.
func addContentHash(response *ContentResponse) {
	if hash, err := hashstructure.Hash(response.Items, hashstructure.FormatV2, nil); err == nil {
		response.Hash = fmt.Sprintf("%d", hash)
	}
}
