package types

type EventType int

const (
	EventNone EventType = iota

	EventStartRecon

	EventOpenScope

	EventOpenReports

	EventDoctor

	EventSettings

	EventExit
)
