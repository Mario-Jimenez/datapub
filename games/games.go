package games

import (
	"context"
	"encoding/json"
	"os"

	"github.com/juju/errors"
)

type FileHandler struct{}

func NewFileHandler() *FileHandler {
	return &FileHandler{}
}

func (*FileHandler) Find(ctx context.Context) ([]Game, error) {
	file, err := os.ReadFile("./games/games.json")
	if err != nil {
		return nil, errors.New(err.Error())
	}

	var data []Game
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, errors.NotValidf(err.Error())
	}

	return data, nil
}
