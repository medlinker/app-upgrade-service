package libconfig

import (
	"encoding/json"
	"io/ioutil"
)

type JsonConfig struct {
	filename string
	obj      interface{}
}

func NewJsonConfig(filename string, obj interface{}) *JsonConfig {
	jsonConfig := &JsonConfig{filename, obj}
	jsonConfig.parse()
	return jsonConfig
}

func (j *JsonConfig) parse() {
	data, err := ioutil.ReadFile(j.filename)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, j.obj)
	if err != nil {
		panic(err)
	}
}
