package inits

import (
	"time"

	_ "github.com/huydeerpets/pintest_mysql/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbname := "default"
	datasource := beego.AppConfig.String("datasource") // sqlite file

	orm.DefaultTimeLoc = time.Local
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase(dbname, "mysql", datasource)

	// sync model and database
	if err := orm.RunSyncdb(dbname, false, true); err != nil {
		panic(err)
	}
}
