package users

import (
	"blog/dao/mysql/users_db"
	"blog/utils/errors"
)

var (
	queryInsertUser = "insert into users(`first_name`, `last_name`, `email`, `password`) values (?, ?, ?, ?);"
	queryGetUserByEmail = "select id, first_name, last_name, email, password from users where email = ?;"
	queryGetUserById = "select id, first_name, last_name, email from users where id = ?;"
)

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServeError("insert db error, " + err.Error())
	}

	defer stmt.Close()

	res, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return errors.NewInternalServeError("db stmt error, " + err.Error())
	}

	uid, err:= res.LastInsertId()
	if err != nil {
		return errors.NewInternalServeError("db last id error, " + err.Error())
	}

	user.Id = uid
	return nil
}

func (user *User) GetByEmail() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUserByEmail)
	if err != nil {
		return errors.NewInternalServeError("invalid email")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Email)
	getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if getErr != nil {
		return errors.NewInternalServeError("db search error")
	}

	return nil
}

func (user *User) GetById() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUserById)
	if err != nil {
		return errors.NewInternalServeError("invalid id")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email)
	if getErr != nil {
		return errors.NewInternalServeError("db search error")
	}

	return nil
}