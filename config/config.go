package config

import (
	"fmt"
	"net"
	"os"
)

type Config struct {
	URL  string
	PORT string
}

func InitConfig() *Config {
	c := &Config{}
	c.URL = os.Getenv("URL")
	c.PORT = os.Getenv("PORT")

	if c.URL == "" {
		c.URL = "localhost"
	}
	if c.PORT == "" {
		port := getPort()
		c.PORT = fmt.Sprintf("%d", port)
	}
	return c
}

func getPort() int {
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:0")

	l, _ := net.ListenTCP("tcp", addr)
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port
}
