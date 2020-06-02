package config

import (
	"fmt"
	"io"
	"io/ioutil"

	"github.com/dimchansky/utfbom"
	"github.com/jeremywohl/flatten"
	"github.com/palantir/stacktrace"
	"muzzammil.xyz/jsonc"
)

type Configuration map[string]string

func (config *Configuration) AddJsonReader(reader io.Reader) error {
	reader = utfbom.SkipOnly(reader)

	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return stacktrace.Propagate(err, "Could not read")
	}

	data, err := jsonToMap(bytes)
	if err != nil {
		return stacktrace.Propagate(err, "Could not parse to json")
	}
	merge(*config, data)

	return nil
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

func merge(m1, m2 map[string]string) {
	for k, v := range m2 {
		m1[k] = v
	}
}
