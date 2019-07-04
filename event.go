package event

type Callback func(...interface{})

type EventCb struct {
	cb   Callback
	evl  *eventL
	prev *EventCb
	next *EventCb
}

type eventL struct {
	first *EventCb
	last  *EventCb
}
type Event struct {
	evs map[int]*eventL
}

func NewEvent() *Event {
	return &Event{make(map[int]*eventL)}
}

func (e *Event) On(name int, cb Callback) *EventCb {
	evs := &EventCb{cb, nil, nil, nil}
	if evl, ok := e.evs[name]; ok == false {
		evln := &eventL{evs, evs}
		evs.evl = evln
		e.evs[name] = evln
	} else {
		evs.evl = evl
		if evl.last == nil {
			evl.first = evs
			evl.last = evs
		} else {
			evs.prev = evl.last
			evl.last.next = evs
			evl.last = evs
		}
	}
	return evs
}

func (e *Event) Off(evs *EventCb) {
	evl := evs.evl
	if evs.prev == nil {
		if evs.next == nil {
			evl.first = nil
			evl.last = nil
		} else {
			evl.first = evs.next
			evs.next.prev = nil
		}
	} else if evs.next == nil {
		evl.last = evs.prev
		evl.last.next = nil
		evs.prev = nil
	} else {
		evs.prev.next = evs.next
		evs.next.prev = evs.prev
		evs.evl = nil //make life easier to gc
	}
}

func (e *Event) Fire(name int, l ...interface{}) {
	if evl, ok := e.evs[name]; ok == true {
		evs := evl.first
		for evs != nil {
			evs.cb(l...)
			evs = evs.next
		}
	}
}
