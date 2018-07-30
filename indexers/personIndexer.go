package indexers

import (
	"context"
	"dewey/common"
	"dewey/types/person"
	"strconv"
	//"github.com/sirupsen/logrus"
)

var esClient = common.ESClient()

func init() {

	//create the Index if it does not exit
	ctx := context.Background()
	if !common.IndexExists(ctx, person.IndexName) {
		common.CreateIndex(ctx, person.IndexName, person.Mapping)
	}

	//register indexer
	IndexerMap["melody.person.created"] = &MelodyPersonIndexer{}
}

//MelodyPersonIndexer indexes a Melody Person type
type MelodyPersonIndexer struct{}

//Index parses the object map, creates a type and calls the ElasticSearch APIs to index a document
func (mp *MelodyPersonIndexer) Index(objMap map[string]interface{}) {

	document, err := MapDocument(objMap)
	if err != nil {
		//log error
		return
	}

	common.IndexNewDocument(person.IndexName, person.TypeName, strconv.Itoa(document.ID), document)
}
