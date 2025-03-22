package config

import "sync"

type SettingValueType int8
type SettingType int

const (
	String SettingValueType = iota
	Int
)

const (
	GinServerPort SettingType = iota
	GinServerIP
	LogLevel
	LogFileName

	// add new here
	LastSettingType // cannot be used
)

type settingsManager struct {
	settingsMap map[SettingType]Setting
	sync.RWMutex
}

type Setting struct {
	ValueType SettingValueType
	Value     interface{}
}
