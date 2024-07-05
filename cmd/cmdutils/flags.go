package cmdutils

import (
	"github.com/spf13/pflag"
	v "github.com/spf13/viper"
)

// GetParamB returns a parameter as a string and a boolean to tell if it is different from the default
func GetParamB(flags *pflag.FlagSet, key string) (string, bool) {
	value, _ := flags.GetString(key)

	// If set on Flags, use it.
	if flags.Changed(key) {
		return value, true
	}

	// If set through viper (env, config), return it.
	if v.IsSet(key) {
		return v.GetString(key), true
	}

	// Otherwise use default value on flags.
	return value, false
}

func GetParam(flags *pflag.FlagSet, key string) string {
	val, _ := GetParamB(flags, key)
	return val
}

// GetIntParamB returns a parameter as int and a boolean to tell if it is different from the default
func GetIntParamB(flags *pflag.FlagSet, key string) (int, bool) {
	value, _ := flags.GetInt(key)

	// If set on Flags, use it.
	if flags.Changed(key) {
		return value, true
	}

	// If set through viper (env, config), return it.
	if v.IsSet(key) {
		return v.GetInt(key), true
	}

	// Otherwise use default value on flags.
	return value, false
}

func GetIntParam(flags *pflag.FlagSet, key string) int {
	val, _ := GetIntParamB(flags, key)
	return val
}

// GetBoolParamB returns a parameter as bool and a boolean to tell if it is different from the default
func GetBoolParamB(flags *pflag.FlagSet, key string) (bool, bool) {
	value, _ := flags.GetBool(key)

	// If set on Flags, use it.
	if flags.Changed(key) {
		return value, true
	}

	// If set through viper (env, config), return it.
	if v.IsSet(key) {
		return v.GetBool(key), true
	}

	// Otherwise use default value on flags.
	return value, false
}

func GetBoolParam(flags *pflag.FlagSet, key string) bool {
	val, _ := GetBoolParamB(flags, key)
	return val
}

// GetFloat64ParamB returns a parameter as float64 and a boolean to tell if it is different from the default
func GetFloat64ParamB(flags *pflag.FlagSet, key string) (float64, bool) {
	value, _ := flags.GetFloat64(key)

	// If set on Flags, use it.
	if flags.Changed(key) {
		return value, true
	}

	// If set through viper (env, config), return it.
	if v.IsSet(key) {
		return v.GetFloat64(key), true
	}

	// Otherwise use default value on flags.
	return value, false
}

func GetFloat64Param(flags *pflag.FlagSet, key string) float64 {
	val, _ := GetFloat64ParamB(flags, key)
	return val
}
