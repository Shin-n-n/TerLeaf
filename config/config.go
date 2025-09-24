package config

import (
    "encoding/json"
    "os"
)

type Config struct {
    Theme      string `json:"theme"`
    WindowSize struct {
        Width  int `json:"width"`
        Height int `json:"height"`
    } `json:"window_size"`
}

func LoadConfig(path string) (*Config, error) {
    file, err := os.Open(path)
	if err != nil {
        return nil, err
    }
    defer file.Close()
    var cfg Config
    err = json.NewDecoder(file).Decode(&cfg)
    return &cfg, err
}