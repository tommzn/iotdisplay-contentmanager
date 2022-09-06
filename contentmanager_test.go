package main

import (
	"context"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ContentManagerTestSuite struct {
	suite.Suite
}

func TestContentManagerTestSuite(t *testing.T) {
	suite.Run(t, new(ContentManagerTestSuite))
}

func (suite *ContentManagerTestSuite) TestGetContent() {

	publisherMock := newPublisherMock()

	mgr := contentManagerForTest()
	mgr.publisher = publisherMock

	request := ContentRefreshRequest{
		Topic:     "iotdisplay/things/Thing001/contents/get",
		ThingName: "Thing001",
	}
	mgr.GetContent(context.Background(), request)

	suite.Len(publisherMock.contentResponse, 1)
}

func (suite *ContentManagerTestSuite) TestGetContentWithError() {

	publisherMock := newPublisherMock()
	publisherMock.shouldReturnError = true

	mgr := contentManagerForTest()
	mgr.publisher = publisherMock

	request := ContentRefreshRequest{
		Topic:     "iotdisplay/things/Thing001/contents/get",
		ThingName: "Thing001",
	}
	mgr.GetContent(context.Background(), request)

	suite.Len(publisherMock.contentResponse, 0)
}

func (suite *ContentManagerTestSuite) TestGetTargetTopic() {

	suite.Equal("iotdisplay/things/Thing001/contents", getTargetTopic("iotdisplay/things/Thing001/contents/get"))
	suite.Equal("iotdisplay/things/Thing001/contents", getTargetTopic("iotdisplay/things/Thing001/contents"))
	suite.Equal("iotdisplay/things/Thing001/contents/update", getTargetTopic("iotdisplay/things/Thing001/contents/update"))
}

func (suite *ContentManagerTestSuite) TestBootstrap() {

	mgr := bootstrap(loadConfigForTest(nil))
	suite.NotNil(mgr)
}

func (suite *ContentManagerTestSuite) TestAddContentHash() {

	response := ContentResponse{
		Items: []ContentItem{
			ContentItem{
				Position: Position{
					X: 10,
					Y: 10,
				},
				Text: "Test-Item",
			},
		},
	}
	addContentHash(&response)
	suite.NotEqual("", response.Hash)
}
