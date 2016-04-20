package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config manages Croft configuration
type Config struct {
	ListenPortUDP    int    `yaml:"listen_port_udp"`
	ListenAddressUDP string `yaml:"listen_address_udp"`
	NetworkUDP       string `yaml:"network_udp"`
	AMQPUri          string `yaml:"amqp_uri"`
}

func (c Config) ParseYamlFile(configFilePath string) error {
	cfgData, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("error reading %s: %v", configFilePath, err)
		return err
	}
	err = yaml.Unmarshal(cfgData, &c)
	if err != nil {
		log.Fatalf("error parsing %s: %v", configFilePath, err)
		return err
	}
	log.Printf("Croft config:\n%+v\n\n", config)
	return nil
}

/*
func (c Config) Parse(data []byte) error {
	err := yaml.Unmarshal(data, &c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//err := yaml.Unmarshal(data, &c)
	//fmt.Printf("- - - %s\n", data)
	fmt.Printf("- - - %+v\n", c)
	return nil
}
*/
