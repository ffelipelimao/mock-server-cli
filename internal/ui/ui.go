package ui

import (
	"fmt"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/input"
)

func GetFilePathUI() (string, error) {
	filePath, err := prompt.New().Ask("Type the path for your response file: ").Input(
		"Must be .json or .xml to RESTful response",
		input.WithHelp(true),
	)
	if err != nil {
		return "", fmt.Errorf("error in read config file %v", err)
	}
	return filePath, nil
}
