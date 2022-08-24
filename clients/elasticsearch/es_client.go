package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/baglabs/bookstore_utils-go/logger"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

var (
	Client esClientInterface = &esClient{}
	log                      = logger.GetLogger()
)

type esClientInterface interface {
	SetClient(*elasticsearch.Client)
	Index(string, interface{}) (*esapi.Response, error)
}

type esClient struct {
	client *elasticsearch.Client
}

func Init() {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}
	Client.SetClient(client)
}

func (c *esClient) SetClient(client *elasticsearch.Client) {
	c.client = client
}

func (c *esClient) Index(index string, doc interface{}) (*esapi.Response, error) {
	data, err := json.Marshal(doc)
	if err != nil {
		log.Printf(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}

	ctx := context.Background()
	req := esapi.IndexRequest{
		Index: index,
		Body:  bytes.NewReader(data),
	}
	result, err := req.Do(ctx, c.client)

	if err != nil {
		log.Printf(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}

	return result, nil
}
