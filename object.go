package ycmodel

type ObjectStorageMessages struct {
	Messages []ObjectMessages `json:"messages"`
}

type ObjectDetails struct {
	BucketID string `json:"bucket_id"`
	ObjectID string `json:"object_id"`
}

type ObjectMessages struct {
	EventMetadata EventMetadata `json:"event_metadata"`
	Details       TimerDetails  `json:"details"`
}
