package service

import (
	"github.com/jacklove/go-gin-example/models"
)

type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

func (t *Tag) ExistByName() (bool, error) {
	return models.ExistTagByName(t.Name)
}

//func (t *Tag) Export() (string, error) {
//	tags, err := t.GetAll()
//	if err != nil {
//		return "", err
//	}
//
//	xlsFile := xlsx.NewFile()
//	sheet, err := xlsFile.AddSheet("标签信息")
//	if err != nil {
//		return "", err
//	}
//
//	titles := []string{"ID", "名称", "创建人", "创建时间", "修改人", "修改时间"}
//	row := sheet.AddRow()
//
//	var cell *xlsx.Cell
//	for _, title := range titles {
//		cell = row.AddCell()
//		cell.Value = title
//	}
//
//	for _, v := range tags {
//		values := []string{
//			strconv.Itoa(v.ID),
//			v.Name,
//			v.CreatedBy,
//			strconv.Itoa(v.CreatedOn),
//			v.ModifiedBy,
//			strconv.Itoa(v.ModifiedOn),
//		}
//
//		row = sheet.AddRow()
//		for _, value := range values {
//			cell = row.AddCell()
//			cell.Value = value
//		}
//	}
//
//	time := strconv.Itoa(int(time.Now().Unix()))
//	filename := "tags-" + time + export.EXT
//
//	dirFullPath := export.GetExcelFullPath()
//	err = file.IsNotExistMkDir(dirFullPath)
//	if err != nil {
//		return "", err
//	}
//
//	err = xlsFile.Save(dirFullPath + filename)
//	if err != nil {
//		return "", err
//	}
//
//	return filename, nil
//}
//
//func (t *Tag) GetAll() ([]models.Tag, error) {
//	var (
//		tags, cacheTags []models.Tag
//	)
//
//	cache := cache_service.Tag{
//		State: t.State,
//
//		PageNum:  t.PageNum,
//		PageSize: t.PageSize,
//	}
//	key := cache.GetTagsKey()
//	if gredis.Exists(key) {
//		data, err := gredis.Get(key)
//		if err != nil {
//			logging.Info(err)
//		} else {
//			json.Unmarshal(data, &cacheTags)
//			return cacheTags, nil
//		}
//	}
//
//	tags, err := models.GetTags(t.PageNum, t.PageSize, t.getMaps())
//	if err != nil {
//		return nil, err
//	}
//
//	gredis.Set(key, tags, 3600)
//	return tags, nil
//}