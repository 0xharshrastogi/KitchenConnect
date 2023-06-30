package config

import (
	"encoding/json"
	"os"
)

const EMPTY_CONNECTION_STRING string = "KitchenConnectDB"

type AppSetting map[string]any

func (s AppSetting) GetConnectionString(name string) string {
	v, found := s["connection"]
	if !found {
		return EMPTY_CONNECTION_STRING
	}
	c, ok := v.(map[string]any)
	if !ok {
		return EMPTY_CONNECTION_STRING
	}
	conn, ok := c[name].(string)
	if !ok {
		return EMPTY_CONNECTION_STRING
	}
	return conn
}

func BuildAppSetting(fp string) func() AppSetting {
	return func() AppSetting {
		return NewAppSetting(fp)
	}
}

func NewAppSetting(fp string) AppSetting {
	f, err := os.OpenFile(fp, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	as := make(AppSetting)
	if err := json.NewDecoder(f).Decode(&as); err != nil {
		return nil
	}
	return as
}
