package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Window_y_offset int
}

func (c Config) Save() {
	file, err := os.Create("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(c)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Config) Load() {
	file, err := os.Open("config.json")
	if err != nil {
		if os.IsNotExist(err) {
			log.Default().Writer().Write([]byte("file does not exist, creating a new one\n"))
			c.Window_y_offset = 39
			c.Save()
			return
		}
		log.Fatal(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(c)
	if err != nil {
		log.Fatal(err)
	}
}
