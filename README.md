# GORM-Find

wrapper funcs for [gorm]() library with go v1.18's new general feature

## Install

```
$ go get -d github.com/xiaosongfu/gorm-find
```

## API

* `func count(querySeg *gorm.DB) (int64, error)` : select count with given querySeg
* `func one[T any](querySeg *gorm.DB) (*T, error)` : select one row with given querySeg
* `func rows[T any](querySeg *gorm.DB, page *Page) ([]*T, error)` : select multi rows with given querySeg and pageable param
* `func rowsJoin[T any](querySeg *gorm.DB, sortTableName string, page *Page) ([]*T, error)` : select multi rows for join with given querySeg and pageable param

> for more detail please view the source code.
