package main

import "context"

// Handler will process content refresh requests.
type Handler interface {

	// GetContent collects and pulished contents for passed AWS IOT device.
	GetContent(context.Context, ContentRefreshRequest)
}
