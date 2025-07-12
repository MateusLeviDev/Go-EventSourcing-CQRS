package store

import (
	"context"
	"errors"
	"io"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/pkg/es"
	"github.com/MateusLeviDev/Go-EventSourcing-CQRS/pkg/logger"
)

type aggregateStore struct {
	log logger.Logger
	db  *esdb.Client
}

func NewAggregatorStore(log logger.Logger, db *esdb.Client) *aggregateStore {
	return &aggregateStore{log: log, db: db}
}

func (a *aggregateStore) Load(ctx context.Context, streamID string, aggregate es.Aggregate) error {
	stream, err := a.db.ReadStream(ctx, streamID, esdb.ReadStreamOptions{
		Direction: esdb.Forwards,
		From:      esdb.Revision(1),
	}, 100)
	if err != nil {
		return err
	}
	defer stream.Close()

	events := make([]es.Event, 0, 100)
	for {
		event, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return err
		}
		events = append(events, es.NewEventFromRecorded(event.Event))
	}

	return aggregate.Load(events)
}
