package configuration

import (
	"os"
	"strconv"
	"time"
	"unsafe"
)

type signed interface {
	int | int8 | int16 | int32 | int64
}

type unsigned interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type floating interface {
	float32 | float64
}

func ReadString(name string, defaultValue string) *string {
	value, exists := os.LookupEnv(name)
	if !exists {
		return &defaultValue
	}

	return &value
}

func ReadUnsigned[T unsigned](name string, defaultValue T) *T {
	stringValue, exists := os.LookupEnv(name)
	if !exists {
		return &defaultValue
	}

	parsedValue, parseErr := strconv.ParseUint(stringValue, 10, int(unsafe.Sizeof(T(0)))*8)
	if parseErr != nil {
		return &defaultValue
	}

	value := T(parsedValue)

	return &value
}

func ReadSigned[T signed](name string, defaultValue T) *T {
	stringValue, exists := os.LookupEnv(name)
	if !exists {
		return &defaultValue
	}

	parsedValue, parseErr := strconv.ParseInt(stringValue, 10, int(unsafe.Sizeof(T(0)))*8)
	if parseErr != nil {
		return &defaultValue
	}

	value := T(parsedValue)

	return &value
}

func ReadFloating[T floating](name string, defaultValue T) *T {
	stringValue, exists := os.LookupEnv(name)
	if !exists {
		return &defaultValue
	}

	parsedValue, parseErr := strconv.ParseFloat(stringValue, int(unsafe.Sizeof(T(0)))*8)
	if parseErr != nil {
		return &defaultValue
	}

	value := T(parsedValue)

	return &value
}

func ReadBool(name string, defaultValue bool) *bool {
	stringValue, exists := os.LookupEnv(name)
	if !exists {
		return &defaultValue
	}

	parsedValue, parseErr := strconv.ParseBool(stringValue)
	if parseErr != nil {
		return &defaultValue
	}

	return &parsedValue
}

func ReadDuration(name string, defaultValue time.Duration, resolution time.Duration) *time.Duration {
	value := ReadUint64(name, 0)
	if value == nil {
		return &defaultValue
	}

	duration := time.Duration(*value) * resolution

	return &duration
}

var (
	ReadUint   = ReadUnsigned[uint]
	ReadUint8  = ReadUnsigned[uint8]
	ReadUint16 = ReadUnsigned[uint16]
	ReadUint32 = ReadUnsigned[uint32]
	ReadUint64 = ReadUnsigned[uint64]

	ReadInt   = ReadSigned[int]
	ReadInt8  = ReadSigned[int8]
	ReadInt16 = ReadSigned[int16]
	ReadInt32 = ReadSigned[int32]
	ReadInt64 = ReadSigned[int64]

	ReadFloat32 = ReadFloating[float32]
	ReadFloat64 = ReadFloating[float64]

	ReadDurationInSeconds = func(name string, defaultValue time.Duration) *time.Duration {
		return ReadDuration(name, defaultValue, time.Second)
	}
)
