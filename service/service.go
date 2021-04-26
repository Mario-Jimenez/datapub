package service

import (
	"context"

	"github.com/Mario-Jimenez/datapub/broker/kafka"
	"github.com/Mario-Jimenez/datapub/config"
	"github.com/Mario-Jimenez/datapub/data"
	"github.com/Mario-Jimenez/datapub/games"
	"github.com/Mario-Jimenez/datapub/logger"
	"github.com/juju/errors"
	log "github.com/sirupsen/logrus"
)

// Run service
func Run(serviceName, serviceVersion string) {
	// load app configuration
	conf, err := config.NewFileConfig()
	if err != nil {
		if errors.IsNotFound(err) {
			log.WithFields(log.Fields{
				"error": errors.Details(err),
			}).Error("Configuration file not found")
			return
		}
		if errors.IsNotValid(err) {
			log.WithFields(log.Fields{
				"error": errors.Details(err),
			}).Error("Invalid configuration values")
			return
		}
		log.WithFields(log.Fields{
			"error": errors.Details(err),
		}).Error("Failed to retrieve secrets")
		return
	}

	// initialize logger
	logger.InitializeLogger(serviceName, serviceVersion, conf.Values().LogLevel)

	producer := kafka.NewProducer("games", conf.Values().KafkaConnection)
	gamesData := games.NewFileHandler()

	dataHandler := data.NewHandler(gamesData, producer, conf.Values().NumberOfThreads)
	err = dataHandler.PublishMessages(context.Background())
	if err != nil {
		log.WithFields(log.Fields{
			"error": errors.Details(err),
		}).Error("Failed to publish messages")
	}

	if err := producer.Close(); err != nil {
		log.WithFields(log.Fields{
			"error": errors.Details(err),
		}).Error("Failed to close publisher")
	}
}
