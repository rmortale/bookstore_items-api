package elasticsearch

import (
	"context"
	"github.com/olivere/elastic/v7"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(client *elastic.Client)
	Index(interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

func InitClient() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
	)
	if err != nil {
		panic(err)
	}
	defer client.Stop()
	Client.setClient(client)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	return c.client.Index().Do(ctx)
}
