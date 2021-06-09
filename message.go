package ycmodel

import (
	"time"
)

type MessageAttributeValue struct {
	DataType string `json:"data_type"`

	// Binary type attributes can store any binary data, such as compressed data,
	// encrypted data, or images.
	BinaryValue *[]byte `json:"binary_value"`

	// Strings are Unicode with UTF-8 binary encoding. For a list of code values, see
	// ASCII Printable Characters
	// (http://en.wikipedia.org/wiki/ASCII#ASCII_printable_characters).
	StringValue *string `json:"string_value"`
}

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

type MessageAttributes map[string]MessageAttributeValue

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
