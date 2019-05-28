package main

import (
	"github.com/alecthomas/kingpin"
)

type Config struct {
	port     string
	host   string
}

func (c *Config) WithFlags() *Config {
	kingpin.Flag("port", "webserver http port").StringVar(&c.port)
	kingpin.Flag("host", "backend api url").StringVar(&c.host)
	kingpin.Parse()

	return c
}
