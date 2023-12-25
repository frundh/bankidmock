package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rojanDinc/bankidmock/api"
	bolt "go.etcd.io/bbolt"
)

var (
	port = envOrDefault("PORT", "8888")
)

func main() {
	log.Println("intitialising server...")

	db, err := bolt.Open("my.db", 0600, nil)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}
	defer log.Println(db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket([]byte("orders"))
	}))

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("orders"))
		return err
	}); err != nil {
		log.Fatalln(err)
	}

	ctrl := api.NewController(db)
	go func() {
		for {
			if err := api.CleanUp(db); err != nil {
				log.Println(err)
			}
			time.Sleep(5 * time.Second)
		}
	}()

	http.HandleFunc("/rp/v6.0/auth", ctrl.Auth)
	http.HandleFunc("/rp/v6.0/collect", ctrl.Collect)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		os.Exit(0)
	}()

	log.Println("starting server on port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func envOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
