package comment

import (
	"forum/database/initialisation"
)

// Add a comment in the database
func DeleteComment(idComment int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM `Comments` WHERE `Id` =?", idComment)
	return err
}

// Delete aLl comments of a Post
func DeleteCommentsOfPost(idCreator int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM `Comments` WHERE `IdPost` =?", idCreator)
	return err
}
