package util

import "testing"

func TestLoadConfig(t *testing.T) {
	conf, err := loadConfig("./conf.json")
	if err != nil {
		t.Errorf("sth wrong %s ", err)
	}
	t.Log(err, conf)
}
