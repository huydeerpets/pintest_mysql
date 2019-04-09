package main

import (
	_ "github.com/huydeerpets/pintest_mysql/conf/inits"
	_ "github.com/huydeerpets/pintest_mysql/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
