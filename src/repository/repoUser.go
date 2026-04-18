package repository

import (
	"database/sql"
	"BACK-END-ONLINESHOP/src/model"
)

func InsertUser(db *sql.DB, User model.User) error {
	query := "INSERT INTO Users (user_id, user_name, role_id) VALUES (?, ?, ?)"

	_, err := db.Exec(query, User.User_id, User.User_name, User.Role_id)

	if err != nil {
		return err 
	}
	return nil
}

func GetAllUsers(db *sql.DB) ([]model.User, error) {
	query :="SELECT user_id, user_name, role_id FROM Users"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	var alluser []model.User

	for rows.Next(){
		user := model.User{}
		errScan := rows.Scan(&user.User_id, &user.User_name, &user.Role_id)

		if errScan != nil {
			return nil, errScan
		}

		alluser = append(alluser, user)
	}

	
	return alluser, nil
}

func UpdateUser(db *sql.DB, User model.User)error{
	query := "UPDATE Users SET user_name = ?, role_id = ? WHERE user_id = ?"

	_,err := db.Exec(query, User.User_name, User.Role_id, User.User_id)

	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(db *sql.DB, userId int) error{
	query := "DELETE FROM Users WHERE user_id = ?"

	_,err := db.Exec(query, userId)
    // err itu variable yang nyimpen error 
	if err != nil {
		return err
	}
	return nil
}