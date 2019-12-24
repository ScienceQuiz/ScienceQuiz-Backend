// https://github.com/xesina/golang-echo-realworld-example-app
package main

import (
	"github.com/ScienceQuiz-Backend/db"
	"github.com/ScienceQuiz-Backend/handler"
	"github.com/ScienceQuiz-Backend/route"
	"github.com/ScienceQuiz-Backend/service"
)

func main() {
	r := route.New()
	v1 := r.Group("api")

	d := db.New()
	db.AutoMigrate(d)

	u := service.NewUserService(d)
	h := handler.NewHandler(u)

	h.Register(v1)
	r.Logger.Fatal(r.Start(":8080"))
}