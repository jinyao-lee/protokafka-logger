package dataaccess

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/IBM/sarama"
	"github.com/jinyao-lee/kafka-logger/internal/config"
)

var _ sarama.ConsumerGroupHandler = (*SaramaConsumer)(nil)

type SaramaConsumer struct {
	consumerGroup     sarama.ConsumerGroup
	topic             config.Topic
	ready             chan bool
	clientHandlerFunc func(ctx context.Context, message *sarama.ConsumerMessage) error
}

// Cleanup implements sarama.ConsumerGroupHandler.
func (c SaramaConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// Setup implements sarama.ConsumerGroupHandler.
func (c SaramaConsumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

// ConsumeClaim implements sarama.ConsumerGroupHandler.
func (c SaramaConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				slog.InfoContext(session.Context(), "message channel was closed")
				return nil
			}

			session.MarkMessage(message, "")
			ctx := session.Context()
			if err := c.clientHandlerFunc(ctx, message); err != nil {
				slog.With("error", err).ErrorContext(ctx, "failed to handle message")
			}

		case <-session.Context().Done():
			return nil
		}
	}
}

func (c SaramaConsumer) Consume(ctx context.Context) error {
	if err := c.consumerGroup.Consume(ctx, []string{c.topic.Name}, c); err != nil {
		return fmt.Errorf("failed to consume topic %s; err=%w", c.topic.Name, err)
	}

	return nil
}

func NewSaramaConsumer(
	topic config.Topic,
	consumer config.Consumer,
	clientHandlerFunc func(ctx context.Context, message *sarama.ConsumerMessage) error,
) (*SaramaConsumer, error) {
	cfg := sarama.NewConfig()
	if consumer.ConsumeFromOldest {
		cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	cg, err := sarama.NewConsumerGroup(topic.Brokers, consumer.GroupID, cfg)
	if err != nil {
		return nil, err
	}

	return &SaramaConsumer{
		consumerGroup:     cg,
		topic:             topic,
		ready:             make(chan bool),
		clientHandlerFunc: clientHandlerFunc,
	}, nil
}
