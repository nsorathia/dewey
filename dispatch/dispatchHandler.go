package dispatch

import (
	"dewey/indexers"
	"dewey/logging"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

//Handle reads a kafka message, maps the event_type to an Indexer and calls index.
func handle(msg *kafka.Message) {
	logging.Logger.WithFields(logrus.Fields{"Topic:": *msg.TopicPartition.Topic, "Message": string(msg.Value)}).Infoln("Begin - handling kafka message")

	//TODO:  see if we can standardize this json instead of using unstructured json ...would be better performant
	var data map[string]interface{}

	if err := json.Unmarshal(msg.Value, &data); err != nil {
		logging.Logger.WithFields(logrus.Fields{"error:": err.Error()}).Errorln("Cold not Unmarshal msg")
	}

	eventType := data["event_type"].(string)
	indexer, ok := indexers.IndexerMap[eventType]
	if !ok {
		logging.Logger.WithFields(logrus.Fields{"eventType:": eventType}).Infoln("Event type not registered")
		return
	}

	//execute async

	go indexer.Index(data)

	// go func(data map[string]interface{}) {
	// 	indexer.Index(data)
	// }(data)

	logging.Logger.Infoln("End - handling kafka message!")
}
