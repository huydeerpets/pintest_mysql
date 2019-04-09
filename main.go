package main

import (
	_ "github.com/huydeerpets/pintest/conf/inits"
	_ "github.com/huydeerpets/pintest/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
