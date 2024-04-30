package main

import (
	"github.com/MJU-Capstone-6/devmark-backend/cmd"

	_ "github.com/MJU-Capstone-6/devmark-backend/docs"
)

//	@title			Devmark API
//	@version		1.0
//	@description	Bookmark service with automatic classification.
//	@contact.name	MilkyMilky0116
//	@contact.url	https://milkymilky0116.github.io
//	@contact.email	sjlee990129@gmail.com

// @host		localhost
// @BasePath	/
func main() {
	cmd.Execute()
}
