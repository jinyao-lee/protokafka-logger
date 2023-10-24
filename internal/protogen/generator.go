package protogen

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"

	"github.com/jinyao-lee/kafka-logger/internal/config"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

const tmpFilePath = "./temp.pb"

type Generator struct{}

func (g Generator) GenerateProtoFile(ctx context.Context, topic config.Topic) (*dynamicpb.Message, error) {
	registry, err := g.createProtoRegistry(ctx, topic.SchemaLocation)
	if err != nil {
		return nil, fmt.Errorf("failed to create proto registry; err=%v", err)
	}

	fileName, err := g.getFileName(topic.SchemaLocation)
	if err != nil {
		return nil, fmt.Errorf("failed to get file name; err=%w", err)
	}

	desc, err := registry.FindFileByPath(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to find file by path; err=%v", err)
	}
	fd := desc.Messages()
	messageDescriptor := fd.ByName(protoreflect.Name(topic.StructName))
	return dynamicpb.NewMessage(messageDescriptor), nil
}

func (g Generator) createProtoRegistry(ctx context.Context, srcFile string) (*protoregistry.Files, error) {
	// Create descriptors using the protoc binary.
	// Imported dependencies are included so that the descriptors are self-contained.
	srcDir := srcFile[:strings.LastIndex(srcFile, "/")]
	fileName, err := g.getFileName(srcFile)
	if err != nil {
		return nil, fmt.Errorf("failed to get file name; err=%w", err)
	}

	inputArgs := []string{
		"--include_source_info",
		"--include_imports",
		fmt.Sprintf("--proto_path=%s/src/", os.Getenv("GOPATH")),
		fmt.Sprintf("--proto_path=%s", srcDir),
		fmt.Sprintf("--descriptor_set_out=%s", tmpFilePath),
		fileName,
	}

	cmdOut, err := exec.CommandContext(ctx, "protoc", inputArgs...).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to run script to generate proto files; cmdOut=%s, err=%w", cmdOut, err)
	}

	defer func() {
		if err := os.Remove(tmpFilePath); err != nil {
			slog.With("error", err).ErrorContext(ctx, "failed to remove tmp file")
		}
	}()

	marshalledDescriptorSet, err := os.ReadFile(tmpFilePath)
	if err != nil {
		return nil, err
	}
	descriptorSet := descriptorpb.FileDescriptorSet{}
	err = proto.Unmarshal(marshalledDescriptorSet, &descriptorSet)
	if err != nil {
		return nil, err
	}

	files, err := protodesc.NewFiles(&descriptorSet)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (g Generator) getFileName(protoFileDir string) (string, error) {
	splitStrings := strings.Split(protoFileDir, "/")
	if len(splitStrings) == 0 {
		return "", fmt.Errorf("invalid protoFileDir %s", protoFileDir)
	}

	return splitStrings[len(splitStrings)-1], nil
}

func NewGenerator() *Generator {
	return &Generator{}
}
