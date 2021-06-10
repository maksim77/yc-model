package ycmodel

import "time"

type EventMetadata struct {
	EventID        string         `json:"event_id"`
	EventType      string         `json:"event_type"`
	CreatedAt      time.Time      `json:"created_at"`
	CloudID        string         `json:"cloud_id"`
	FolderID       string         `json:"folder_id"`
	TracingContext TracingContext `json:"tracing_context"`
}

type TracingContext struct {
	TraceID       string `json:"trace_id"`
	SpanID        string `json:"span_id"`
	ParrentSpanID string `json:"parent_span_id"`
}
