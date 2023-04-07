package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Profiles []Profile `json:"Profiles"`
	Active   Profile
}

type Profile struct {
	Environment string `json:"Environment"`
	Secrets     struct {
		Foo string `json:"Foo"`
	} `json:"Secrets"`
}

var (
	cfg Config
)

func init() {
	b, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	profile := os.Getenv("PROFILE")
	if len(profile) < 1 {
		log.Fatal("Missing PROFILE")
	}
	for _, p := range cfg.Profiles {
		if p.Environment == profile {
			cfg.Active = p
		}
	}
	if len(cfg.Active.Environment) < 1 {
		log.Fatal("Invalid PROFILE")
	}
}
