package repositories

import (
	"desafio-store/models"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	//_ "github.com/lib/pq"
)

func init() {
	verifyDatabaseTables()
}

var (
	once sync.Once
	err  error
	conn *dbConnection
)

type dbConnection struct {
	Pgdb     *gorm.DB
	host     string
	port     string
	user     string
	password string
	dbName   string
	driver   string
}

// PgConnect é o metodo que irá retorna a instacia de conexão com o banco de dados
func (dcon *dbConnection) PgConnect() *dbConnection {
	once.Do(func() {

		dcon.loadDatabaseConfiguration()
		// Struct Singleton
		dcon.Pgdb, err = gorm.Open(dcon.driver, dcon.getUrl())

		if err != nil {
			log.Println("Error PgConnect::", err)
		}

	})

	log.Println("conectado ao banco de dados")
	return dcon
}

func (dcon *dbConnection) loadDatabaseConfiguration() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dcon.host = os.Getenv("DB_HOST")
	dcon.port = os.Getenv("DB_PORT")
	dcon.user = os.Getenv("DB_USER")
	dcon.password = os.Getenv("DB_PASSWORD")
	dcon.dbName = os.Getenv("DB_NAME")
	dcon.driver = os.Getenv("DB_DRIVER")

}

func (dcon dbConnection) getUrl() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dcon.host, dcon.port, dcon.user, dcon.password, dcon.dbName)
}

// verifyDatabaseTables() esse metodo, verifica se já existem no banco as tabelas a serem utilizadas
// pela aplicação, caso não exista ela cria
func verifyDatabaseTables() {
	var connection = dbConnection{}

	conn = connection.PgConnect()
	db := conn.Pgdb

	if err := conn.Pgdb.DB().Ping(); err != nil {

		log.Fatal("Banco de dados não acessivel")
	}

	if !db.HasTable(&models.Account{}) {
		log.Println("Criando tabela Account")
		db.AutoMigrate(&models.Account{})
	}

	if !db.HasTable(&models.Transfer{}) {
		log.Println("Criando tabela Account")
		db.AutoMigrate(&models.Transfer{})
	}

}
