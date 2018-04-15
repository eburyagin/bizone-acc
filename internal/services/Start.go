package services

import (
	"log"

	"github.com/eburyagin/bizone-acc/internal/cfg"
)

func StartServices(config *cfg.Config) error {
	log.Println("Запуск сервисов...")
	StartGetAccount_v1(config)
	log.Println("Сервисы запущены.")
	return nil
}
