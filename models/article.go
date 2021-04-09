package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model
	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) {
	scope.SetColumn("created_on", time.Now().Unix())
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) {
	scope.SetColumn("modified_on", time.Now().Unix())
}

func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ?", id).First(&article).Error
	if err != nil {
		return false, nil
	}
	if article.ID > 0 {
		return true,nil
	}

	return false,nil
}

func GetArticleTotal(maps interface{}) int {
	var count int
	db.Model(&Article{}).Count(&count)
	return count
}

func GetArticle(id int) (*Article,error) {
	var article Article
	err := db.Where("id = ?", id).First(&article).Related(&article.Tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article,nil
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Update(data)
	return true
}

func AddArticle(data map[string]interface{}) error {
	article := &Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	}

	if err := db.Create(article).Error; err != nil {
		return err
	}
	return nil
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(&Article{})
	return true
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (article Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&article)
	return article
}
