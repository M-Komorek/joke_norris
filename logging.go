package main

import (
	"context"
	"log"
	"time"

	"github.com/M-Komorek/joke_norris/dto"
)

type LoggingService struct {
	service Service
}

func NewLoggingService(service Service) Service {
	return &LoggingService{
		service: service,
	}
}

func (loggingService *LoggingService) GetChuckNorrisJoke(ctx context.Context) (chuckNorrisJoke *dto.ChuckNorrisJoke, err error) {
	defer func(start_time time.Time) {
		log.Printf("joke=%+v; err=%v; took=%v\n", chuckNorrisJoke, err, time.Since(start_time))
	}(time.Now())

	return loggingService.service.GetChuckNorrisJoke(ctx)
}
