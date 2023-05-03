package config

import (
	"channels/internal/models"
	"encoding/json"
	"io"
	"log"
	"os"
)

func GetConfig(s string) (*models.Config, *models.DbKey, error) {
	file, err := os.Open("./config/config.json")
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
		return nil, nil, err
	}

	var config models.Config
	var dbKey models.DbKey

	switch s {
	case "db":
		err = json.Unmarshal(bytes, &dbKey)
		if err != nil {
			log.Println(err)
			return nil, nil, err
		}
		return nil, &dbKey, nil

	case "host":
		err = json.Unmarshal(bytes, &config)
		if err != nil {
			log.Println(err)
			return nil, nil, err
		}
		return &config, nil, nil

	default:
		err = json.Unmarshal(bytes, &config)
		if err != nil {
			log.Println(err)
			return nil, nil, err
		}

		err = json.Unmarshal(bytes, &dbKey)
		if err != nil {
			log.Println(err)
			return nil, nil, err
		}

		return &config, &dbKey, nil
	}
}
