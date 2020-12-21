package main

import (
	"context"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	rtr *mux.Router
	ctx = context.Background()
)

type Movie struct {
	Uid      string       `json:"uid,omitempty"`
	Title    string       `json:"title,omitempty"`
	Summary  string       `json:"summary,omitempty"`
	Year     int          `json:"year,omitempty"`
	Released time.Time    `json:"released,omitempty"`
	Genre    []string     `json:"genre,omitempty"`
	Cast     []CastMember `json:"cast,omitempty"`
}

type CastMember struct {
	Uid         string  `json:"uid,omitempty"`
	Name        string  `json:"name,omitempty"`
	Role        string  `json:"role,omitempty"`
	Title       string  `json:"title,omitempty"`
	Filmography []Movie `json:"filmography,omitempty"`
}

type User struct {
	Uid       string `json:"uid,omitempty"`
	Username  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	Friends   []User `json:"friends,omitempty"`
	Favorites struct {
		Movies []Movie      `json:"movies,omitempty"`
		Actors []CastMember `json:"actor,omitempty"`
	} `json:"favorites,omitempty"`
}

func main() {

	rtr = mux.NewRouter()
	rtr.HandleFunc("/init", initDB)

	s := &http.Server{
		Addr:           ":1080",
		Handler:        handlers.LoggingHandler(io.MultiWriter(nil, os.Stdout), rtr),
		MaxHeaderBytes: 1 << 62,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("server failed", err)
	}

}

func newClient() *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial("alpha:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}


func initDB(w http.ResponseWriter, r *http.Request) {

	dgraphClient := newClient()

	op := &api.Operation{}

	op.Schema = `
		title: string @index(exact) .
		summary: string .
		year: int .
		released: datetime .
		genre: [string] .
		Cast: [uid] .

		type Movie {
			title: string
			summary: string 
			year: int
			released: datetime
			genre: [string]
			Cast: [CastMember]
		}
	`

	err := dgraphClient.Alter(ctx, op)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
}
