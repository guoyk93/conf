package conf

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	Loaders = []Loader{
		&jsonLoader{},
		&yamlLoader{},
	}
)

type Loader interface {
	Extensions() []string
	Unmarshal(buf []byte, out interface{}) error
}

func Load(name string, out interface{}) (err error) {
	dir := os.Getenv("CONF_DIR")
	if dir == "" {
		dir = "conf"
	}
	var buf []byte
	for _, loader := range Loaders {
		for _, ext := range loader.Extensions() {
			if buf, err = ioutil.ReadFile(filepath.Join(dir, name+"."+ext)); err != nil {
				if !os.IsNotExist(err) {
					return
				}
				err = nil
				continue
			}
			err = loader.Unmarshal(buf, out)
			return
		}
	}
	err = os.ErrNotExist
	return
}
