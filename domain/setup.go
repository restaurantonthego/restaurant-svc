package domain

import (
	"log"

	// "github.com/google/uuid"
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
	"github.com/looplab/eventhorizon/commandhandler/aggregate"
	"github.com/looplab/eventhorizon/commandhandler/bus"

	// "github.com/looplab/eventhorizon/eventhandler/projector"
	// "github.com/looplab/eventhorizon/eventhandler/saga"
	"github.com/restaurantonthego/restaurant-svc/middleware"
)

// Setup configures the domain.
func Setup(
	eventStore eh.EventStore,
	eventBus eh.EventBus,
	commandBus *bus.CommandHandler) {

	// Add a logger as an observer.
	eventBus.AddObserver(eh.MatchAny(), &middleware.Logger{})

	// Create the aggregate repository.
	aggregateStore, err := events.NewAggregateStore(eventStore, eventBus)
	if err != nil {
		log.Fatalf("could not create aggregate store: %s", err)
	}

	// Create the aggregate command handler and register the commands it handles.
	restaurantCommandHandler, err := aggregate.NewCommandHandler(RestaurantAggregateType, aggregateStore)
	if err != nil {
		log.Fatalf("could not create command handler: %s", err)
	}
	commandHandler := eh.UseCommandHandlerMiddleware(restaurantCommandHandler, middleware.LoggingMiddleware)
	commandBus.SetHandler(commandHandler, CreateRestaurantCommandType)
	commandBus.SetHandler(commandHandler, ChangeRestaurantNameCommandType)

	// // Create and register a read model for individual invitations.
	// invitationProjector := projector.NewEventHandler(
	// 	NewInvitationProjector(), invitationRepo)
	// invitationProjector.SetEntityFactory(func() eh.Entity { return &Invitation{} })
	// eventBus.AddHandler(eh.MatchAnyEventOf(
	// 	InviteCreatedEvent,
	// 	InviteAcceptedEvent,
	// 	InviteDeclinedEvent,
	// 	InviteConfirmedEvent,
	// 	InviteDeniedEvent,
	// ), invitationProjector)

	// // Create and register a read model for a guest list.
	// guestListProjector := NewGuestListProjector(guestListRepo, eventID)
	// eventBus.AddHandler(eh.MatchAnyEventOf(
	// 	InviteAcceptedEvent,
	// 	InviteDeclinedEvent,
	// 	InviteConfirmedEvent,
	// 	InviteDeniedEvent,
	// ), guestListProjector)

	// Setup the saga that responds to the accepted guests and limits the total
	// amount of guests, responding with a confirmation or denial.
	// responseSaga := saga.NewEventHandler(NewResponseSaga(2), commandBus)
	// eventBus.AddHandler(eh.MatchEvent(InviteAcceptedEvent), responseSaga)
}
