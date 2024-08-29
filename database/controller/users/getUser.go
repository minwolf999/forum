package users

import (
	"forum/database/initialisation"
	"forum/structure"
)

// Retrieve an User in the database by an UUID and a password
func GetUser(UUID, password string) (structure.Users, error) {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return structure.Users{}, err
	}
	defer db.Close()

	datas, err := db.Query("SELECT * FROM `Users` WHERE `UUID`=? AND `Password`=?", UUID, password)
	if err != nil {
		return structure.Users{}, err
	}
	defer datas.Close()

	var user structure.Users
	for datas.Next() {
		err = datas.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.RegistrationDate, &user.Role, &user.Connected, &user.ImagePath, &user.PromoteRequest, &user.Uuid)
		if err != nil {
			return structure.Users{}, err
		}
	}

	return user, nil
}

// Retrieve all request mod users
func GetAllReqMod() error {
	structure.Html.AllReqMod = []structure.Users{}
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	datas, err := db.Query("SELECT * FROM `Users` WHERE `PromoteRequest`=?", 1)
	if err != nil {
		return err
	}
	defer datas.Close()

	for datas.Next() {
		var user structure.Users
		err = datas.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.RegistrationDate, &user.Role, &user.Connected, &user.ImagePath, &user.PromoteRequest, &user.Uuid)
		if err != nil {
			return err
		}
		structure.Html.AllReqMod = append(structure.Html.AllReqMod, user)
	}

	return nil
}

// Retrieve all mods users
func GetAllMod() error {
	structure.Html.AllMod = []structure.Users{}
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	datas, err := db.Query("SELECT * FROM `Users` WHERE `Role`=?", 1)
	if err != nil {
		return err
	}
	defer datas.Close()

	for datas.Next() {
		var user structure.Users
		err = datas.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.RegistrationDate, &user.Role, &user.Connected, &user.ImagePath, &user.PromoteRequest, &user.Uuid)
		if err != nil {
			return err
		}
		structure.Html.AllMod = append(structure.Html.AllMod, user)
	}

	return nil
}

// Retrieve an User in the database by his ID
func GetUserById(id int) (string, string, error) {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return "", "", err
	}
	defer db.Close()

	datas, err := db.Query("SELECT `Username`, `ImagePath` FROM `Users` WHERE `Id`=?", id)
	if err != nil {
		return "", "", err
	}
	defer datas.Close()

	for datas.Next() {
		var name, imagePath string
		err = datas.Scan(&name, &imagePath)
		return name, imagePath, err
	}
	return "", "", nil
}

// Retrieve an User in the database by his UUID
func GetUserByUUID(UUID string) (structure.Users, error) {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return structure.Users{}, err
	}
	defer db.Close()

	datas, err := db.Query("SELECT * FROM `Users` WHERE `UUID`=?", UUID)
	if err != nil {
		return structure.Users{}, err
	}
	defer datas.Close()

	var user structure.Users

	for datas.Next() {
		err = datas.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.RegistrationDate, &user.Role, &user.Connected, &user.ImagePath, &user.PromoteRequest, &user.Uuid)
		if err != nil {
			return structure.Users{}, err
		}
	}

	return user, nil
}

func GetUserByEmail(email string) (structure.Users, error) {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return structure.Users{}, err
	}
	defer db.Close()

	datas, err := db.Query("SELECT * FROM `Users` WHERE `Email`=?", email)
	if err != nil {
		return structure.Users{}, err
	}
	defer datas.Close()

	var user structure.Users

	for datas.Next() {
		err = datas.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.RegistrationDate, &user.Role, &user.Connected, &user.ImagePath, &user.PromoteRequest, &user.Uuid)
		if err != nil {
			return structure.Users{}, err
		}
	}

	return user, nil
}