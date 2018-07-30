package common

import (
	"fmt"
	"context"
	"github.com/olivere/elastic"
	"sync"
)

var esClient *elastic.Client
var once sync.Once

func init() {
	esClient = ESClient()
}

//ESClient is a singleton method that returns an *elastic.Client
func ESClient() *elastic.Client {
	once.Do(func() {
		instance, err := elastic.NewClient()
		if err != nil {
			panic(err)
		}
		esClient = instance
	})
	return esClient
}

//IndexExists checks to see if a index exist in the ElasticSearch cluster.
func IndexExists(ctx context.Context, indexName string) bool {
	exists, err := esClient.IndexExists(indexName).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	return exists
}

//CreateIndex creates a new index given the name and typeMapping.
func CreateIndex(ctx context.Context, indexName, typeMapping string) {
	index, err := esClient.CreateIndex(indexName).BodyString(typeMapping).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !index.ShardsAcknowledged {
		// Not acknowledged
	}
}

//IndexNewDocument uses the Elasticsearch apis to put a new document in the index
func IndexNewDocument(index, doctype, id string, document interface{}) {
	response, err := esClient.Index().
		Index(index).
		Type(doctype).
		Id(id).
		BodyJson(document).
		Do(context.Background())

	if err != nil || response.Status != 0 {
		//log an error
	}

	fmt.Println(document);
}
