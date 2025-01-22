package types

import "fmt"

type ForwardRequestToConfig struct {
	BackendServiceRoute map[string]string `yaml:"forwardRequestTo"`
}

func (config *ForwardRequestToConfig) PrintConfig() {
	for key, value := range config.BackendServiceRoute {
		fmt.Println(key, value)
	}
}