package main

import (
	"github.com/cisordeng/beego/xenon"

	_ "nature/model"
	_ "nature/rest"
)

func main() {
	xenon.Run()
}
