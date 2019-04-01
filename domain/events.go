package domain

import (
	eh "github.com/looplab/eventhorizon"
)

const (
	// Created is the event after a restaurant is created
	Created = eh.EventType("restaurant:created")
	// Deleted is the event after a restaurant is deleted
	Deleted = eh.EventType("restaurant:deleted")
	// NameChanged is the event after a restaurant name is changed
	NameChanged = eh.EventType("restaurant:name_changed")
)

func init() {
	eh.RegisterEventData(NameChanged, func() eh.EventData {
		return &NameChangedData{}
	})
}

// NameChangedData contains the data for NameChanged event
type NameChangedData struct {
	Name string `json:"name" bson:"name"`
}
