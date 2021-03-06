package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type service struct {
	Name        string   `json:"name"`
	Label       string   `json:"label"`
	Tags        []string `json:"tags"`
	Plan        string   `json:"plan"`
	Credentials struct {
		Url string `json:"url"`
	} `json:"credentials"`
}

func main() {
	var services map[string][]service
	if err := json.Unmarshal([]byte(os.Getenv("VCAP_SERVICES")), &services); err != nil {
		log.Fatal(err)
	}

	service := os.Getenv("PG_BROKER_NAME")
	if service == "" {
		log.Fatal("$PG_BROKER_NAME is not set")
	}

	conn, err := sql.Open("postgres", services[service][0].Credentials.Url)
	if err != nil {
		log.Fatal(err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		var ver string
		conn.QueryRow("SELECT version()").Scan(&ver)

		rw.Write([]byte(ver))
		rw.Write([]byte("\n"))
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
