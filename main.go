package main

import (
	"github.com/nautilus/gateway"
	"github.com/nautilus/graphql"
	"log"
	"net/http"
	"time"
)

func main() {

	//Start the service that accepts uploads
	go StartServiceUpload()

	schemas, err := graphql.IntrospectRemoteSchemas([]string{"http://localhost:5000/query"}...)
	if err != nil {
		log.Fatal(err)
	}

	factory := gateway.QueryerFactory(func(ctx *gateway.PlanningContext, url string) graphql.Queryer {
		//return graphql.NewSingleRequestQueryer(url) //Uncomment for working Queryer
		return graphql.NewMultiOpQueryer(url, 10*time.Millisecond, 1000)
	})

	gw, err := gateway.New(schemas, gateway.WithQueryerFactory(&factory))
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/graphql", gw.GraphQLHandler)

	// start the server
	log.Println("Starting server on port 4040")
	if err = http.ListenAndServe(":4040", nil); err != nil {
		log.Fatal(err.Error())
	}
}
