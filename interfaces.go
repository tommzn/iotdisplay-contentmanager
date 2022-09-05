package main

import (
	"context"

	"github.com/aws/aws-sdk-go/service/iotdataplane"
)

// Handler will process content refresh requests.
type Handler interface {

	// GetContent collects and pulished contents for passed AWS IOT device.
	GetContent(context.Context, ContentRefreshRequest)
}

// ContentPublisher publishes generated contents for AWS IOT devices.
type ContentPublisher interface {

	// Sends given content response to passed topic.
	Send(response ContentResponse, topic string) error
}

// AwsIotClient is used to publish messages to MQTT topics on AWS IOT.
type AwsIotClient interface {

	// Publish sends messages to topic passed via request input
	Publish(*iotdataplane.PublishInput) (*iotdataplane.PublishOutput, error)
}
