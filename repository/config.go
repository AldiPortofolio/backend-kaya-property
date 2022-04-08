package repository

import (
	"fmt"
	"kaya-backend/models"

	logger "kaya-backend/library/logger/v2"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres is a must
	"github.com/kelseyhightower/envconfig"
)

// DbEnv ..
type DbEnv struct {
	DbUser    string `envconfig:"DB_POSTGRES_USER" default:"postgres"`
	DbPass    string `envconfig:"DB_POSTGRES_PASS" default:"admin123"`
	DbName    string `envconfig:"DB_POSTGRES_NAME" default:"kaya"`
	DbAddres  string `envconfig:"DB_POSTGRES_ADDRESS" default:"8.215.68.35"`
	DbPort    string `envconfig:"DB_POSTGRES_PORT" default:"5432"`
	DbDebug   bool   `envconfig:"DB_POSTGRES_DEBUG" default:"true"`
	DbType    string `envconfig:"DB_POSTGRES_TYPE" default:"postgres"`
	SslMode   string `envconfig:"DB_POSTGRES_SSL_MODE" default:"disable"`
	DbTimeout string `envconfig:"DB_POSTGRES_TIMEOUT" default:"30"`
}

var (
	// Dbcon ..
	Dbcon *gorm.DB

	// Errdb ..
	Errdb error
	dbEnv DbEnv
)

func init() {
	fmt.Println("DB POSTGRES")

	err := envconfig.Process("Database_KAYA", &dbEnv)
	if err != nil {
		fmt.Println("Failed to get Database KAYA env:", err)
	}

	if DbOpen() != nil {
		fmt.Println("Can't Open ", dbEnv.DbName, " DB", DbOpen())
	}
	Dbcon = GetDbCon()
	Dbcon = Dbcon.LogMode(true)
}

// DbOpen ..
func DbOpen() error {
	args := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=%s", dbEnv.DbAddres, dbEnv.DbPort, dbEnv.DbUser, dbEnv.DbPass, dbEnv.DbName, dbEnv.SslMode, dbEnv.DbTimeout)
	Dbcon, Errdb = gorm.Open("postgres", args)
	log := logger.InitLog()
	if Errdb != nil {
		log.Error(fmt.Sprintf("open db Err :%s ", Errdb.Error()))
		return Errdb
	}

	if errping := Dbcon.DB().Ping(); errping != nil {
		log.Error(fmt.Sprintf("Db Not Connect test Ping : %s", errping.Error()))
		fmt.Println("Can't Open db Postgres")
		return errping
	}
	log.Info("Connect DB success")
	return nil
}

// GetDbCon ..
func GetDbCon() *gorm.DB {
	//TODO looping try connection until timeout
	// using channel timeout
	log := logger.InitLog()
	if errping := Dbcon.DB().Ping(); errping != nil {
		log.Error(fmt.Sprintf("Db Not Connect test Ping : %s", errping.Error()))
		//errping = nil
		if errping = DbOpen(); errping != nil {
			log.Error(fmt.Sprintf("try to connect again but error : %s", errping.Error()))
		}
	}
	Dbcon.LogMode(true)
	return Dbcon
}

// DbPostgres ..
type DbPostgres struct {
	General models.GeneralModel
}

// TotalRow ..
type TotalRow struct {
	Total int64 `gorm:"column(total)"`
}

// AsyncRawQuery ..
func AsyncRawQuery(query string, order string, res interface{}, gormchan chan *gorm.DB) {

	sql := Dbcon.Raw(query + order).Scan(res)

	gormchan <- sql
}

func (databae *DbPostgres) GetConfig() *gorm.DB {
	return Dbcon
}
