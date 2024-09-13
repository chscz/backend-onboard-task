package main

import (
	"log"
	"os"

	"github.com/chscz/backend-onboard-task/internal/auth"
	"github.com/chscz/backend-onboard-task/internal/config"
	"github.com/chscz/backend-onboard-task/internal/handler"
	"github.com/chscz/backend-onboard-task/internal/mysql"
	"github.com/chscz/backend-onboard-task/internal/router"
)

func main() {
	cfg, err := config.LoadFromEnv()
	if err != nil {
		panic(err)
	}
	log.Println(cfg)

	db, err := mysql.InitMySQL(cfg.MySQL)
	if err != nil {
		panic(err)
	}

	ua := auth.NewUserAuth(cfg.JWT)
	uh := handler.NewUserHandler(mysql.UserRepo{DB: db}, ua)
	mh := handler.NewPostHandler(mysql.PostRepo{DB: db}, ua)

	r := router.InitRouter(uh, mh, cfg.LocalDebugMode)

	if err = r.Run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
