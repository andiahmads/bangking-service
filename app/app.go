package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/andiahmads/bangking-service/domain"
	"github.com/andiahmads/bangking-service/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func sanityCheck() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}
}

func Start() {

	//check database
	sanityCheck()
	router := mux.NewRouter()

	dbClient := getDbClient()

	//panggil repo
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	//panggil service
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}

	ah := AccountHandlers{service.NewAccountService(accountRepositoryDb)}
	// accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customer", ch.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	router.HandleFunc("/customer/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

	//if env not found

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	fmt.Println(port, address)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWORD")
	dbAddr := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	client.SetMaxOpenConns(10)
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetConnMaxIdleTime(10)
	return client
}
