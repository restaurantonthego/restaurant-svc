package domain

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	eh "github.com/looplab/eventhorizon"
	"github.com/looplab/eventhorizon/aggregatestore/events"
)

func init() {
	eh.RegisterAggregate(func(id uuid.UUID) eh.Aggregate {
		return &RestaurantAggregate{
			AggregateBase: events.NewAggregateBase(RestaurantAggregateType, id),
		}
	})
}

// RestaurantAggregateType is the type for the restaurant aggregate
const RestaurantAggregateType = eh.AggregateType("restaurant")

// RestaurantAggregate is the aggregate for restaurant
type RestaurantAggregate struct {
	*events.AggregateBase
	name              string
	created           bool
	numberNameChanges int
}

// TimeNow is a mockable version of time.Now.
var TimeNow = time.Now

// HandleCommand implements the HandleCommand method of the
// eventhorizon.CommandHandler interface.
func (a *RestaurantAggregate) HandleCommand(ctx context.Context, cmd eh.Command) error {
	switch cmd.(type) {
	case *CreateRestaurantCommand:
		// An aggregate can only be created once.
		if a.created {
			return errors.New("already created")
		}
	default:
		// All other events require the aggregate to be created.
		if !a.created {
			return errors.New("not created")
		}
	}

	switch cmd := cmd.(type) {
	case *CreateRestaurantCommand:
		a.StoreEvent(Created, nil, TimeNow())
	case *ChangeRestaurantNameCommand:
		tooManyNameChanges := a.numberNameChanges > 2
		if tooManyNameChanges {
			return errors.New("Number of name changes exceeded")
		}
		a.StoreEvent(NameChanged, &NameChangedData{
			Name: cmd.Name,
		}, TimeNow())
	default:
		return fmt.Errorf("could not handle command: %s", cmd.CommandType())
	}
	return nil
}

// ApplyEvent implements the ApplyEvent method of the
// eventhorizon.Aggregate interface.
func (a *RestaurantAggregate) ApplyEvent(ctx context.Context, event eh.Event) error {
	switch event.EventType() {
	case Created:
		a.created = true
	case NameChanged:
		data, ok := event.Data().(*NameChangedData)
		if !ok {
			return errors.New("invalid data for NameChangedData")
		}
		a.name = data.Name
		a.numberNameChanges++
	default:
		return fmt.Errorf("could not apply the event %s", event.EventType())
	}
	return nil
}
