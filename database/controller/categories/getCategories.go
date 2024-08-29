package categories

import (
	"forum/database/initialisation"
	"forum/structure"
)

// Retrieve all the Categories in the database
func GetCategories() error {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return err
	}
	defer db.Close()
	
	datas, err := db.Query("SELECT * FROM `Categories`")
	if err != nil {
		return err
	}
	defer datas.Close()

	structure.Html.Categories = []structure.Categorie{}

	for datas.Next() {
		var tmp structure.Categorie

		err = datas.Scan(&tmp.Id, &tmp.Name)
		if err != nil {
			return err
		}

		structure.Html.Categories = append(structure.Html.Categories, tmp)
	}

	return nil
}

// Retrieve the Category by its identifier in the database
func GetCategoriesById(id int) (string, error) {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return "", err
	}
	defer db.Close()

	datas, err := db.Query("SELECT `Name` FROM `Categories` WHERE `Id`=?", id)
	if err != nil {
		return "", err
	}
	defer datas.Close()

	for datas.Next() {
		var res string

		err = datas.Scan(&res)
		return res, err
	}

	return "", nil
}
