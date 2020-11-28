package sqlstore

import (
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

type InstitutionRepository struct {
	store *Store
}

func (institutionRepository *InstitutionRepository) Create(model *models.Institution) error {
	if err := model.Validate(); err != nil {
		return err
	}
	
	return institutionRepository.store.db.QueryRow(
		"INSERT INTO institutions (title, address) VALUES ($1,$2) RETURNING id",
		model.Title,
		model.Address,
	).Scan(&model.Id)
}

func (institutionRepository *InstitutionRepository) FindByTitle(title string) ([]models.Institution, error) {
	rows, err := institutionRepository.store.db.Query(
		"SELECT * FROM institutions WHERE disabled = FALSE AND title LIKE $1",
		title+"%",
	)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	institutions := []models.Institution{}
	
	for rows.Next() {
		model := models.Institution{}
		err := rows.Scan(&model.Id, &model.Title, &model.Address)
		if err != nil {
			return nil, err
		}
		institutions = append(institutions, model)
	}
	
	return institutions, nil
}

func (institutionRepository *InstitutionRepository) DeleteById(id int64) error {
	_, err := institutionRepository.store.db.Exec("DELETE FROM institutions WHERE id = $1", id)
	return err
}
