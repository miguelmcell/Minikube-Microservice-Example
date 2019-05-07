package main

import (
	"log"
	"context"
	"os"
	"strings"
	"fmt"
	"strconv"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/mongodb/mongo-go-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Player struct {
	Score int `json:"score"`
	Name string `json:"name"`
}


func Routes() *chi.Mux {
	router := chi.NewRouter()
	cors := cors.New(cors.Options{
	        AllowedOrigins:   []string{"*"},
	        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
	        AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
        })
	router.Use(cors.Handler)
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)
	router.Get("/getPlayers", getPlayers)
	router.Post("/postPlayer/{player}", postPlayer)
	router.Post("/deletePlayer/{name}", deletePlayer)
	router.Get("/getStatus", getStatus)
	return router
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("good"))
}

func deletePlayer(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r,"name")
	filter := bson.D{{"name", name}}
	clientOptions := options.Client().ApplyURI("mongodb://main_admin:Janeth1998@mongodb-service.default.svc.cluster.local:27017/admin")
	client, err := mongo.Connect(context.TODO(), clientOptions)

        if err != nil {
                log.Fatal(err)
        }

        // Check the connection
        err = client.Ping(context.TODO(), nil)

        if err != nil {
         log.Fatal(err)
        }

        fmt.Println("Connected to MongoDB!")

        collection := client.Database("test").Collection("Players")

        if err != nil {
            log.Fatal(err)
        }
	deleteResult, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	render.JSON(w, r, deleteResult)

}

func postPlayer(w http.ResponseWriter, r *http.Request) {
	// Add new player to leaderboard
	player := chi.URLParam(r, "player")
	//ex: migs,100
	s := strings.Split(player, "_")
	name, sc := s[0], s[1]
	score, err := strconv.Atoi(sc)
	if err != nil {
	        // handle error
		fmt.Println(err)
		os.Exit(2)
	}
	clientOptions := options.Client().ApplyURI("mongodb://main_admin:Janeth1998@mongodb-service.default.svc.cluster.local:27017/admin")
	newPlayer := Player{
		Score: score,
		Name: name,
	}
	// Connect to MongoDB
        client, err := mongo.Connect(context.TODO(), clientOptions)

        if err != nil {
		log.Fatal(err)
        }

        // Check the connection
        err = client.Ping(context.TODO(), nil)

        if err != nil {
         log.Fatal(err)
        }

        fmt.Println("Connected to MongoDB!")

	collection := client.Database("test").Collection("Players")

        if err != nil {
            log.Fatal(err)
        }
	insertResult, err := collection.InsertOne(context.TODO(), newPlayer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insertResult.InsertedID)
	render.JSON(w, r, newPlayer)
}


func getPlayers(w http.ResponseWriter, r *http.Request) {
	clientOptions := options.Client().ApplyURI("mongodb://main_admin:Janeth1998@mongodb-service.default.svc.cluster.local:27017/admin")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
	log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
	 log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("test").Collection("Players")

	if err != nil {
	    log.Fatal(err)
	}

	var result []*Player
	options := options.FindOptions{}
	options.Sort = bson.D{{"score", -1}}
	limit := int64(5)
	options.Limit = &limit
	cur, err := collection.Find(context.TODO(), bson.D{}, &options)
	if err != nil {
	   log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
	  var elem Player
	  err := cur.Decode(&elem)
	  if err != nil {
	        log.Fatal(err)
	  }

	  result = append(result, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	for _, play := range result{
		fmt.Printf("%d\n",play.Score)
	}
	render.JSON(w, r, result)
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
