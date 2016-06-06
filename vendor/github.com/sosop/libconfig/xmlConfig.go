package libconfig

import (
	"encoding/xml"
	"io/ioutil"
)

type XmlConfig struct {
	filename string
	obj      interface{}
}

func NewXmlConfig(filename string, obj interface{}) *XmlConfig {
	xmlConfig := &XmlConfig{filename, obj}
	xmlConfig.parse()
	return xmlConfig
}

func (x *XmlConfig) parse() {
	data, err := ioutil.ReadFile(x.filename)
	if err != nil {
		panic(err)
	}
	err = xml.Unmarshal(data, x.obj)
	if err != nil {
		panic(err)
	}
}
