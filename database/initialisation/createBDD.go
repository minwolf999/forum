package initialisation

// Create the database tabs if they do not exist and fill the Category table
func CreateBDD() error {
	db, err := OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `Users` (`Id` INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, `Username` TEXT UNIQUE NOT NULL, `Email` TEXT UNIQUE NOT NULL, `Password` TEXT NOT NULL, `RegistrationDate` TEXT NOT NULL, `Role` INTEGER NOT NULL DEFAULT 0, `Connected` INTEGER NOT NULL DEFAULT 0, `ImagePath` TEXT NOT NULL DEFAULT '/front/static/imgs/Default_pfp.svg', `PromoteRequest` INTEGER NOT NULL DEFAULT 0 , `UUID` TEXT UNIQUE NOT NULL);")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `Categories` (`Id` INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, `Name` TEXT UNIQUE NOT NULL);")
	if err != nil {
		return err
	}

	var categorie = []string{"Cuisine et Gastronomie", "Bricolage et Travaux Manuels", "Sports et Activités Physique", "Arts et Culture", "Technologies et Informatique", "Santé et Bien-Être", "Voyage et Découvertes", "Musique et Divertissement", "Mode et Beauté", "Automobile et Mécanique", "Science et Innovation", "Environement et Écologie", "Éducation et Apprentissage", "Finance et Économie", "Animaux et Nature"}
	for i, v := range categorie {
		_, err = db.Exec("INSERT INTO `Categories` VALUES(?,?);", i, v)
		if err != nil {
			return err
		}
	}

	_, err = db.Exec("CREATE TABLE `Posts` (`Id` INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, `IdCategories` TEXT NOT NULL, IdCreator INTEGER NOT NULL REFERENCES`Users`(`Id`), `Name` TEXT UNIQUE NOT NULL, `Description` TEXT NOT NULL, `CreationDate` TEXT NOT NULL, `ImagePath` TEXT);")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE `ReportPosts` (`Id` INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, `IdPost` INTEGER UNIQUE NOT NULL REFERENCES `Posts`(`Id`));")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `LikesPosts` (`Id` INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, `IdPost` INTEGER NOT NULL REFERENCES `Posts`(`Id`), `IdUser` INTEGER NOT NULL REFERENCES `Users`(`Id`));")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `DislikesPosts` (`Id` INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, `IdPost` INTEGER NOT NULL REFERENCES `Posts`(`Id`), `IdUser` INTEGER NOT NULL REFERENCES `Users`(`Id`));")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `Comments` (`Id` INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, `IdPost` INTEGER NOT NULL REFERENCES `Posts`(`Id`), `IdCreator` INTEGER NOT NULL REFERENCES `Users`(`Id`), `Text` TEXT NOT NULL, `CreationDate` TEXT NOT NULL, `ImagePath` TEXT);")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `LikesComments` (`Id` INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, `IdComment` INTEGER NOT NULL REFERENCES `Comments`(`Id`), `IdUser` INTEGER NOT NULL REFERENCES `Users`(`Id`));")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `DislikesComments` (`Id` INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL, `IdComment` INTEGER NOT NULL REFERENCES `Comments`(`Id`), `IdUser` INTEGER NOT NULL REFERENCES `Users`(`Id`));")
	return err
}
