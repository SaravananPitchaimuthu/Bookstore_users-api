package users

import (
	"fmt"

	"github.com/SaravananPitchaimuthu/Bookstore_users-api/datasources/mysql/users_db"
	"github.com/SaravananPitchaimuthu/Bookstore_users-api/utils/errors"
	"github.com/SaravananPitchaimuthu/Bookstore_users-api/utils/mysql_utils"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO users(first_name,last_name,email,date_created,password,status) VALUES (?,?,?,?,?,?);"
	queryGetUser     = "SELECT id,first_name,last_name,email,date_created,password,status FROM users WHERE id=?;"
	errorNoRows      = "no rows in result set"
	queryUpdateUser  = "UPDATE users SET first_name=?,last_name=?,email=?,status=? WHERE id=?;"
	queryDeleteUser  = "DELETE FROM users WHERE id=?;"
	queryFindUser    = "SELECT id,first_name,last_name,email,date_created,password,status FROM users WHERE status=?;"
)

// var (
// 	usersDB = make(map[int64]*User)
// )

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Password, &user.Status); err != nil {
		return mysql_utils.ParseError(err)
		// if strings.Contains(err.Error(), errorNoRows) {
		// 	return errors.NewNotFoundError(fmt.Sprintf("user id %d not found", user.Id))
		// }
		// return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d:%s", user.Id, err.Error()))
	}
	// if err := users_db.Client.Ping(); err != nil {
	// 	panic(err)
	// }
	// result := usersDB[user.Id]

	// if result == nil {
	// 	return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	// }

	// user.Id = result.Id
	// user.FirstName = result.FirstName
	// user.LastName = result.LastName
	// user.Email = result.Email
	// user.DateCreated = date_utils.GetNowString()

	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)

	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	// if saveErr != nil {
	// 	sqlErr, ok := saveErr.(*mysql.MySQLError)
	// 	if !ok {
	// 		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user : %s", saveErr.Error()))
	// 	}
	// 	// if strings.Contains(err.Error(), indexUniqueEmail) {
	// 	// 	return errors.NewBadRequestError(fmt.Sprintf("email id %s already exists", user.Email))
	// 	// }
	// 	fmt.Println(sqlErr.Number)
	// 	fmt.Println(sqlErr.Message)

	// 	switch sqlErr.Number {
	// 	case 1062:
	// 		errors.NewInternalServerError(fmt.Sprintf("email id %s already exists", user.Email))
	// 	}
	// 	return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user :%s", saveErr.Error()))
	// }
	//without validating execute the query
	//result, err := users_db.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)

	userId, err := insertResult.LastInsertId()

	if err != nil {
		return mysql_utils.ParseError(err)
	}
	// if err != nil {
	// 	return errors.NewInternalServerError(fmt.Sprintf("error when trying to save: %s", err.Error()))
	// }

	user.Id = userId
	// current := usersDB[user.Id]

	// if current != nil {
	// 	if user.Email == current.Email {
	// 		return errors.NewBadRequestError(fmt.Sprintf("email id %s already exists", user.Email))
	// 	}
	// 	return errors.NewBadRequestError(fmt.Sprintf("user id %d already exists", user.Id))
	// }
	//
	// usersDB[user.Id] = user
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Id)

	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Id); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUser)

	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)

	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Password, &user.Status); err != nil {
			fmt.Println(err)
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil

}
