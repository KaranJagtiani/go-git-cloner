package utils

import (
	"io/ioutil"

	"github.com/KaranJagtiani/go-git-cloner/types"

	"gopkg.in/yaml.v2"
)

func ParseYaml(file string) (*types.Config, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	c := &types.Config{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
