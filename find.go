package gormfind

import (
	"fmt"

	"gorm.io/gorm"
)

// Count 查询总记录数
//
// querySeg := db.Model(&model.Badge{}).Where("`badge_type`=?", badgeType)
// total, err := gormfind.Count(querySeg)
//
func Count(querySeg *gorm.DB) (int64, error) {
	var total int64
	if err := querySeg.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

// Row 单表单行查询
//
// case1: 查询 Badge 表的全部字段
// querySeg := db.Where("`id`=?", id)
// row, err := gormfind.One[model.Badge](querySeg)
//
// case2: 查询 Badge 表的部分字段
// querySeg := db.Model(&model.Badge{}).Select("name,tag").Where("`id`=?", id)
// row, err := gormfind.One[model.BadgeSlim](querySeg)
//
func Row[T any](querySeg *gorm.DB) (*T, error) {
	var d T
	tx := querySeg.First(&d)

	if tx.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &d, nil
}

// Rows 单表多行查询
// 1. 范型 T 可以是查询结果的类型，或某张表的类型
//
// case1: 查询 Badge 表的全部字段
// querySeg := db.Where("`badge_type`=?", badgeType)
// rows, err := gormfind.Rows[model.Badge](querySeg)
//
// case2: 查询 Badge 表的部分字段
// querySeg := db.Model(&model.Badge{}).Select("name,tag").Where("`badge_type`=?", badgeType)
// rows, err := gormfind.Rows[model.BadgeSlim](querySeg)
//
func Rows[T any](querySeg *gorm.DB, page *Page) ([]*T, error) {
	if page != nil {
		// 拼接 ORDER BY
		if page.SortField != nil && page.Order != nil {
			querySeg.Order(fmt.Sprintf("`%s` %s", *page.SortField, *page.Order))
		}

		// 拼接 LIMIT
		querySeg.Offset(page.Size * (page.Page - 1)).Limit(page.Size)
	}

	// 执行查询
	var d []*T
	if err := querySeg.Find(&d).Error; err != nil {
		return nil, err
	}

	return d, nil
}

// RowsJoin 多表联接多行查询
// 1. 范型 T 可以是查询结果的类型，或某张表的类型。通常 join 查询的结果会单独定义一个 struct
// 2. sortTableName 是 [排序] 条件所属的表
//
// rowsQuerySeg := db.Where("like_dao.`chain_id` = ?", chainID)
// rowsQuerySeg.Select("dao.name,like_dao.address").Joins("left join dao on like_dao.chain_id = dao.chain_id")
// rows, err := gormfind.RowsJoin[model.DaoSlim](rowsQuerySeg, "dao", convertPage(page))
//
func RowsJoin[T any](querySeg *gorm.DB, sortTableName string, page *Page) ([]*T, error) {
	if page != nil {
		// 拼接 ORDER BY
		if page.SortField != nil && page.Order != nil {
			querySeg.Order(fmt.Sprintf("%s.`%s` %s", sortTableName, *page.SortField, *page.Order))
		} // 和 rows[T]() 不同的地方就在这里 !!!

		// 拼接 LIMIT
		querySeg.Offset(page.Size * (page.Page - 1)).Limit(page.Size)
	}

	// 执行查询
	var d []*T
	if err := querySeg.Find(&d).Error; err != nil {
		return nil, err
	}

	return d, nil
}
