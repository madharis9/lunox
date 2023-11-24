package statemachine

import (
	"fmt"

	"github.com/looplab/fsm"
)

const (
	Booking Entity = "booking"

	// BookingCreatedEvent is the event name for trip created
	BookingCreatedEvent   = "booking-created"
	BookingPlacedEvent    = "booking-placed"
	BookingCancelledEvent = "booking-cancelled"
	BookingExpiredEvent   = "booking-expired"
)

func (f *FiniteState) SetEvents() (fsm.Events, error) {
	switch f.EntityName {
	case Booking:
		f.Events = fsm.Events{
			{Name: "create-booking", Src: []string{BookingCreatedEvent}, Dst: BookingPlacedEvent},
			{Name: "cancel-booking", Src: []string{BookingPlacedEvent, BookingExpiredEvent}, Dst: BookingCancelledEvent},
		}
	default:
		return nil, fmt.Errorf("invalid entity name : [%v]", f.EntityName)
	}
	return f.Events, nil
}
