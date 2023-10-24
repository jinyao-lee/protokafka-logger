package handler

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/jinyao-lee/kafka-logger/internal/dataaccess"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/dynamicpb"
)

type Handler struct {
	message       dynamicpb.Message
	messageWriter dataaccess.MessageWriter
}

func (h Handler) Handle(ctx context.Context, msg *sarama.ConsumerMessage) error {
	if err := proto.Unmarshal(msg.Value, &h.message); err != nil {
		return fmt.Errorf("message does not match any registered proto message")
	}

	meta := dataaccess.MessageMeta{
		Topic:     msg.Topic,
		Partition: msg.Partition,
		Offset:    msg.Offset,
	}
	if err := h.messageWriter.Log(meta, h.message); err != nil {
		return fmt.Errorf("failed to log message; err=%w", err)
	}

	return nil

}

func NewHandler(message dynamicpb.Message, messageWriter dataaccess.MessageWriter) Handler {
	return Handler{
		message:       message,
		messageWriter: messageWriter,
	}
}
