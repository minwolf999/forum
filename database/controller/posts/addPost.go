package posts

import (
	"forum/database/initialisation"
	"forum/structure"
	"time"
)

// Add a Reportpost in the database
func CreateNewPost(idCategories string, idCreator int, name, description, imagePath string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO `Posts`(`IdCategories`, `IdCreator`, `Name`, `Description`, `CreationDate`, `ImagePath`) VALUES(?,?,?,?,?,?)", idCategories, idCreator, name, description, time.Now().Format("2006-01-02 15:04:05"), imagePath)
	return err
}

// Add the id of a Post in ReportPost table
func AddToReportPosts(id int) error {
	structure.Html.ReportPosts = []structure.Post{}
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO `ReportPosts` (`IdPost`) VALUES(?)", id)
	return err
}
