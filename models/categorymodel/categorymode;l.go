package categorymodel

import (
	"project-crud/config"
	"project-crud/entities"
)

func GetAll() []entities.Category {
 // sama seperti di expressjs -> primsa.findAll
	rows, err := config.DB.Query(`SELECT * FROM categories`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// buat array kosong lalu dengan type dari struct
	var categories []entities.Category
	
	// looping hasil dari query satu persatu
	for rows.Next() {
		var category entities.Category
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdateAt)
		if err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}
	//lalu return biar diterima oleh controller
	return categories
}