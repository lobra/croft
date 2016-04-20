package main

// Config manages Croft configuration
type Config struct {
	ListenPortUDP    int    `yaml:"listen_port_udp"`
	ListenAddressUDP string `yaml:"listen_address_udp"`
	NetworkUDP       string `yaml:"network_udp"`
	AMQPUri          string `yaml:"amqp_uri"`
}

/*
func (c Config) Parsezz(data []byte) error {
	c.OrionBaseURL = "http://OrionBaseURL"
	//err := yaml.Unmarshal(data, &c)
	//fmt.Printf("- - - %s\n", data)
	fmt.Printf("- - - %+v\n", c)
	return nil
}
*/
