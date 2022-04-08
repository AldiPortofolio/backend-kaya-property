package repository

import (
	"fmt"
	"github.com/nleeper/goment"
	"kaya-backend/models"
	"math/rand"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

// NewTransactionRepository ..
func NewTransactionRepository(gen *models.GeneralModel, db *gorm.DB) *transactionRepository {
	return &transactionRepository{
		General: gen,
		DB:      db,
	}
}

// TransactionRepository ..
type (
	TransactionRepository interface {
		Save(models.Transactions) (models.Transactions, error)
		WithTrx(*gorm.DB) transactionRepository
		FindStatusById(statusID uint64) (models.Status, error)
		FindPaymentMethodByName(name string) (models.PaymentMethod, error)
		GetRunNum(name string) (string, error)
		FindByNoOrder(noOrder string) (models.Transactions, error)
		GetRandomNumber() int
	}
	transactionRepository struct {
		General *models.GeneralModel
		DB      *gorm.DB
	}
)

func (repo transactionRepository) WithTrx(trxHandle *gorm.DB) transactionRepository {
	fmt.Println(">>> transactionRepository - WithTrx <<<")
	defer timeTrack(time.Now(), "transactionRepository-WithTrx")
	repo.DB = trxHandle
	return repo
}

func (repo transactionRepository) Save(req models.Transactions) (models.Transactions, error) {
	fmt.Println(">>> transactionRepository - Save <<<")
	defer timeTrack(time.Now(), "Save")

	err := repo.DB.Save(&req).Error
	if err != nil {
		return req, err
	}

	return req, nil
}

func (repo transactionRepository) FindStatusById(statusId uint64) (models.Status, error) {
	fmt.Println(">>> transactionRepository - status <<<")
	defer timeTrack(time.Now(), "status")

	res := models.Status{}

	err := Dbcon.Where("id = ?", statusId).Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repo transactionRepository) FindPaymentMethodByName(name string) (models.PaymentMethod, error) {
	fmt.Println(">>> transactionRepository - status <<<")
	defer timeTrack(time.Now(), "status")

	res := models.PaymentMethod{}

	err := Dbcon.Where("name = ?", name).Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
func (repo transactionRepository) GetRunNum(code string) (string, error) {
	runNums := models.RunNums{}
	err := Dbcon.Where("num_code = ?", code).Find(&runNums).Error
	if err != nil {
		return "", err
	}

	var format = runNums.CodeFormat

	var now, _ = goment.New()
	nowString := now.Format(format)
	fmt.Println("now", nowString)
	var number, _ = strconv.Atoi(runNums.LastNum)
	number = number + 1
	var midle = nowString

	no := fmt.Sprintf("%06d", number)

	fmt.Println("no", no)

	resCode := runNums.Prefix + "-" + midle + "-" + no

	err = Dbcon.Model(&runNums).Update(models.RunNums{
		ID:             runNums.ID,
		Prefix:         runNums.Prefix,
		CodeFormat:     runNums.CodeFormat,
		CodeLen:        runNums.CodeLen,
		LastNum:        strconv.Itoa(number),
		Desc:           runNums.Desc,
		NumCode:        runNums.NumCode,
		LastCodeFormat: midle,
	}).Error
	if err != nil {
		return "", err
	}

	return resCode, err
}

func (repo transactionRepository) FindByNoOrder(noOrder string) (models.Transactions, error) {
	fmt.Println(">>> Database - FindByNoOrder <<<")
	defer timeTrack(time.Now(), "FindByNoOrder")

	res := models.Transactions{}

	db := repo.DB
	if err := db.Where("no_transaction = ?", noOrder).Preload("PaymentMethod").Preload("Status").Preload("TransactionDetail.Property").Find(&res).Error; err != nil {
		return res, err
	}

	return res, nil
}

func (repo transactionRepository) GetRandomNumber() int {
	fmt.Println(">>> Database - Check Random Number <<<")
	defer timeTrack(time.Now(), "CheckRandomNumber")

	var number = 100 + rand.Intn(999-100)

	res := models.RandomNumber{}

	db := repo.DB

	trx := db.Where("random_number = ?", number).Where("status = ?", 1000).Find(&res)
	if trx.Error == nil {
		repo.GetRandomNumber()
	}

	return number
}
