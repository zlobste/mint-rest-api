package sqlstore

import (
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

type DishRepository struct {
	store *Store
}

func (d *DishRepository) Create(model *models.Dish) error {
	if err := model.Validate(); err != nil {
		return err
	}
	
	return d.store.db.QueryRow(
		"INSERT INTO dishes (name, description, cost, menu_id) VALUES ($1,$2, $3, $4) RETURNING id",
		model.Name,
		model.Description,
		model.Cost,
		model.MenuId,
	).Scan(&model.Id)
}

func (d *DishRepository) FindById(id int64) (*models.Dish, error) {
	model := &models.Dish{}
	if err := d.store.db.QueryRow("SELECT id, title, description, organization_id FROM dishes WHERE id=$1", id).
		Scan(&model.Id, &model.Name, &model.Description, &model.Cost, &model.MenuId); err != nil {
		return nil, err
	}
	return model, nil
}

func (d *DishRepository) DeleteById(id int64) error {
	_, err := d.store.db.Exec("DELETE FROM dishes where id = $1", id)
	return err
}

func (d *DishRepository) GetAllDishes() ([]models.Dish, error) {
	rows, err := d.store.db.Query("select * from dishes")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	dishes := []models.Dish{}
	
	for rows.Next(){
		model := models.Dish{}
		err := rows.Scan(&model.Id, &model.Name, &model.Description, &model.Cost, &model.MenuId)
		if err != nil{
			return  nil, err
		}
		dishes = append(dishes, model)
	}
	
	return dishes, nil
}