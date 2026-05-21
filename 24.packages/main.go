package main

import (
	"fmt"

	"github.com/riteshdawale02/podcast/24.packages/auth"

	"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("ritesh", "dawale")
	session := auth.GetSession()

	fmt.Println(session)
	
	color.Red(session)
	color.Green(session)
	color.Yellow(session)

	// powerful command :  go mod tidy
	
}