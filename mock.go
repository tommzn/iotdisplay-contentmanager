package main

import (
	"errors"

	"github.com/aws/aws-sdk-go/service/iotdataplane"
)

type publisherMock struct {
	contentResponse   map[string][]ContentResponse
	shouldReturnError bool
}

func newPublisherMock() *publisherMock {
	return &publisherMock{
		contentResponse:   make(map[string][]ContentResponse),
		shouldReturnError: false,
	}
}

func (mock *publisherMock) Send(response ContentResponse, topic string) error {

	if mock.shouldReturnError {
		return errors.New("Am error has occured!")
	}

	if _, ok := mock.contentResponse[topic]; !ok {
		mock.contentResponse[topic] = []ContentResponse{}
	}
	mock.contentResponse[topic] = append(mock.contentResponse[topic], response)

	return nil
}

type iotClientMock struct {
	publishInputs     []*iotdataplane.PublishInput
	shouldReturnError bool
}

func newIotClientMock() *iotClientMock {
	return &iotClientMock{
		publishInputs:     []*iotdataplane.PublishInput{},
		shouldReturnError: false,
	}
}

func (mock *iotClientMock) Publish(request *iotdataplane.PublishInput) (*iotdataplane.PublishOutput, error) {

	mock.publishInputs = append(mock.publishInputs, request)

	if mock.shouldReturnError {
		return nil, errors.New("Am error has occured!")
	}
	return &iotdataplane.PublishOutput{}, nil
}
