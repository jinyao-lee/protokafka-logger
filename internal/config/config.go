package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

var ErrConfigNotFound = errors.New("config.yaml file is not found in root folder of project")

type Root struct {
	Topics   []Topic  `yaml:"topics"`
	Consumer Consumer `yaml:"consumer"`
}

type Consumer struct {
	GroupID             string `yaml:"groupID"`
	WithStartTimeSuffix bool   `yaml:"withStartTimeSuffix"`
	ConsumeFromOldest   bool   `yaml:"consumeFromOldest"`
}

type Topic struct {
	Name           string   `yaml:"name"`
	Brokers        []string `yaml:"brokers"`
	SchemaLocation string   `yaml:"schemaLocation"`
	StructName     string   `yaml:"structName"`
}

func ParseConfig(fileDir string) (*Root, error) {
	file, err := os.ReadFile(fileDir)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("failed to open file %s; err=%w", fileDir, err)
		}

		return nil, ErrConfigNotFound
	}

	root := &Root{}
	if err := yaml.Unmarshal(file, root); err != nil {
		return nil, fmt.Errorf("failed to unmarshal yaml file %s; err=%w", fileDir, err)
	}

	root.manipulateGroupName()

	return root, nil
}

func (r *Root) manipulateGroupName() {
	cgName := r.Consumer.GroupID
	if cgName == "" {
		cgName = "unnamed-consumer-group"
	}

	if r.Consumer.WithStartTimeSuffix {
		cgName += "-" + strconv.FormatInt(time.Now().Unix(), 10)
	}
	r.Consumer.GroupID = cgName
}
