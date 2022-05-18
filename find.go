package gormfind

import (
	"fmt"

	"gorm.io/gorm"
)

// 查询总记录数
func count(querySeg *gorm.DB) (int64, error) {
	var total int64
	if err := querySeg.Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

// 单表单行查询
func one[T any](querySeg *gorm.DB) (*T, error) {
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

// 单表多行查询
// 1. 范型 T 是查询结果的类型，不是某张表的类型
func rows[T any](querySeg *gorm.DB, page *Page) ([]*T, error) {
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

// 多表联接多行查询
// 1. 范型 T 是查询结果的类型，不是某张表的类型
// 2. sortTableName 是 [排序] 条件所属的表
func rowsJoin[T any](querySeg *gorm.DB, sortTableName string, page *Page) ([]*T, error) {
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
