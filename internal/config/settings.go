package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

type SettingsManager interface {
	SetSetting(SettingType, interface{}) error
	GetSetting(SettingType) (Setting, error)
}

var (
	settingsManagerInstance *settingsManager = nil
	once                    sync.Once
)

func GetSettingsManager() SettingsManager {
	once.Do(func() {
		settingsManagerInstance = &settingsManager{
			settingsMap: make(map[SettingType]Setting),
		}
		settingsManagerInstance.initSettingsMap()
	})
	return settingsManagerInstance
}

func (sm *settingsManager) SetSetting(st SettingType, v interface{}) error {
	if st == LastSettingType {
		return fmt.Errorf("Setting %v cannot be used", st)
	}

	sm.Lock()
	defer sm.Unlock()

	setting, _ := sm.settingsMap[st]

	switch setting.ValueType {
	case String:
		if _, ok := v.(string); !ok {
			return fmt.Errorf("String value is expected, got %T", v)
		}
	case Int:
		if _, ok := v.(int); !ok {
			return fmt.Errorf("Int value is expected, got %T", v)
		}
	default:
		return fmt.Errorf("Unexpected type %T", v)
	}

	setting.Value = v
	sm.settingsMap[st] = setting

	return nil
}

func (sm *settingsManager) GetSetting(st SettingType) (Setting, error) {
	sm.RLock()
	defer sm.RUnlock()

	if st == LastSettingType {
		return Setting{}, fmt.Errorf("Setting %v cannot be used", st)
	}

	setting, exists := sm.settingsMap[st]

	if !exists {
		return Setting{}, fmt.Errorf("")
	}

	return setting, nil
}

func (sm *settingsManager) initSettingsMap() {
	var (
		tmpStringVal string
		//tmpIntVal    int
	)

	tmpStringVal = getEnvOrPanic("GIN_SERVER_PORT")
	sm.settingsMap[GinServerPort] = Setting{ValueType: String, Value: tmpStringVal}

	tmpStringVal = getEnvOrPanic("GIN_SERVER_IP")
	sm.settingsMap[GinServerIP] = Setting{ValueType: String, Value: tmpStringVal}

	tmpStringVal = getEnvOrPanic("LOG_FILE_NAME")
	sm.settingsMap[LogFileName] = Setting{ValueType: String, Value: tmpStringVal}

	tmpStringVal = getEnvOrPanic("LOG_LEVEL")
	sm.settingsMap[LogLevel] = Setting{ValueType: String, Value: tmpStringVal}

	// add new above
	if len(sm.settingsMap) != int(LastSettingType) {
		logrus.Panicf("Settings map size and settings count mismatched! Add new setting to settings map")
	}
}

func getEnvOrPanic(v string) string {
	res := os.Getenv(v)

	if len(res) == 0 {
		logrus.Panicf("Couldn't extract %s from env", v)
	}
	return res
}
