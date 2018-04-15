package main

import (
	"flag"
	"log"
	"runtime"

	"github.com/eburyagin/bizone-acc/internal/cfg"
	"github.com/eburyagin/bizone-acc/internal/ds"
	"github.com/eburyagin/bizone-acc/internal/services"
)

var cf = flag.String("c", "./config.json", "конфигурационный файл")

func main() {

	flag.Parse()

	config, err := cfg.Load(*cf)
	if err != nil {
		log.Fatalln("Ошибка загрузки конфигурации!")
	}

	_, err = ds.Connect(config)
	if err != nil {
		log.Fatalln("Ошибка подключения к хранилищу данных!")
	}

	services.StartServices(config)

	runtime.Goexit()

}
