package main

import (
	"fmt"

	"github.com/spf13/viper"

	"assig3/pkg/store/postgres"
	deliveryHttp "assig3/services/contact/internal/delivery/http"
	repositoryStorage "assig3/services/contact/internal/repository/storage/postgres"
	useCaseContact "assig3/services/contact/internal/useCase/contact"
	useCaseGroup "assig3/services/contact/internal/useCase/group"
)

func main() {
	conn, err := postgres.New(postgres.Settings{
		Host:     "localhost",
		Port:     5432,
		Database: "assig",
		User:     "assig",
		Password: "22641037",
	})
	if err != nil {
		panic(err)
	}
	defer conn.Pool.Close()

	var (
		repoStorage  = repositoryStorage.New(conn.Pool, repositoryStorage.Options{})
		ucContact    = useCaseContact.New(repoStorage, useCaseContact.Options{})
		ucGroup      = useCaseGroup.New(repoStorage, useCaseGroup.Options{})
		listenerHttp = deliveryHttp.New(ucContact, ucGroup, deliveryHttp.Options{})
	)

	go func() {
		fmt.Printf("service started successfully on http port: %d", viper.GetUint("HTTP_PORT"))
		if err = listenerHttp.Run(); err != nil {
			panic(err)
		}
	}()

	fmt.Println("Press Enter to exit")
	fmt.Scanln() 
}
