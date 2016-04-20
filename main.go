package main

import (
	"flag"
	"log"
)

var (
	config Config
)

func main() {
	log.Print("Croft is ALIVE")

	//// go run */*.go -config=croft_config.yml
	var configFilePath string
	flag.StringVar(&configFilePath, "config", "/non/existent/filez", "the YAML config file")
	flag.Parse()
	/*
		cfgData, err := ioutil.ReadFile(configFilePath)
		if err != nil {
			log.Fatal(err)
		}
	*/
	config = Config{}
	if err := config.ParseYamlFile(configFilePath); err != nil {
		log.Fatal(err)
	}
	/*
		err = yaml.Unmarshal([]byte(cfgData), &config)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
	*/
	/*
		if err = config.Parse(cfgData); err != nil {
			log.Fatal(err)
		}
	*/
	//fmt.Printf("Croft config:\n%+v\n\n", config)

	publisher, err := connectPublisher()
	if err != nil {
		log.Fatalf("Failed to connect publisher: %s", err.Error())
	}

	messages := make(chan interface{})
	go readUDPMessages(config.ListenPortUDP, messages)
	for msg := range messages {
		err = publisher.Publish(msg)
		if err != nil {
			log.Printf("Failed to publish message %#v: %s", msg, err.Error())
		}
	}
}

func connectPublisher() (Publisher, error) {
	publisher, err := ConnectRabbitPublisher()
	if err != nil {
		return nil, err
	}

	err = publisher.Configure()
	if err != nil {
		return nil, err
	}

	return publisher, nil
}
