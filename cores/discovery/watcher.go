package discovery

import "time"

type Watcher interface {
	Next() (*Result, error)
	Stop()
}

type Result struct {
	Action  EventType
	Service *Service
}

type EventType int

const (
	Create EventType = iota
	Delete
	Update
)

func (t EventType) String() string {
	switch t {
	case Create:
		return "create"
	case Delete:
		return "delete"
	case Update:
		return "update"
	default:
		return "unknown"
	}
}

type Event struct {
	Id        string
	Type      EventType
	Timestamp time.Time
	Service   *Service
}
