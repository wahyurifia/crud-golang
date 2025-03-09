package productmodel

import (
	"project-crud/config"
	"project-crud/entities"
)

func GetAll() []entities.Product {
	rows, err := config.DB.Query(`
	SELECT
		products.id,
		products.name,
		categories.name as category_name,
		products.stock,
		products.description,
		products.created_at,
		products.updated_at
	FROM products
	JOIN categories ON products.category_id = categories.id
	`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var products []entities.Product

	for rows.Next() {
		var product entities.Product
		err := rows.Scan(
			&product.Id, 
			&product.Name, 
			&product.Category.Name,
			&product.Stock, 
			&product.Description, 
			&product.CreatedAt, 
			&product.UpdatedAt)
		if err != nil {
			panic(err)
		}
		products = append(products, product)
	}
	return products
}

func GetById(id int) entities.Product {
	row := config.DB.QueryRow(`
	SELECT id, name, category_id, stock, description 
	FROM products
	WHERE id = $1`,id)

	var product entities.Product
	err := row.Scan(&product.Id, &product.Name, &product.Category.Id, &product.Stock, &product.Description)
	if err !=  nil {panic(err.Error())}
	return product
}

func Create(product entities.Product)bool {
	var lastInsertID int
	err := config.DB.QueryRow(`
	INSERT INTO products (name, category_id, stock, description, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
	product.Name, product.Category.Id, product.Stock, product.Description, product.CreatedAt, product.UpdatedAt).Scan(&lastInsertID)

	if err != nil {panic(err)}

	return lastInsertID > 0
}

func Update(id int, product entities.Product)bool {
	result , err := config.DB.Exec(`
	UPDATE products SET 
		name = $1, 
		stock = $2, 
		category_id = $3,
		description = $4,
		updated_at = $5
	WHERE id = $6`,
	&product.Name, &product.Stock, &product.Category.Id, &product.Description, &product.UpdatedAt, id)
	status, err := result.RowsAffected()
	if err != nil {panic(err)}

	return status > 0
}

func Delete(id int)error {
	_, err := config.DB.Exec(`
	DELETE FROM products WHERE id = $1`, id)
	if err != nil {
		panic(err)
	}
	return err
}