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

	mgr := newContentManager(loggerForTest())
	request := ContentRefreshRequest{
		Topic:     "iotdisplay/things/Thing001/contents/get",
		ThingName: "Thing001",
	}
	mgr.GetContent(context.Background(), request)
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
