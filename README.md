[![codecov](https://codecov.io/gh/DAtek/fsm/graph/badge.svg?token=LrWVnbgjTj)](https://codecov.io/gh/DAtek/fsm) [![Go Report Card](https://goreportcard.com/badge/github.com/DAtek/fsm)](https://goreportcard.com/report/github.com/DAtek/fsm)

# FSM

A dead simple Finite State Machine in Golang.

## Why?
I looked at a couple of other libs on GitHub, but all of them were more complicated. Events, sources, destinations, callbacks, background context, whatever... I don't need those.  

A finite state machine can be defined only with _states_, if we assign each state a single _transition_, which returns the next state.

## Usage
Look at the tests.
