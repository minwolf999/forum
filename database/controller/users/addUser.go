package users

import (
	"forum/database/initialisation"
	"time"
)

// Add a User in the database
func AddUser(username, email, password string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO `Users`(`Username`, `Email`, `Password`, `RegistrationDate`,`UUID`) VALUES(?,?,?,?,?)", username, email, password, time.Now().Format("2006-01-02 15:04:05"),"")
	return err
}
