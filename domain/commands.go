package domain

import (
	"github.com/google/uuid"
	eh "github.com/looplab/eventhorizon"
)

const (
	// CreateRestaurantCommandType CreateRestaurantCommand type
	CreateRestaurantCommandType = eh.CommandType("restaurant:create")
	// ChangeRestaurantNameCommandType ChangeRestaurantNameCommand type
	ChangeRestaurantNameCommandType = eh.CommandType("restaurant:change_name")
)

var _ = eh.Command(&CreateRestaurantCommand{})

// CreateRestaurantCommand creates a restaurant
type CreateRestaurantCommand struct {
	ID uuid.UUID `json:"id"`
}

// AggregateID stub
func (c *CreateRestaurantCommand) AggregateID() uuid.UUID { return c.ID }

// AggregateType stub
func (c *CreateRestaurantCommand) AggregateType() eh.AggregateType { return RestaurantAggregateType }

// CommandType stub
func (c *CreateRestaurantCommand) CommandType() eh.CommandType { return CreateRestaurantCommandType }

// ChangeRestaurantNameCommand creates a restaurant
type ChangeRestaurantNameCommand struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// AggregateID stub
func (c *ChangeRestaurantNameCommand) AggregateID() uuid.UUID { return c.ID }

// AggregateType stub
func (c *ChangeRestaurantNameCommand) AggregateType() eh.AggregateType { return RestaurantAggregateType }

// CommandType stub
func (c *ChangeRestaurantNameCommand) CommandType() eh.CommandType {
	return ChangeRestaurantNameCommandType
}
