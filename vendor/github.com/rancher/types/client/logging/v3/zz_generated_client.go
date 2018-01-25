package client

import (
	"github.com/rancher/norman/clientbase"
)

type Client struct {
	clientbase.APIBaseClient

	Logging        LoggingOperations
	ProjectLogging ProjectLoggingOperations
}

func NewClient(opts *clientbase.ClientOpts) (*Client, error) {
	baseClient, err := clientbase.NewAPIClient(opts)
	if err != nil {
		return nil, err
	}

	client := &Client{
		APIBaseClient: baseClient,
	}

	client.Logging = newLoggingClient(client)
	client.ProjectLogging = newProjectLoggingClient(client)

	return client, nil
}
