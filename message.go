package ycmodel

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type MessageQueue struct {
	Messages []Messages `json:"messages"`
}

type EventMetadata struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	CreatedAt time.Time `json:"created_at"`
}

type Attributes struct {
	Senttimestamp string `json:"SentTimestamp"`
}

type MessageAttributes map[string]types.MessageAttributeValue

type Message struct {
	MessageID              string            `json:"message_id"`
	Md5OfBody              string            `json:"md5_of_body"`
	Body                   string            `json:"body"`
	Attributes             Attributes        `json:"attributes"`
	MessageAttributes      MessageAttributes `json:"message_attributes"`
	Md5OfMessageAttributes string            `json:"md5_of_message_attributes"`
}

type Details struct {
	QueueID string  `json:"queue_id"`
	Message Message `json:"message"`
}

type Messages struct {
	EventMetadata EventMetadata `json:"event_metadata"`
	Details       Details       `json:"details"`
}
