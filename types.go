package main

import log "github.com/tommzn/go-log"

// ContentManager delivers content for requested devices.
type ContentManager struct {
	logger    log.Logger
	publisher ContentPublisher
}

// AwsIotPublisher is used to send content response message to MQTT topics on AWS IOT.
type AwsIotPublisher struct {
	logger    log.Logger
	iotClient AwsIotClient
}

// ContentRefreshRequest is used to trigger content refresh for passed AWS IOT device.
type ContentRefreshRequest struct {
	Topic     string `json:"topic_name"`
	ThingName string `json:"thing_name"`
}

// ContentResponse contains entire content an AWS IOT device should display.
type ContentResponse struct {
	Hash  string        `json:"content_hash"`
	Items []ContentItem `json:"items"`
}

// ContentItem is a single item, e.g. a text.
type ContentItem struct {
	Position Position `json:"position"`
	Text     string   `json:"text"`
}

// Position defines where to place content items on a screen.
type Position struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}
