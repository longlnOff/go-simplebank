package initialize

import (
	"fmt"

	_ "github.com/lib/pq"

	"github.com/longln/go-simplebank/global"
)
func Run() {
	// 1. Load config
	LoadConfig()
	// 2. Init logger
	InitLogger()

	// 3. Database init
	InitDataBase()

	// 4. Cache (redis) init

	// 5. Init Router
	r := InitRouter()
	addressString := fmt.Sprintf("%s:%d", global.Config.ServerConfig.Address, global.Config.ServerConfig.Port)
	r.Run(addressString)
}