package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mikeleg/gokitexample/customer"
)

func getenvWithDefault(name, defaultValue string) string {
	val := os.Getenv(name)
	if val == "" {
		val = defaultValue
	}
	return val
}

func main() {

	var (
		connectionString = flag.String("conn", getenvWithDefault("DATABASE_URL", "postgres://postgres:postgres@localhost/crm?sslmode=disable"), "PostgreSQL connection string")
		listenAddr       = flag.String("addr", getenvWithDefault("SERVICE_PORT", ":8080"), "HTTP address to listen on")
	)
	r := mux.NewRouter()

	if *connectionString == "" {
		log.Fatalln("Please pass the connection string using the -conn option")
	}

	db, err := sqlx.Connect("postgres", *connectionString)
	if err != nil {
		log.Fatalf("Unable to establish connection: %v", err)
	}
	cs := customer.NewCustomerService(db)

	r.Handle("/api/customers", customer.MakeFetchAllCustomer(cs, nil)).Methods("GET")

	fmt.Println("listening on port ", *listenAddr)
	printAllRoutes(r)

	http.ListenAndServe(*listenAddr, r)
}

func printAllRoutes(r *mux.Router) {
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, _ := route.GetPathTemplate()
		met, _ := route.GetMethods()
		fmt.Println(tpl, met)
		return nil
	})
}
