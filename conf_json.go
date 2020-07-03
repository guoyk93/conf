package conf

import "encoding/json"

type jsonLoader struct {
}

func (j *jsonLoader) Extensions() []string {
	return []string{"json"}
}

func (j *jsonLoader) Unmarshal(buf []byte, out interface{}) error {
	return json.Unmarshal(buf, out)
}

func remapJSON(m map[string]interface{}, out interface{}) (err error) {
	var buf []byte
	if buf, err = json.Marshal(m); err != nil {
		return
	}
	if err = json.Unmarshal(buf, out); err != nil {
		return
	}
	return
}
