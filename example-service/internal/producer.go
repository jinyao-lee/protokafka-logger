package internal

import (
	"context"
	"log/slog"

	"github.com/IBM/sarama"
	"github.com/jinyao-lee/kafka-logger/example-service/pb"
	"google.golang.org/protobuf/proto"
)

const topicName = "sample-topic-name"

type SaramaProducer struct {
	producer sarama.SyncProducer
}

func (p SaramaProducer) ProduceMessage(ctx context.Context, input *pb.MySampleKafkaMessage) error {
	b, err := proto.Marshal(input)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: topicName,
		Value: sarama.ByteEncoder(b),
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	slog.InfoContext(ctx, "successfully produced Kafka message", slog.Attr{Key: "partition", Value: slog.AnyValue(partition)}, slog.Attr{Key: "offset", Value: slog.AnyValue(offset)})
	return nil
}

func NewSaramaProducer() (*SaramaProducer, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Errors = true
	cfg.Producer.Return.Successes = true
	cfg.Net.MaxOpenRequests = 1

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, cfg)
	if err != nil {
		return nil, err
	}

	return &SaramaProducer{
		producer: producer,
	}, nil
}
