package repository

import (
	"fmt"
	"kaya-backend/models"
	"kaya-backend/models/request"
	"time"

	"github.com/jinzhu/gorm"
)

// NewBlogRepository ..
func NewBlogRepository(gen *models.GeneralModel, db *gorm.DB) *blogRepository {
	return &blogRepository{
		General: gen,
		DB:      db,
	}
}

// BlogRepository ..
type (
	BlogRepository interface {
		Filter(request.TagBlog) ([]models.Blog, error)
		FilterBlogTag(req request.TagBlog) ([]models.ResBlogTags, error)
		Detail(slug string) (models.BlogDetail, error)
		WithTrx(*gorm.DB) blogRepository
	}
	blogRepository struct {
		General *models.GeneralModel
		DB      *gorm.DB
	}
)

func (repo blogRepository) WithTrx(trxHandle *gorm.DB) blogRepository {
	fmt.Println(">>> tagRepository - WithTrx <<<")
	defer timeTrack(time.Now(), "tagRepository-WithTrx")
	repo.DB = trxHandle
	return repo
}

func (repo blogRepository) Filter(req request.TagBlog) ([]models.Blog, error) {
	fmt.Println(">>> blogRepository - Filter <<<")
	defer timeTrack(time.Now(), "Filter")

	var res []models.Blog
	db := repo.DB

	if len(req.Id) > 0  {
		db.Where("id in ?", req.Id)
	}

	err := db.Where("status = ?", "PUBLISH").Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repo blogRepository) Detail(slug string) (models.BlogDetail, error){
	fmt.Println(">>> blogRepository - Detail <<<")
	defer timeTrack(time.Now(), "Detail")

	var res models.BlogDetail
	db := repo.DB
	err := db.Where("slug = ?", slug).Where("status = ?", "PUBLISH").Preload("BlogTags").Preload("BlogTags.Tag").Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (repo blogRepository) FilterBlogTag(req request.TagBlog) ([]models.ResBlogTags, error) {
	fmt.Println(">>> blogRepository - Filter <<<")
	defer timeTrack(time.Now(), "Filter")

	var res []models.ResBlogTags
	db := repo.DB
	
	db = db.Table("blog_tags BT").
			Select("B.*," + " BT.tag_id ").
			Joins("LEFT JOIN blogs B on BT.blog_id = B.id")
	
	if len(req.Id) > 0 {
		db = db.Where("BT.tag_id in (?)", req.Id)
	}

	if req.NotId > 0 {
		db = db.Where("B.id != ?", req.NotId)
	}

	err := db.Where("B.status = ?", "PUBLISH").Find(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}
