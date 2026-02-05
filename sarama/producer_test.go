package sarama

import (
	"testing"

	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
)

var addr = []string{"localhost:9094"}

func TestSyncProducer(t *testing.T) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(addr, config)

	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	//config.Producer.Partitioner = sarama.NewHashPartitioner
	//config.Producer.Partitioner = sarama.NewRandomPartitioner
	//config.Producer.Partitioner = sarama.NewManualPartitioner
	//config.Producer.Partitioner = sarama.NewCustomPartitioner
	assert.NoError(t, err)

	for i := 0; i < 100; i++ {
		_, _, err = producer.SendMessage(&sarama.ProducerMessage{
			Topic: "Test Topic",
			//Key: nil,
			Value: sarama.StringEncoder("这是一条信息"),
			Headers: []sarama.RecordHeader{{
				Key:   []byte("KEY"),
				Value: []byte("VALUE"),
			}},
			Metadata: "这是Metadata",
			//Offset:    0,
			//Partition: 0,
			//Timestamp: time.Time{},
		})
	}
}

func TestAsyncProducer(t *testing.T) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	producer, err := sarama.NewAsyncProducer(addr, config)
	assert.NoError(t, err)

	msgs := producer.Input()
	msgs <- &sarama.ProducerMessage{
		Topic: "Test Async Topic",
		Value: sarama.StringEncoder("Async Sarama:这是一条信息"),
		Headers: []sarama.RecordHeader{{
			Key:   []byte("KEY"),
			Value: []byte("VALUE"),
		}},
		Metadata: "这是Metadata",
	}

	select {
	case msg := <-producer.Successes():
		t.Log("消息发送成功:", string(msg.Value.(sarama.StringEncoder)))
	case err := <-producer.Errors():
		t.Log("消息发送失败:", err.Err, err.Msg)
	}
}
