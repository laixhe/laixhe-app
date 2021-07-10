package dbx

import "gorm.io/gorm"

// Get 查询单条
func Get(dest interface{}, query string, values ...interface{}) error {
	return client.Raw(query, values...).Scan(dest).Error
}

// Select 查询多条
func Select(dest interface{}, query string, values ...interface{}) error {
	return client.Raw(query, values...).Scan(dest).Error
}

// Exec 原生执行 SQL
func Exec(query string, values ...interface{}) error {
	return client.Exec(query, values...).Error
}

// 事务
type transaction struct {
	db *gorm.DB
}

// Begin 开始事务
func Begin() (*transaction, error) {
	tx := &transaction{db: client.Begin()}
	return tx, tx.db.Error
}

// Rollback 回滚事务
func (tx *transaction) Rollback() error {
	return tx.db.Rollback().Error
}

// Commit 提交事务
func (tx *transaction) Commit() error {
	return tx.db.Commit().Error
}

// Get 查询单条
func (tx *transaction) Get(dest interface{}, query string, values ...interface{}) error {
	return tx.db.Raw(query, values...).Scan(dest).Error
}

// Select 查询多条
func (tx *transaction) Select(dest interface{}, query string, values ...interface{}) error {
	return tx.db.Raw(query, values...).Scan(dest).Error
}

// Exec 执行多条
func (tx *transaction) Exec(query string, values ...interface{}) error {
	return tx.db.Exec(query, values...).Error
}
