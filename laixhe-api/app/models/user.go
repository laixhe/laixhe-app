package models

import (
	"github.com/laixhe/goutil/db_sqlx"
)

type User struct {
	Id   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Age  int    `db:"age" json:"age"`
}

func UserList() ([]*User, error) {

	query := "SELECT `id`,`name`,`age` FROM `user`"

	user := make([]*User, 0, 10)

	err := db_sqlx.SQLX().Select(&user, query)
	if err != nil {
		return nil, err
	}

	return user, nil

}
