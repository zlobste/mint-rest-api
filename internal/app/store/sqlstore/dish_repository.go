package sqlstore

import (
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

type DishRepository struct {
	store *Store
}

func (dishRepository *DishRepository) Create(model *models.Dish) error {
	if err := model.Validate(); err != nil {
		return err
	}
	
	return dishRepository.store.db.QueryRow(
		"INSERT INTO dishes (name, description, cost) VALUES ($1,$2, $3, $4) RETURNING id",
		model.Title,
		model.Description,
		model.Cost,
	).Scan(&model.Id)
}

func (dishRepository *DishRepository) FindById(id int64) (*models.Dish, error) {
	model := &models.Dish{}
	if err := dishRepository.store.db.QueryRow("SELECT id, title, description, disabled FROM dishes WHERE id=$1", id).
		Scan(&model.Id, &model.Title, &model.Description, &model.Cost, &model.Disabled); err != nil {
		return nil, err
	}
	return model, nil
}

func (dishRepository *DishRepository) DeleteById(id int64) error {
	_, err := dishRepository.store.db.Exec("DELETE FROM dishes where id = $1", id)
	return err
}

func (dishRepository *DishRepository) GetAllDishes() ([]models.Dish, error) {
	rows, err := dishRepository.store.db.Query("SELECT * FROM dishes WHERE disabled = false")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	dishes := []models.Dish{}
	
	for rows.Next(){
		model := models.Dish{}
		err := rows.Scan(&model.Id, &model.Title, &model.Description, &model.Cost)
		if err != nil{
			return  nil, err
		}
		dishes = append(dishes, model)
	}
	
	return dishes, nil
}
