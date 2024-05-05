package configuration

import (
	"time"
)

// Configuration ...
type Configuration struct {
	mongoDB *MongoDBConfiguration
}

// NewConfiguration ...
func NewConfiguration() (*Configuration, error) {
	cfg := &Configuration{
		mongoDB: NewMongoDBConfiguration(),
	}

	return cfg, nil
}

func (c *Configuration) getString(value *string, defaultValue string) string {
	return getValueHelper(value, defaultValue)
}

func (c *Configuration) getInt(value *int, defaultValue int) int {
	return getValueHelper(value, defaultValue)
}

func (c *Configuration) getInt64(value *int64, defaultValue int64) int64 {
	return getValueHelper(value, defaultValue)
}

func (c *Configuration) getBool(value *bool, defaultValue bool) bool {
	return getValueHelper(value, defaultValue)
}

func (c *Configuration) getDuration(value *time.Duration, defaultValue time.Duration) time.Duration {
	return getValueHelper(value, defaultValue)
}

func (c *Configuration) getUint64(value *uint64, defaultValue uint64) uint64 {
	return getValueHelper(value, defaultValue)
}

func (c *Configuration) getStringArray(value []string, defaultValue []string) []string {
	if value != nil {
		return value
	}

	return defaultValue
}

func getValueHelper[T any](value *T, defaultValue T) T {
	if value != nil {
		return *value
	}

	return defaultValue
}
