package usuarios

import "main/models"

type UseCase interface {
	Fetch() ([]models.Usuario, error)
	Get(ID int64) (models.Usuario, error)
	AddUser(username, password string) (models.Usuario, error)
	UpdateUser(id int64, username, password string) (models.Usuario, error)
	RemoveUser(id int64) error
}
