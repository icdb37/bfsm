package config

import (
	"fmt"
	"os"
)

const (
	envDisableAutoEnv     = "BFSM_CONFIG_DISABLE_AUTO_ENV"
	envConfigFile         = "BFSM_CONFIG_FILE"
	envDisableConfigWatch = "BFSM_CONFIG_DISABLE_WATCH"
)

var ConfigFile string

func MustInitConfig() {
	if err := Init(ConfigFile); err != nil {
		fmt.Printf("init config failed: %v\n", err)
		os.Exit(1)
	}
}

func Init(path string) error {
	if os.Getenv(envDisableAutoEnv) == "" {
		vp.AutomaticEnv()
	}

	if path == "" {
		path = os.Getenv(envConfigFile)
	}

	if path == "" {
		return nil
	}

	vp.SetConfigFile(path)
	if err := vp.ReadInConfig(); err != nil {
		return err
	}

	if !vp.GetBool(envDisableConfigWatch) {
		vp.WatchConfig()
	}

	return nil
}
