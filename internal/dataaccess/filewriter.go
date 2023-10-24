package dataaccess

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/dynamicpb"
)

type MessageWriter struct {
	file *os.File
}

type MessageMeta struct {
	Topic     string
	Partition int32
	Offset    int64
}

func (w MessageWriter) Log(meta MessageMeta, message dynamicpb.Message) error {
	jsonBytes, err := protojson.Marshal(&message)
	if err != nil {
		return fmt.Errorf("failed to marshal recognized proto message to JSON; err=%w", err)
	}

	output := fmt.Sprintf("time=%v, partition=%v, offset=%v, message=%v\n", time.Now(), meta.Partition, meta.Offset, string(jsonBytes))
	if _, err := w.file.WriteString(output); err != nil {
		return fmt.Errorf("failed to write message to file; err=%w", err)
	}

	return nil
}

func NewMessageWriter(topicName string) (*MessageWriter, func(), error) {
	if err := os.MkdirAll("log", os.ModePerm); err != nil {
		return nil, nil, fmt.Errorf("failed to create log directory; err=%w", err)
	}

	fileName := "log/" + topicName + ".log"
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open log file; err=%w", err)
	}

	dep := &MessageWriter{file: file}
	dispose := func() {
		if err := dep.file.Close(); err != nil {
			slog.With("error", err).With("file_path", fileName).Error("failed to close log file")
		}
	}
	return dep, dispose, nil
}
