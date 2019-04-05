package main

import (
	"github.com/cisordeng/beego/xenon"

	_ "gravity/models"
	_ "gravity/rest"
)

func main() {
	xenon.Run()
}
