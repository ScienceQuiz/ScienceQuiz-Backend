// https://github.com/xesina/golang-echo-realworld-example-app
package main

import (
	"github.com/ScienceQuiz-Backend/db"
	"github.com/ScienceQuiz-Backend/handler"
	"github.com/ScienceQuiz-Backend/router"
	"github.com/ScienceQuiz-Backend/service"
)

func main() {
	r := route.New()
	v1 := r.Group("api")

	d, err := db.New()
	if err == nil {
		defer d.Close()
	}
	db.AutoMigrate(d)

	u := service.NewUserService(d)
	q := service.NewQuizService(d)
	h := handler.NewHandler(u, q)

	h.Register(v1)
	r.Logger.Fatal(r.Start(":8000"))
}