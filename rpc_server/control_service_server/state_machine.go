// SPDX-License-Identifier: Apache-2.0
// Copyright Pionix GmbH and Contributors to EVerest

package control_service_server

import (
	"github.com/looplab/fsm"
)

const (
	StateSetup   string = "Setup"
	StateReady   string = "Ready"
	StateRunning string = "Running"
	StateStopped string = "Stopped"
)

const (
	EventReady string = "Ready"
	EventStart string = "Start"
	EventStop  string = "Stop"
)

func NewStateMachine() *fsm.FSM {
	return fsm.NewFSM(
		StateSetup,
		fsm.Events{
			{Name: EventReady, Src: []string{StateSetup}, Dst: StateReady},
			{Name: EventStart, Src: []string{StateReady}, Dst: StateRunning},
			{Name: EventStop, Src: []string{StateRunning}, Dst: StateStopped},
		},
		fsm.Callbacks{},
	)
}
