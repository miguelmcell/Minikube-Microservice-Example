package main

import (
	"log"
	"os"
	"fmt"
	"net/http"
	"net/url"
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/go-chi/chi/middleware"
	_ "github.com/denisenkom/go-mssqldb"
	//"github.com/miguelmcell/Pre-merge-Gating-JenkinsX/app/backend/src/todo"
)

type Player struct {
	Score int `json:"score"`
	Name string `json:"name"`
}


func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)
	router.Get("/api", getPlayers)
	//router.Route("/v1", func(r chi.Router) {
	//	r.Mount("/api", todo.Routes())
	//})
	return router
}

func getPlayers(w http.ResponseWriter, r *http.Request) {
	var score int

	query := url.Values{}
	query.Add("app name", "MyAppName")

	 u := &url.URL{
	     Scheme:   "sqlserver",
	      User:     url.UserPassword(os.Getenv("DBUSER"), os.Getenv("DBPASS")),
	      Host:     fmt.Sprintf("%s:%d",os.Getenv("DBIP"), 32291),
	     // Path:  instance, // if connecting to an instance instead of a port
	    RawQuery: query.Encode(),
	}
	db, err := sql.Open("sqlserver", u.String())

	//db, err := sql.Open("mysql", "SA:<YourStrong!Passw0rd>@tcp(192.168.99.100:32291)/master")
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("select score from Players")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&score)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(score)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	players := []Player{
		{
			Score: score,
			Name: "Migs",
		},
	}
	render.JSON(w, r, players)
}
func main() {
	router := Routes()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ... func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err %s\n", err.Error())
	}
	log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVERPORT"), router))
}
