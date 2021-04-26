package data

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/Mario-Jimenez/datapub/games"
	"github.com/Mario-Jimenez/datapub/publisher"
	"github.com/juju/errors"
	log "github.com/sirupsen/logrus"
)

type Data interface {
	Find(context.Context) ([]games.Game, error)
}

type Handler struct {
	wg    *sync.WaitGroup
	guard chan struct{}

	data      Data
	publisher publisher.Publisher
}

// NewHandler creates a handler for data management
func NewHandler(data Data, publisher publisher.Publisher, NumberOfThreads int) *Handler {
	return &Handler{
		wg:        &sync.WaitGroup{},
		guard:     make(chan struct{}, NumberOfThreads),
		data:      data,
		publisher: publisher,
	}
}

func (h *Handler) PublishMessages(ctx context.Context) error {
	games, err := h.data.Find(ctx)
	if err != nil {
		return errors.Trace(err)
	}

	for _, game := range games {
		b, err := json.Marshal(game)
		if err != nil {
			log.WithFields(log.Fields{
				"game":  game,
				"error": err.Error(),
			}).Error("json marshal failed")

			continue
		}

		h.guard <- struct{}{}
		h.wg.Add(1)
		go h.publishMessage(ctx, b)
	}

	h.wg.Wait()

	return nil
}

func (h *Handler) publishMessage(ctx context.Context, message []byte) {
	defer func() {
		<-h.guard
		h.wg.Done()
	}()

	// publish a message
	wait := 1
	for {
		if err := h.publisher.Publish(ctx, message); err != nil {
			log.WithFields(log.Fields{
				"message": string(message),
				"wait":    fmt.Sprintf("Retrying in %d second(s)", wait),
				"error":   err.Error(),
			}).Warning("Failed to publish message. Retrying...")
			time.Sleep(time.Duration(wait) * time.Second)
			if wait <= 60 {
				wait += 3
			}
			continue
		}

		log.WithFields(log.Fields{
			"message": string(message),
		}).Debug("Message published successfully")

		return
	}
}
