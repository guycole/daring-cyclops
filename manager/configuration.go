// Copyright 2021 Guy Cole. All rights reserved.
// Use of this source code is governed by a GPL-3 license that can be found in the LICENSE file.
package main

import (
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
}

var configuration Configuration

func init() {
	loadConfig()
}

func loadConfig() {
	file, err := os.Open("configuration.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}

	decoder := json.NewDecoder(file)
	configuration = Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}
