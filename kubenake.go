package main

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	os "os"

	yaml "gopkg.in/yaml.v2"
)

// Secret is a struct containing a k8s secret spec
type Secret struct {
	APIVersion string                 `yaml:"apiVersion"`
	Kind       string                 `yaml:"kind"`
	Data       map[string]interface{} `yaml:"data"`
	Type       string                 `yaml:"type"`
	Metadata   map[string]interface{} `yaml:"metadata"`
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readStdin() ([]byte, error) {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Println("Cannot Read Stdin you need to run kubenake after a pipe")
		return []byte{}, err
	}
	return data, err
}

func decode(secret Secret) Secret {
	data := secret.Data
	for key, value := range data {
		decodedValue, err := b64.StdEncoding.DecodeString(value.(string))
		check(err)
		secret.Data[key] = string(decodedValue)
	}
	return secret
}

func main() {

	data, err := readStdin()
	check(err)
	secret := Secret{}
	err = yaml.Unmarshal([]byte(data), &secret)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	decodedSecret := decode(secret)

	decodedYaml, err := yaml.Marshal(&decodedSecret)
	check(err)
	fmt.Println(string(decodedYaml))
}
