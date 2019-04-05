package main

import (
	"github.com/cisordeng/beego/xenon"

	_ "gravity/model"
	_ "gravity/rest"
)

func main() {
	xenon.Run()
}
