package main

import (
	"flag"
	"log"
	"strconv"

	"net/http"

	"rest-example-go/config"
	"rest-example-go/controller"
	"rest-example-go/repository"

	_ "github.com/lib/pq"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {
	flag.String("config", "config", "config file name")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	configFile := viper.GetString("config")

	viper.SetConfigName(configFile)
	//viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	var configuration config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	log.Printf("database is %s", configuration.Database.ConnectionString)
	log.Printf("port for this application is %d", configuration.Server.Port)

	db, err := sqlx.Connect(configuration.Database.ConnectionType, configuration.Database.ConnectionString)
	if err != nil {
		log.Printf("Error: %s. Failed to connect to DB\n", err)
	}
	defer db.Close()

	ur := repository.NewUserRepository(db)
	user := controller.NewUser(ur)

	r := mux.NewRouter()

	r.HandleFunc("/users", user.FindAll).Methods("GET")
	r.HandleFunc("/users", user.Add).Methods("POST")
	r.HandleFunc("/users/{id}", user.Update).Methods("PUT")
	r.HandleFunc("/users/{id}", user.Delete).Methods("DELETE")
	r.HandleFunc("/users/{id}", user.FindOneByID).Methods("GET")

	if err := http.ListenAndServe(":"+strconv.Itoa(configuration.Server.Port), r); err != nil {
		log.Fatal(err)
	}
}
