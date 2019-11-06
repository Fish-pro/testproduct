package main

import (
	"fmt"
	"github.com/leyle/ginbase/dbandmq"
	"github.com/leyle/ginbase/middleware"
	"github.com/leyle/testproduct/config"
	"github.com/leyle/testproduct/context"
	"github.com/leyle/testproduct/handler"
	"github.com/leyle/userandrole/auth"
	"os"
	. "github.com/leyle/userandrole/api"
)

func main() {
	// 命令行参数
	// 从配置读取参数
	var err error
	conf := config.LoadConf()

	// 数据库连接初始化
	// 初始化数据库 mongodb 和 redis
	ds := dbandmq.NewDs(conf.Mongodb.Host,
		conf.Mongodb.Port,
		conf.Mongodb.User,
		conf.Mongodb.Passwd,
		conf.Mongodb.DbName)
	defer ds.Close()

	rOpt := &dbandmq.RedisOption{
		Host: conf.Redis.Host,
		Port: conf.Redis.Port,
		Passwd: conf.Redis.Passwd,
		DbNum: conf.Redis.Db,
	}
	redisC, err := dbandmq.NewRedisClient(rOpt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	authDs := dbandmq.NewDs(conf.Auth.Mongodb.Host,
		conf.Auth.Mongodb.Port,
		conf.Auth.Mongodb.User,
		conf.Auth.Mongodb.Passwd,
		conf.Auth.Mongodb.DbName)
	defer authDs.Close()
	uo := initUserOption(conf, authDs)

	ctx := &context.Context{
		Ds: ds,
		R: redisC,
		Conf: conf,
		Au: uo,
	}

	// gin middleware 初始化
	e := middleware.SetupGin()
	apiRouter := e.Group("/api")

	// router 配置
	handler.PRouter(ctx, apiRouter.Group(""))

	// 程序启动
	addr := ctx.Conf.Server.GetServerAddr()
	err = e.Run(addr)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func initUserOption(conf *config.Config, ds *dbandmq.Ds) *UserOption {
	var err error

	rOpt := &dbandmq.RedisOption{
		Host: conf.Auth.Redis.Host,
		Port: conf.Auth.Redis.Port,
		Passwd: conf.Auth.Redis.Passwd,
		DbNum: conf.Auth.Redis.Db,
	}

	redisC, err := dbandmq.NewRedisClient(rOpt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	uo := &UserOption{
		Ds:        ds,
		R:         redisC,
	}

	ao := &auth.Option{
		R:   redisC,
		Ds: ds,
	}
	err = ao.InitAuth()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	AuthOption = ao

	return uo
}