package main

import (
	"fmt"

	"github.com/longln/go-simplebank/api"
	"github.com/longln/go-simplebank/global"
	db "github.com/longln/go-simplebank/internal/database"
	"github.com/longln/go-simplebank/internal/initialize"
)

func main() {
	initialize.LoadConfig()
	initialize.InitLogger()
	initialize.InitDataBase()

	store := db.NewStore(global.TestDB)
	server := api.NewServer(store)
	address := fmt.Sprintf("%s:%d", global.Config.ServerConfig.Address, global.Config.ServerConfig.Port)
	err := server.Start(address)
	if err != nil {
		panic(err)
	}
}