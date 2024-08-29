package users

import (
	"errors"
	"fmt"
	"forum/database/initialisation"
	"forum/structure"
	verificationfunction "forum/verificationFunction"

	"golang.org/x/crypto/bcrypt"
)

// Update the informations of a user by his ID
func UpdateUser(username, email, password, imagePath string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	if !verificationfunction.EmailVerif(email) {
		return errors.New("incorrect email format")
	}

	var request string
	if password == "" && imagePath == "" {
		request = fmt.Sprintf("UPDATE `Users` SET `Username`='%s', `Email`='%s' WHERE `Id`=?", username, email)
	} else if verificationfunction.PasswordVerif(password) && imagePath == "" {
		cryptPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)

		request = fmt.Sprintf("UPDATE `Users` SET `Username`='%s', `Email`='%s', `Password`='%s' WHERE `Id`=?", username, email, cryptPassword)
	} else if password == "" && imagePath != "" {
		request = fmt.Sprintf("UPDATE `Users` SET `Username`='%s', `Email`='%s', `ImagePath`='%s' WHERE `Id`=?", username, email, imagePath)
	} else if verificationfunction.PasswordVerif(password) && imagePath != "" {
		cryptPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)

		request = fmt.Sprintf("UPDATE `Users` SET `Username`='%s', `Email`='%s', `Password`='%s', `ImagePath`='%s' WHERE `Id`=?", username, email, cryptPassword, imagePath)
	} else {
		return errors.New("incorrect password ! The password must contain 8 characters, 1 uppercase letter, 1 special character, 1 number")
	}

	_, err = db.Exec(request, structure.Html.User.Id)
	if err != nil {
		return err
	}

	user, err := GetUser(email, password)
	structure.Html.User = user
	return err
}

func UpdateUserRole(id, role int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE `Users` SET `Role`=? WHERE `Id` =?", role, id)
	return err
}

func UpdateUserUUID(id int, uuid string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE `Users` SET `UUID`=? WHERE `Id` =?", uuid, id)
	return err
}

// Set the Connected Row of a User in the BDD at 1
func ConnectUser(id int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE `Users` SET `Connected`=1 WHERE `Id`=?", id)
	if err != nil {
		structure.Html.User = structure.Users{}
		return err
	}

	return nil
}

// Set the Connected Row of a User in the BDD at 0
func DisconnectUser(id int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE `Users` SET `Connected`=? WHERE `Id`=?", 0, id)
	if err != nil {
		return err
	}

	structure.Html.User = structure.Users{}
	return nil
}

// Update the promote request
func UpdatePromoteRequest(id int, newStatus int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE `Users` SET `PromoteRequest`=? WHERE `Id`=?", newStatus, id)
	if err != nil {
		return err
	}

	return nil
}
