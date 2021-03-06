package common

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Conf *configuration

//profile variables
type configuration struct {
	SonicHost        string `yaml:"sonic_host"`
	SonicPassword    string `yaml:"sonic_password"`
	SonicCollection  string `yaml:"sonic_collection"`
	SonicBucket      string `yaml:"sonic_bucket"`
	SonicQueryLimit  int    `yaml:"sonic_query_limit"`
	SonicQueryOffset int    `yaml:"sonic_query_offset"`
	SearchPort       string `yaml:"search_port"`
}

func (c *configuration) getConf() *configuration {

	path := flag.String("c", "conf.yaml", "配置文件的路径")

	flag.Parse()

	yamlFile, err := ioutil.ReadFile(*path)
	if err != nil {
		panic(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		panic(err.Error())
	}
	return c
}

func init() {
	conf := configuration{}
	Conf = conf.getConf()
}
