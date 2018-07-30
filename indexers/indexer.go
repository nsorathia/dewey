package indexers

//IndexerMap is a map of type Indexers
var IndexerMap = make(map[string]Indexer)

//Indexer is an abstraction for a specific indexer type
type Indexer interface {
	Index(objMap map[string]interface{})
}
