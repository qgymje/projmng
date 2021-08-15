package repository

import (
	"encoding/json"
	"projmng/pkg/event"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ProjectCommandKafka struct {
	producer *kafka.Producer
	topic    string
}

func (pck *ProjectCommandKafka) Create(evt event.Project) error {
	return pck.publish(evt)
}

func (pck *ProjectCommandKafka) Update(evt event.Project) error {
	return pck.publish(evt)
}

func (pck *ProjectCommandKafka) Delete(evt event.Project) error {
	return pck.publish(evt)
}

func (pck *ProjectCommandKafka) publish(evt event.Project) error {
	data, err := json.Marshal(&evt)
	if err != nil {
		return err
	}

	err = pck.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &pck.topic, Partition: kafka.PartitionAny},
		Value:          data,
	}, nil)

	return err
}
