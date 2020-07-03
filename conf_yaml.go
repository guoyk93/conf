package conf

import (
	"gopkg.in/yaml.v2"
)

type yamlLoader struct{}

func (y *yamlLoader) Extensions() []string {
	return []string{"yaml", "yml"}
}

func (y *yamlLoader) Unmarshal(buf []byte, out interface{}) (err error) {
	var m map[string]interface{}
	if err = yaml.Unmarshal(buf, &m); err != nil {
		return
	}
	if err = remapJSON(m, out); err != nil {
		return
	}
	return
}
