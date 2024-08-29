package likespost

import "forum/database/initialisation"

// Get all the likes of a post using its ID
func GetLikesPost(id int) ([]int, error) {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	datas, err := db.Query("SELECT `IdUser` FROM `LikesPosts` WHERE `IdPost`=?", id)
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

// // Get all the dislikes of a post using its ID
func GetDislikesPost(id int) ([]int, error) {
	db, err := initialisation.OpenBDD()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	datas, err := db.Query("SELECT `IdUser` FROM `DislikesPosts` WHERE `IdPost`=?", id)
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