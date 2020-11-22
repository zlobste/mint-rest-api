package sqlstore

import (
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

type InstitutionRepository struct {
	store *Store
}

func (i *InstitutionRepository) Create(model *models.Institution) error {
	if err := model.Validate(); err != nil {
		return err
	}

	return i.store.db.QueryRow(
		"INSERT INTO institutions (title, address) VALUES ($1,$2) RETURNING id",
		model.Title,
		model.Address,
	).Scan(&model.Id)
}

func (i *InstitutionRepository) FindByTitle(title string) ([]models.Institution, error) {
	rows, err := i.store.db.Query("SELECT * FROM institutions WHERE disabled = false AND title LIKE $1", title + "%")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	institutions := []models.Institution{}
	
	for rows.Next(){
		model := models.Institution{}
		err := rows.Scan(&model.Id, &model.Title, &model.Address)
		if err != nil {
			return  nil, err
		}
		institutions = append(institutions, model)
	}
	
	return institutions, nil
}

func (i *InstitutionRepository) DeleteById(id int64) error {
	_, err := i.store.db.Exec("DELETE FROM institutions where id = $1", id)
	return err
}

