package service

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/jinyao-lee/kafka-logger/internal/config"
	"github.com/jinyao-lee/kafka-logger/internal/dataaccess"
	"github.com/jinyao-lee/kafka-logger/internal/handler"
	"github.com/jinyao-lee/kafka-logger/internal/protogen"
)

func Start() error {
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rootConfig, err := config.ParseConfig("config.yaml")
	if err != nil {
		return err
	}

	generator := protogen.NewGenerator()

	disposers := make([]func(), 0, len(rootConfig.Topics))
	defer func() {
		for _, dispose := range disposers {
			dispose()
		}
	}()

	for idx := range rootConfig.Topics {
		topic := rootConfig.Topics[idx]
		wg.Add(1)

		generatedFile, err := generator.GenerateProtoFile(ctx, topic)
		if err != nil {
			return err
		}

		messageWriter, dispose, err := dataaccess.NewMessageWriter(topic.Name)
		if err != nil {
			return err
		}
		disposers = append(disposers, dispose)

		h := handler.NewHandler(*generatedFile, *messageWriter)
		c, err := dataaccess.NewSaramaConsumer(topic, rootConfig.Consumer, h.Handle)
		if err != nil {
			return err
		}

		go func(c *dataaccess.SaramaConsumer) {
			defer wg.Done()
			slog.With("topic", topic.Name).InfoContext(ctx, "starting to consume topic")
			if err := c.Consume(ctx); err != nil {
				slog.
					With("error", err).
					With("topic", topic.Name).
					ErrorContext(ctx, "failed to consume topic")
			}
		}(c)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	cancel()

	slog.InfoContext(ctx, "waiting for all goroutines to finish")

	wg.Wait()
	return nil
}
