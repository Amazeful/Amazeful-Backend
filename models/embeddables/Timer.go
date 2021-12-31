package embeddables

import "time"

type Timer struct {
	Enabled     bool          `bson:"enabled" json:"enabled"`
	MinMessages int           `bson:"minMessages" json:"minMessages"`
	Interval    time.Duration `bson:"interval" json:"interval"`
	Stream      StreamStatus  `bson:"stream" json:"stream"`
}
