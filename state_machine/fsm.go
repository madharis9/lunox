package statemachine

import (
	"context"

	"github.com/looplab/fsm"
)

type State string
type Entity string

type FiniteState struct {
	EntityName       Entity
	InitialState     State
	TransitionChange State
	Events           fsm.Events
	Fsm              *fsm.FSM
}

func NewFiniteState(entity Entity, initial State) (*FiniteState, error) {
	finite := &FiniteState{EntityName: entity, InitialState: initial}
	events, err := finite.SetEvents()
	if err != nil {
		return nil, err
	}
	// Create a new FSM with the given initial state.
	finite.Fsm = fsm.NewFSM(string(initial), events, fsm.Callbacks{})
	return finite, nil
}

func (f *FiniteState) ChangeState(transitionName string) error {
	return f.Fsm.Event(context.Background(), transitionName)
}
