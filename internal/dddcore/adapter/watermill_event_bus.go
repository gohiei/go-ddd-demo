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

type WatermillEventBus struct {
	router *message.Router
	pubsub *gochannel.GoChannel
}

func (eb *WatermillEventBus) Post(e dddcore.Event) {
	jsonData, err := json.Marshal(e)

	if err != nil {
	}

	msg := message.NewMessage(e.GetID(), jsonData)

	fmt.Println("post: ", string(jsonData))

	err = eb.pubsub.Publish(e.GetName(), msg)

	if err != nil {

	}
}

func (eb *WatermillEventBus) PostAll(ar dddcore.AggregateRoot) {
	for _, event := range ar.GetDomainEvents() {
		eb.Post(event)
	}
}

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

func (eb *WatermillEventBus) Unregister(h dddcore.EventHandler) {

}

var _ dddcore.EventBus = (*WatermillEventBus)(nil)

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

	// @todo err
	go router.Run(ctx)
	<-router.Running()

	return WatermillEventBus{
		router: router,
		pubsub: pubsub,
	}
}
