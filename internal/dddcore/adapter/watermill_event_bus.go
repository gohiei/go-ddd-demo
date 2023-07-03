package dddcore

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"

	"cypt/internal/dddcore"
)

var (
	ctx = context.Background()
)

// WatermillEventBus is an implementation of the EventBus interface using Watermill pub/sub library.
type WatermillEventBus struct {
	router *message.Router
	pubsub *gochannel.GoChannel
}

// Post publishes an event to the event bus.
func (eb *WatermillEventBus) Post(e dddcore.Event) {
	jsonData, err := json.Marshal(e)

	if err != nil {
		// Handle error
	}

	msg := message.NewMessage(e.GetID(), jsonData)

	fmt.Println("post: ", string(jsonData))

	err = eb.pubsub.Publish(e.GetName(), msg)

	if err != nil {
		// Handle error
	}
}

// PostAll publishes all the domain events of an aggregate root to the event bus.
func (eb *WatermillEventBus) PostAll(ar dddcore.AggregateRoot) {
	for _, event := range ar.GetDomainEvents() {
		eb.Post(event)
	}
}

// Register registers an event handler with the event bus.
func (eb *WatermillEventBus) Register(h dddcore.EventHandler) {
	eb.router.AddNoPublisherHandler(
		h.Name(),
		h.EventName(),
		eb.pubsub,
		func(msg *message.Message) error {
			h.When(h.EventName(), msg.Payload)
			return nil
		},
	)

	eb.router.RunHandlers(ctx)
}

// Unregister unregisters an event handler from the event bus.
func (eb *WatermillEventBus) Unregister(h dddcore.EventHandler) {
	// Implementation for unregistering an event handler
}

var _ dddcore.EventBus = (*WatermillEventBus)(nil)

// NewWatermillEventBus creates a new instance of WatermillEventBus.
func NewWatermillEventBus() WatermillEventBus {
	logger := watermill.NewStdLogger(false, false)
	router, err := message.NewRouter(message.RouterConfig{}, logger)

	if err != nil {
		panic(err)
	}

	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(
		middleware.CorrelationID,
		middleware.Retry{
			MaxRetries:      3,
			InitialInterval: time.Millisecond * 100,
			Logger:          logger,
		}.Middleware,
		middleware.Recoverer,
	)

	pubsub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)

	go router.Run(ctx)
	<-router.Running()

	return WatermillEventBus{
		router: router,
		pubsub: pubsub,
	}
}
