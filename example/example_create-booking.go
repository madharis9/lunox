package main

import (
	"fmt"
	"log"

	statemachine "github.com/kikan/fsm/state_machine"
)

func main() {

	// Create a new finite state machine with the specified entity and event.
	// Returns the created state machine and any error encountered during initialization.
	// Create a new FSM with the given initial state.
	// BookingEntity is the entity name for booking state machine for cancel-booking event
	// this flow will used for cancel booking
	sm, err := statemachine.NewFiniteState(statemachine.BookingEntity, statemachine.BookingCreatedEvent)
	if err != nil {
		panic(err)
	}

	fmt.Println("current 1:" + sm.Fsm.Current())

	currentState := sm.Fsm.Current()
	err = sm.ChangeState("create-booking")
	if err != nil {
		log.Panic("err : ", err)
	}
	dest := sm.Fsm.Current()
	if currentState == dest {
		log.Panic("error")
	}

	fmt.Println("Expected change state:" + sm.Fsm.Current())

}
