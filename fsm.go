package fsm

type StateName string

const STATE_FINAL = StateName("STATE_FINAL")

type State[Context any] struct {
	Name    StateName
	Transit func(ctx *Context) StateName
}

type IFSM interface {
	Run()
	CurrentStateName() StateName
}

type fsm[Context any] struct {
	states           map[StateName]*State[Context]
	ctx              *Context
	currentStateName StateName
}

func NewFSM[Context any](states []*State[Context], ctx *Context) IFSM {
	stateMap := map[StateName]*State[Context]{}

	for _, state := range states {
		stateMap[state.Name] = state
	}

	return &fsm[Context]{
		states:           stateMap,
		ctx:              ctx,
		currentStateName: states[0].Name,
	}
}

func (f *fsm[T]) Run() {
	for f.currentStateName != STATE_FINAL {
		f.currentStateName = f.states[f.currentStateName].Transit(f.ctx)
	}
}

func (f *fsm[T]) CurrentStateName() StateName {
	return f.currentStateName
}
