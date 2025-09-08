package config

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/rotisserie/eris"
	"io/ioutil"
	"log"
	"nso/errs"
	"os"
)

var config *AppConfig

var gameConfig *GameConfig

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		err2 := godotenv.Load(".env")
		if err2 != nil {
			log.Panicln(errs.ToString(eris.Wrap(err, "Error loading .env file")))
		}
	}
	log.Println("Load config file success")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// get current dir of config file
	filePath := os.Getenv("CONFIG_FILE")
	log.Println("Load config file from path: " + filePath)
	file, err := os.Open(filePath)
	if err != nil {
		log.Panicln(errs.ToString(eris.Wrap(err, "Error reading config file")))
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Panicln(errs.ToString(eris.Wrap(err, "Error reading all bytes from file")))
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Panicln(errs.ToString(eris.Wrap(err, "Error parsing config file")))
	}

	// Check if res path is exists
	if _, err := ioutil.ReadDir(config.ResPath); err != nil {
		log.Panicln("Not found res path", eris.Wrap(err, "Not found res path "))
	}

	open, err := os.Open(os.Getenv("APPLICATION_FILE"))
	if err != nil {
		log.Panicln(errs.ToString(eris.Wrap(err, "Error reading application file")))
	}
	defer open.Close()
	data, err = ioutil.ReadAll(open)
	if err != nil {
		log.Panicln(errs.ToString(eris.Wrap(err, "Error reading all bytes from file")))
	}
	err = json.Unmarshal(data, &gameConfig)
	if err != nil {
		log.Panicln(errs.ToString(eris.Wrap(err, "Error parsing application file")))
	}
}
func GetAppConfig() *AppConfig {
	return config
}

func GetGameConfig() *GameConfig {
	return gameConfig
}
