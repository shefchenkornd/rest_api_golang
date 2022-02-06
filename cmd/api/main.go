package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/shefchenkornd/rest_api/internal/app/api"
	"log"
)

var (
	configPath string
)

func init() {
	// Скажем, что наше приложение будет на этапе запуска получать путь до конфиг файла из внешнего мира
	flag.StringVar(&configPath, "path", "./configs/api.toml", "path to config file in .toml format")
	// ./rest_api -help
}

func main() {
	flag.Parse() // здесь парсятся переменные командной строки
	log.Println("Starting app...")


	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config) // Десериализуем значение .toml файла
	if err != nil {
		log.Println("Can't find config file. Using default values:", err)
	}

	app := api.NewApp(config)
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
