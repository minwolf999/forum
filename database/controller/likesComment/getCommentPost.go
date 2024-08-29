package likesComment

import "forum/database/initialisation"

// Get all the likes of a comment using its identifier
func GetLikesComment(id int) ([]int, error) {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	datas, err := db.Query("SELECT `IdUser` FROM `LikesComments` WHERE `IdComment`=?", id)
	if err != nil {
		return nil, err
	}
	defer datas.Close()

	var tab []int
	for datas.Next() {
		var i int
		err = datas.Scan(&i)
		if err != nil {
			return nil, err
		}

		tab = append(tab, i)
	}

	return tab, nil
}

// Get all the dislikes of a comment using its identifier
func GetDislikesComment(id int) ([]int, error) {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	datas, err := db.Query("SELECT `IdUser` FROM `DislikesComments` WHERE `IdComment`=?", id)
	if err != nil {
		return nil, err
	}
	defer datas.Close()

	var tab []int
	for datas.Next() {
		var i int
		err = datas.Scan(&i)
		if err != nil {
			return nil, err
		}

		tab = append(tab, i)
	}

	return tab, nil
}