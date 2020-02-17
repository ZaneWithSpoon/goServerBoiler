package db

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbServiceInstance *DbService

type DbService struct {
	running bool
	host string
	port int
	user string
	password string
	dbname string
	Gorm *gorm.DB
}

// GetAPIService
func GetDBService(host string, port int, user string, dbname string) *DbService {
	//if db hasn't been connected yet
	if dbServiceInstance == nil {
		//Get db password from PostCloudBackend/password.txt
		fmt.Println("reading db password from file")
		file, err := os.Open("password.txt")
		if err != nil {
			log.Fatal(err)
			fmt.Println("missing PostCloudBackend/password.txt for db")
		}
		b, err := ioutil.ReadAll(file)

		//Create dbService Object
		dbServiceInstance = &DbService{
			running: false,
			host: host,
			port: port,
			user: user,
			password: string(b),
			dbname: dbname,
		}
	}

	return dbServiceInstance
}

func Start(this *DbService) {
	//configure DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		this.host, this.port, this.user, this.password, this.dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	this.Gorm = db
	this.running = true

	//test db connection
	fmt.Println("connected to DB errors: ")
	fmt.Println(GetDbService().Gorm.DB().Ping())
}

// Close
func (this *DbService) Close() {
	this.Gorm.Close()
}

func GetDbService() *DbService {
	return dbServiceInstance
}