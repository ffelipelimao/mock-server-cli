package serialize

import (
	"encoding/json"
	"os"
)

type Serializer struct {
	filePath string
}

func NewSerializer(filePath string) *Serializer {
	return &Serializer{
		filePath: filePath,
	}
}

func (s *Serializer) Parse() (*Config, error) {
	file, err := os.ReadFile(s.filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
