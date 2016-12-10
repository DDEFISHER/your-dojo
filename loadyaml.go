package main

import "gopkg.in/yaml.v2"
import "log"
import "io/ioutil"

type Questions struct {
	List []Question `yaml:"questions"`
}

type Question map[string]string

//type Questions []Question `yaml:"a"`

func LoadYaml() Questions {

	file, _ := ioutil.ReadFile("data.yaml")

	questions := Questions{}
	err := yaml.Unmarshal(file, &questions)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return questions

}
