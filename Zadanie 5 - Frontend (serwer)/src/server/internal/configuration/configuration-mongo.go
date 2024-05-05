package configuration

import "time"

const (
	mongoDBHostEnv              = "MONGODB_HOST"
	mongoDBNameEnv              = "MONGODB_NAME"
	mongoDBUserEnv              = "MONGODB_USERNAME"
	mongoDBPassEnv              = "MONGODB_PASSWORD"
	mongoDBParametersEnv        = "MONGODB_PARAMETERS"
	mongoDBEnableDNSSeedlistEnv = "MONGODB_ENABLE_DNS_SEEDLIST"
	mongoDBConnectionTimeoutEnv = "MONGODB_CONNECTION_TIMEOUT"
	mongoDBSessionPoolLimitEnv  = "MONGODB_SESSION_POOL_LIMIT"
	mongoDBMaxIdleTimeEnv       = "MONGODB_MAX_IDLE_TIME"
)

const (
	mongoDBHostDefault                     = "localhost:27017"
	mongoDBUserDefault                     = "mongo_user"
	mongoDBPassDefault                     = "mongo_pass"
	mongoDBNameDefault                     = "Backend"
	mongoDBParametersDefault               = ""
	mongoDBEnableDNSSeedlistDefault        = false
	mongoDBConnectionTimeoutDefault        = 10 * time.Second
	mongoDBSessionPoolLimitDefault  uint64 = 32
	mongoDBMaxIdleTimeDefault              = 10 * time.Second
)

// MongoDBConfiguration ...
type MongoDBConfiguration struct {
	Host              *string
	Name              *string
	User              *string
	Pass              *string
	Params            *string
	EnableDNSSeedlist *bool
	ConnectionTimeout *time.Duration
	SessionPoolLimit  *uint64
	MaxIdleTime       *time.Duration
}

// NewMongoDBConfiguration ...
func NewMongoDBConfiguration() *MongoDBConfiguration {
	return &MongoDBConfiguration{
		Host:              ReadString(mongoDBHostEnv, mongoDBHostDefault),
		Name:              ReadString(mongoDBNameEnv, mongoDBNameDefault),
		User:              ReadString(mongoDBUserEnv, mongoDBUserDefault),
		Pass:              ReadString(mongoDBPassEnv, mongoDBPassDefault),
		Params:            ReadString(mongoDBParametersEnv, mongoDBParametersDefault),
		EnableDNSSeedlist: ReadBool(mongoDBEnableDNSSeedlistEnv, mongoDBEnableDNSSeedlistDefault),
		ConnectionTimeout: ReadDurationInSeconds(mongoDBConnectionTimeoutEnv, mongoDBConnectionTimeoutDefault),
		SessionPoolLimit:  ReadUint64(mongoDBSessionPoolLimitEnv, mongoDBSessionPoolLimitDefault),
		MaxIdleTime:       ReadDurationInSeconds(mongoDBMaxIdleTimeEnv, mongoDBMaxIdleTimeDefault),
	}
}

// GetMongoDBHost ...
func (c *Configuration) GetMongoDBHost() string {
	return c.getString(c.mongoDB.Host, mongoDBHostDefault)
}

// GetMongoDBUser ...
func (c *Configuration) GetMongoDBUser() string {
	return c.getString(c.mongoDB.User, mongoDBUserDefault)
}

// GetMongoDBPass ...
func (c *Configuration) GetMongoDBPass() string {
	return c.getString(c.mongoDB.Pass, mongoDBPassDefault)
}

// GetMongoDBName ...
func (c *Configuration) GetMongoDBName() string {
	return c.getString(c.mongoDB.Name, mongoDBNameDefault)
}

// GetMongoDBParams ...
func (c *Configuration) GetMongoDBParams() string {
	return c.getString(c.mongoDB.Params, mongoDBParametersDefault)
}

// GetMongoDBEnableDNSSeedlist ...
func (c *Configuration) GetMongoDBEnableDNSSeedlist() bool {
	return c.getBool(c.mongoDB.EnableDNSSeedlist, mongoDBEnableDNSSeedlistDefault)
}

// GetMongoDBConnectionTimeout ...
func (c *Configuration) GetMongoDBConnectionTimeout() time.Duration {
	return c.getDuration(c.mongoDB.ConnectionTimeout, mongoDBConnectionTimeoutDefault)
}

// GetMongoDBSessionPoolLimit ...
func (c *Configuration) GetMongoDBSessionPoolLimit() uint64 {
	return c.getUint64(c.mongoDB.SessionPoolLimit, mongoDBSessionPoolLimitDefault)
}

// GetMongoDBMaxIdleTime ...
func (c *Configuration) GetMongoDBMaxIdleTime() time.Duration {
	return c.getDuration(c.mongoDB.MaxIdleTime, mongoDBMaxIdleTimeDefault)
}
