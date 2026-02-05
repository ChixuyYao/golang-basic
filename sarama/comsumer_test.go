package sarama

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
)

type ConsumerHandler struct {
}

func (c ConsumerHandler) Setup(session sarama.ConsumerGroupSession) error {
	log.Println("SetUp")
	return nil
}

func (c ConsumerHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	log.Println("CleanUp")
	return nil
}

func (c ConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	msgs := claim.Messages()
	const batchSize = 10 // 每个批次10条数据
	for {
		log.Println("开始一个新的批次")

		batch := make([]*sarama.ConsumerMessage, 0, batchSize)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		var done = false
		var eg errgroup.Group
		for i := 0; i < batchSize && !done; i++ {
			select {
			case <-ctx.Done():
				done = true // 超时
			case msg, ok := <-msgs:
				if ok {
					cancel()
					return nil
				}
				batch = append(batch, msg)
				eg.Go(func() error {
					// 并发处理
					log.Println(string(msg.Value))
					return nil
				})
			}
		}
		cancel()

		err := eg.Wait()
		if err != nil {
			log.Println(err)
			continue
		}

		// 凑够一批数据就可以处理了
		for _, msg := range batch {
			session.MarkMessage(msg, "")
		}
	}

}

func TestConsumer(t *testing.T) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumerGroup(addr, "DEMO", config)
	assert.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()
	start := time.Now()
	err = consumer.Consume(ctx,
		[]string{"Test Async Topic"}, ConsumerHandler{})
	assert.NoError(t, err)

	t.Log(time.Since(start).String())
}
