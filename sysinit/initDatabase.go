package sysinit

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
)

//初始化数据连接
func InitDatabase() {
	//读取配置文件，设置数据库参数
	//数据库类别
	dbType := beego.AppConfig.String("db_type")
	//连接名称
	dbAlias := beego.AppConfig.String(dbType + "::db_alias")
	//数据库名称
	dbName := beego.AppConfig.String(dbType + "::db_name")
	//数据库连接用户名
	dbUser := beego.AppConfig.String(dbType + "::db_user")
	//数据库连接用户名
	dbPwd := beego.AppConfig.String(dbType + "::db_pwd")
	//数据库IP（域名）
	dbHost := beego.AppConfig.String(dbType + "::db_host")
	//数据库端口
	dbPort := beego.AppConfig.String(dbType + "::db_port")
	dbCharset := beego.AppConfig.String(dbType + "::db_charset")
	//数据库时区
	timezone := beego.AppConfig.String(dbType + "::db_timezone")

	dsn := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=" + dbCharset
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}

	orm.RegisterDriver(dbType, orm.DRMySQL)
	orm.RegisterDataBase(dbAlias, dbType, dsn, 30)
	//如果是开发模式，则显示命令信息
	isDev := false
	if beego.AppConfig.String("runmode") == "dev" {
		isDev = true
	}
	//自动建表
	//orm.RunSyncdb("default", false, isDev)
	orm.Debug = isDev
}
