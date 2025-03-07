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
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}
	//lalu return biar diterima oleh controller
	return categories
}

func GetById(id int) (entities.Category) {
	row := config.DB.QueryRow(`SELECT id, name FROM categories WHERE id = $1`,id)

	var category entities.Category

	err := row.Scan(&category.Id, &category.Name)
	if err != nil {panic(err)}

	return category
}

func Create(category entities.Category)bool {
	var lastInsertID int
	err := config.DB.QueryRow(`
	INSERT INTO categories (name, created_at, updated_at)
	VALUES ($1, $2, $3) RETURNING id`,
	category.Name, category.CreatedAt, category.UpdatedAt).Scan(&lastInsertID)
	if err != nil {panic(err)}

return lastInsertID > 0 
}

func Update(id int, category entities.Category)bool {
	result, err := config.DB.Exec(`
	UPDATE categories SET name = $1, updated_at = $2 WHERE id = $3`,
	category.Name, category.UpdatedAt, id)
	if err != nil {panic(err)}
	
	status, err := result.RowsAffected()
	if err != nil {panic(err)}
	
	return status > 0
}

func Delete(id int)error {
	_, err := config.DB.Exec(`
	DELETE FROM categories WHERE id = $1`,id)
	if err != nil {panic(err)}

	return err
}
