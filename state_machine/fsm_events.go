package statemachine

import (
	"fmt"

	"github.com/looplab/fsm"
)

const (
	BookingEntity Entity = "booking"

	// BookingCreatedEvent is the event name for trip created
	BookingCreatedEvent   = "booking-created"
	BookingPlacedEvent    = "booking-placed"
	BookingCancelledEvent = "booking-cancelled"
	BookingExpiredEvent   = "booking-expired"
)

// SetEvents sets the events for the finite state machine based on the entity name.
// It returns the list of events and an error if the entity name is invalid.
func (f *FiniteState) SetEvents() (fsm.Events, error) {
	switch f.EntityName {
	case BookingEntity:
		f.Events = fsm.Events{
			{Name: "create-booking", Src: []string{BookingCreatedEvent}, Dst: BookingPlacedEvent},
			{Name: "cancel-booking", Src: []string{BookingPlacedEvent, BookingExpiredEvent}, Dst: BookingCancelledEvent},
		}
	default:
		return nil, fmt.Errorf("invalid entity name : [%v]", f.EntityName)
	}
	return f.Events, nil
}
