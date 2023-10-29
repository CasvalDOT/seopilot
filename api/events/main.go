package events

type event struct {
	ch chan any
}

type Event interface {
	Listen()
	Dispatch(data any)
}

func (e *event) listen(cb func(any)) {
	for data := range e.ch {
		cb(data)
	}
}

func (e *event) Dispatch(data any) {
	e.ch <- data
}
