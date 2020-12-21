package sqlstore

import (
	"github.com/zlobste/mint-rest-api/internal/app/models"
)

type InstitutionRepository struct {
	store *Store
}

func (InstitutionRepository *InstitutionRepository) Create(model *models.Institution) error {
	if err := model.Validate(); err != nil {
		return err
	}

	return InstitutionRepository.store.db.QueryRow(
		"INSERT INTO institutions (title, address) VALUES ($1,$2) RETURNING id",
		model.Title,
		model.Address,
	).Scan(&model.Id)
}

func (InstitutionRepository *InstitutionRepository) FindByTitle(title string) ([]models.Institution, error) {
	rows, err := InstitutionRepository.store.db.Query(
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

func (InstitutionRepository *InstitutionRepository) GetAllInstitutions() ([]models.Institution, error) {
	rows, err := InstitutionRepository.store.db.Query("SELECT id, address, title FROM institutions")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	institutions := []models.Institution{}

	for rows.Next() {
		model := models.Institution{}
		err := rows.Scan(&model.Id, &model.Address, &model.Title)
		if err != nil {
			return nil, err
		}
		institutions = append(institutions, model)
	}

	return institutions, nil
}

func (InstitutionRepository *InstitutionRepository) DeleteById(id int64) error {
	_, err := InstitutionRepository.store.db.Exec("DELETE FROM institutions WHERE id = $1", id)
	return err
}
