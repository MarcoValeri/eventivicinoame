package models

import (
	"eventivicinoame/database"
	"eventivicinoame/util"
	"fmt"
)

type UserAdmin struct {
	Id       int
	Email    string
	Password string
}

func UserAdminNew(getUserAdminId int, getUserAdminEmail, getUserAdminPassword string) UserAdmin {
	setNewUserAdmin := UserAdmin{
		Id:       getUserAdminId,
		Email:    getUserAdminEmail,
		Password: getUserAdminPassword,
	}
	return setNewUserAdmin
}

func UserAdminAddNewToDB(getNewUserAdmin UserAdmin) error {
	db := database.DatabaseConnection()
	defer db.Close()

	hashThePassword, errHashPassword := util.PasswordHash(getNewUserAdmin.Password)
	if errHashPassword != nil {
		fmt.Println("Error to hash the password:", errHashPassword)
	}

	query, err := db.Query("INSERT INTO admin_users (email, password) VALUES (?, ?)", getNewUserAdmin.Email, hashThePassword)
	if err != nil {
		fmt.Println("Error adding user:", err)
		return err
	}
	defer query.Close()

	return nil
}

func UserAdminLoginIp(getUserAdminIp, getUserAdminEmail, getUserAdminPassword string) error {
	db := database.DatabaseConnection()
	defer db.Close()

	hashThePassword, errHashPassword := util.PasswordHash(getUserAdminPassword)
	if errHashPassword != nil {
		fmt.Println("Error to hash the password:", errHashPassword)
	}

	query, err := db.Query("INSERT INTO admin_user_logins (user_admin_ip, user_admin_email, user_admin_password, failed_login) VALUES (?, ?, ?, ?) ON DUPLICATE KEY UPDATE failed_login = failed_login + 1", getUserAdminIp, getUserAdminEmail, hashThePassword, 1)
	if err != nil {
		fmt.Println("Error storing attempt user login to the db:", err)
		return err
	}
	defer query.Close()

	return nil
}

func UserAdminBannedByIp(getUserAdminIp string) (bool, error) {
	db := database.DatabaseConnection()
	defer db.Close()

	query := "SELECT COUNT(*) FROM admin_user_logins WHERE user_admin_ip=? AND failed_login >= 3"

	var count int
	err := db.QueryRow(query, getUserAdminIp).Scan(&count)
	if err != nil {
		fmt.Println("Error getting user admin banned:", err)
		return true, err
	}

	return count > 0, nil
}

func UserAdminLogin(getUserAdminEmail, getUserAdminPassword string) bool {
	db := database.DatabaseConnection()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM admin_users WHERE email=?", getUserAdminEmail)
	if err != nil {
		fmt.Println("Error to user admin logic query:", err)
		return false
	}
	defer rows.Close()

	var setUserAdminEmail, setUserAdminPassword string
	for rows.Next() {
		var adminUserId int
		var adminUserEmail string
		var adminUserPassword string
		err = rows.Scan(&adminUserId, &adminUserEmail, &adminUserPassword)
		if err != nil {
			fmt.Println("Error to user admin logic fetching data:", err)
			return false
		}
		setUserAdminEmail = adminUserEmail
		setUserAdminPassword = adminUserPassword
	}

	if len(setUserAdminEmail) > 0 && len(setUserAdminPassword) > 0 {
		userAdminPasswordMatch := util.PasswordHashChecker(getUserAdminPassword, setUserAdminPassword)
		if userAdminPasswordMatch {
			return true
		}
	}

	return false
}
