package usuarios

import "main/models"

type Repository interface {
	Fetch() ([]models.Usuario, error)
	Get(ID int64) (models.Usuario, error)
	Add(models.Usuario) (models.Usuario, error)
	Update(ID int64, name, password string) error
	Remove(ID int64) error
}
