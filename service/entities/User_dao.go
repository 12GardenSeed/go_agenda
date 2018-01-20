package entities

import (
	"crypto/md5"
	"fmt"
	"io"
)

type userDao DataAccessObject

func (dao *userDao) Insert(user *User) error {
	hash := md5.New()
	io.WriteString(hash, user.Username)
	digest := fmt.Sprintf("&x", hash.Sum(nil))
	sqlStmt := `
        INSERT INTO users(key, username, password, email, phone) VALUES(
            '` + digest + `',
            '` + user.Username + `',
            '` + user.Password + `',
            '` + user.Email + `',
            '` + user.Phone + `'
        );`
	result, err := db.Exec(sqlStmt)
	CheckErr(err)
	id, err := result.LastInsertId()
	CheckErr(err)
	user.ID = int(id)
	user.Key = digest
	return nil
}

func (dao *userDao) FindAll() ([]User, error) {
	sqlStmt := `SELECT * FROM users`

	rows, err := dao.Query(sqlStmt)
https: //github.com/go-sql-driver/go-sqlite3
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	userList := make([]User, 0, 0)
	for rows.Next() {
		users := User{}
		err := rows.Scan(&users.ID, &users.Key, &users.Username, &users.Password, &users.Email, &users.Phone)
		CheckErr(err)
		userList = append(userList, users)
	}

	return userList, nil
}

func (dao *userDao) FindBy(col string, value string) (User, error) {
	sqlStmt := `SELECT * FROM users WHERE` + col + ` = '` + value + `';`

	rows, err := dao.Query(sqlStmt)
	defer rows.Close()
	CheckErr(err)
	users := User{}
	if rows.Next() {
		err = rows.scan(
			&users.ID, &users.Key, &users.Username, &users.Password, &users.email, &users.Phone)
	}
	CheckErr(err)
	return users, nil
}

func (dao *userDao) DeleteByKey(key string) error {
	sqlStmt := `DELETE FROM users WHERE key = '` + key + `';`
	_, err := dao.Exec(sqlStmt)
	CheckErr(err)
	return nil
}
