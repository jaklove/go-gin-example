package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model
	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}


func GetTags(pageNum int,pageSize int,maps interface{})(tags []Tag)  {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{})(count int)  {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?",id).First(&tag)
	if tag.ID > 0{
		return true
	}

	return false
}

func ExistTagByName(name string)(bool,error)  {
	var tag Tag
	err := db.Select("id").Where("name = ?", name).First(&tag).Error
	if err != nil{
		return false,nil
	}
	if tag.ID > 0 {
		return true,nil
	}
	return false,nil
}

func (tag *Tag)BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("created_on",time.Now().Unix())
	return nil
}

func (tag *Tag)BeforeUpdate(scope *gorm.Scope)  {
	scope.SetColumn("modified_on",time.Now().Unix())
}

func AddTag(name string,state int,createdBy string)bool  {
	db.Create(&Tag{
		Name: name,
		State: state,
		CreatedBy: createdBy,
	})

	return true
}

func DelTag(id int)bool  {
	db.Where("id =?",id).Delete(&Tag{})
	return true
}

func UpdateTag(id int,data interface{})bool  {
	db.Model(&Tag{}).Where("id = ?",id).Update(data)
	return true
}