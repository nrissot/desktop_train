package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Window_y_offset   int
	Only_Clean_Tracks bool
	Tracks_Visible    bool
	Log_Statistics    bool
}

func (c Config) Save() {
	err := os.MkdirAll(".desktop_train", os.ModeDir)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(".desktop_train/config.json")
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
	file, err := os.Open(".desktop_train/config.json")
	if err != nil {
		if os.IsNotExist(err) {
			log.Default().Writer().Write([]byte("config file does not exist, creating a new one\n"))
			c.Window_y_offset = 39
			c.Only_Clean_Tracks = false
			c.Tracks_Visible = true
			c.Log_Statistics = true
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

type Stats struct {
	Train_Counts int
}

func (s Stats) Save() {
	err := os.MkdirAll(".desktop_train", os.ModeDir)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create(".desktop_train/stats.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(s)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Stats) Load() {
	file, err := os.Open(".desktop_train/stats.json")
	if err != nil {
		if os.IsNotExist(err) {
			log.Default().Writer().Write([]byte("stats file does not exist, creating a new one\n"))
			s.Train_Counts = 0
			s.Save()
			return
		}
		log.Fatal(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(s)
	if err != nil {
		log.Fatal(err)
	}
}
