package elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/rmortale/bookstore_utils-go/logger"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(client *elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	log := logger.GetLogger()
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
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

func (c *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
		Index(index).
		BodyJson(doc).
		Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}
	return result, nil
}
