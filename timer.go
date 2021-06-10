package ycmodel

type TimerMessages struct {
	Messages []TimerMessages `json:"messages"`
}

type TimerDetails struct {
	TriggerID string `json:"trigger_id"`
}
type TimerMessage struct {
	EventMetadata EventMetadata `json:"event_metadata"`
	Details       TimerDetails  `json:"details"`
}
