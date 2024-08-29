package comment

import (
	"forum/database/initialisation"
	"time"
)

// Add a comment in the database
func CreatNewComment(idPost, idCreator int, text, imagePath string) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO `Comments`(`IdPost`, `IdCreator`, `Text`, `CreationDate`, `ImagePath`) VALUES(?,?,?,?,?)", idPost, idCreator, text, time.Now().Format("2006-01-02 15:04:05"), imagePath)
	return err
}
