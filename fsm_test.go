package fsm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type context struct {
	count int
}

type stateWithContext = State[context]

func TestFSM(t *testing.T) {
	t.Run("FSM stores current state name", func(t *testing.T) {
		// given
		state1 := stateWithContext{
			Name:    "state1",
			Transit: func(ctx *context) StateName { return "state2" },
		}

		state2 := stateWithContext{
			Name:    "state2",
			Transit: func(ctx *context) StateName { return STATE_FINAL },
		}

		fsm := NewFSM([]*stateWithContext{&state1, &state2}, &context{})

		//  when
		fsm.Run()

		//  then
		assert.Equal(t, STATE_FINAL, fsm.CurrentStateName())
	})

	t.Run("FSM executes transitions", func(t *testing.T) {
		//  given
		c := context{}

		state1 := stateWithContext{
			Name: "state1",
			Transit: func(ctx *context) StateName {
				ctx.count++
				return "state2"
			},
		}

		state2 := stateWithContext{
			Name: "state2",
			Transit: func(ctx *context) StateName {
				ctx.count += 2
				return STATE_FINAL
			},
		}

		fsm := NewFSM([]*stateWithContext{&state1, &state2}, &c)

		//  when
		fsm.Run()

		// then
		assert.Equal(t, 3, c.count)
	})
}
