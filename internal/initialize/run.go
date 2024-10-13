package initialize

import (
	_ "github.com/lib/pq"
)
func Run() {
	// 1. Load config
	LoadConfig()
	// 2. Init logger
	InitLogger()

	// 3. Database init
	InitDataBase()

	// 4. Cache (redis) init

	// 5. Run Server
	
}