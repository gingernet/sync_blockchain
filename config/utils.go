package config

import (
	"time"
	"github.com/spf13/viper"
)

// getInt
func getInt(key string, defaultValue int) int {
	var (
		value int
	)
	if value = viper.GetInt(key); value == 0 {
		return defaultValue
	}
	return value
}

// getInt64
func getInt64(key string, defaultValue int64) int64 {
	var (
		value int64
	)
	if value = viper.GetInt64(key); value == 0 {
		return defaultValue
	}
	return value
}

// getFloat64
func getFloat64(key string, defaultValue float64) float64 {
	var (
		value float64
	)
	if value = viper.GetFloat64(key); value == 0 {
		return defaultValue
	}
	return value
}

// getString
func getString(key string, defaultValue string) string {
	var (
		value string
	)
	if value = viper.GetString(key); value == "" {
		return defaultValue
	}
	return value
}

// getStringSlice
func getStringSlice(key string, defaultValue []string) []string {
	var (
		value []string
	)
	if value = viper.GetStringSlice(key); len(value) == 0 {
		return defaultValue
	}
	return value
}

// getDuration
func getDuration(key string, defaultValue time.Duration) time.Duration {
	var (
		value string
	)
	if value = viper.GetString(key); value == "" {
		return defaultValue
	}
	if duration, err := time.ParseDuration(value); err == nil {
		return duration
	}
	return defaultValue
}

// getbool
func getbool(key string, defaultValue bool) bool {
	var (
		value bool
	)
	if value = viper.GetBool(key); value == false {
		return defaultValue
	}
	return value
}


