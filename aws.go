package main

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotdataplane"
	config "github.com/tommzn/go-config"
	log "github.com/tommzn/go-log"
)

func newAwsIotPublisher(conf config.Config, logger log.Logger) *AwsIotPublisher {

	sess, _ := session.NewSession(&aws.Config{
		Region:   conf.Get("aws.iot.region", nil),
		Endpoint: conf.Get("aws.iot.endpoint", nil),
	})
	return &AwsIotPublisher{
		logger:    logger,
		iotClient: iotdataplane.New(sess),
	}
}

func (publisher *AwsIotPublisher) Send(response ContentResponse, topic string) error {

	payload, err := json.Marshal(response)
	if err != nil {
		publisher.logger.Error("JSON encoding error: ", err)
		return err
	}

	publishInput := &iotdataplane.PublishInput{
		Topic:   aws.String(topic),
		Payload: payload,
		Qos:     aws.Int64(0),
	}

	if _, err := publisher.iotClient.Publish(publishInput); err != nil {
		publisher.logger.Error("Erro at publoshing to AWS IOT: ", err)
		return err
	}
	return nil
}
