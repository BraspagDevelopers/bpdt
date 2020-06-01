package configuration

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jeremywohl/flatten"
	"muzzammil.xyz/jsonc"
)

type ConfigurationBuilder struct {
	jsonFiles []string
}

func New() *ConfigurationBuilder {
	return &ConfigurationBuilder{}
}

func (b *ConfigurationBuilder) AddJsonFile(jsonFile string) *ConfigurationBuilder {
	if b.jsonFiles == nil {
		b.jsonFiles = make([]string, 0)
	}
	b.jsonFiles = append(b.jsonFiles, jsonFile)
	return b
}

var separator = flatten.SeparatorStyle{
	Middle: "__",
}

func jsonToMap(bytes []byte) (result map[string]string, err error) {
	var data map[string]interface{}
	err = jsonc.Unmarshal(bytes, &data)
	if err != nil {
		return
	}
	data, err = flatten.Flatten(data, "", separator)
	if err != nil {
		return
	}
	result = make(map[string]string)
	for k, v := range data {
		result[k] = fmt.Sprint(v)
	}
	return
}

func (b *ConfigurationBuilder) Build() (result map[string]string, err error) {
	result = make(map[string]string)
	for _, jf := range b.jsonFiles {
		var f *os.File
		f, err = os.Open(jf)
		if err != nil {
			return
		}

		var bytes []byte
		bytes, err = ioutil.ReadAll(f)
		if err != nil {
			return
		}

		var data map[string]string
		data, err = jsonToMap(bytes)
		if err != nil {
			return
		}

		for k, v := range data {
			result[k] = v
		}
	}
	return
}
