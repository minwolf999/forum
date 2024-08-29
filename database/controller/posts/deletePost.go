package posts

import (
	"forum/database/initialisation"
)

// Delete a post in the database
func DeletePost(id int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM `Posts` WHERE `Id` =?", id)
	return err
}

// Delete a post in the ReportPost table
func DeleteReportPost(id int) error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM `ReportPosts` WHERE `IdPost` =?", id)
	return err
}
