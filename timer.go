package ycmodel

type Timer struct {
	Messages []Messages `json:"messages"`
}

type TimerDetails struct {
	TriggerID string `json:"trigger_id"`
}
type TimerMessages struct {
	EventMetadata EventMetadata `json:"event_metadata"`
	Details       TimerDetails  `json:"details"`
}
