package usuarios

import "main/models"

type Repository interface {
	FetchUsuarios() ([]models.Usuario, error)
	FetchPedidos(ID uint64) ([]models.Pedido, error)
	Get(ID uint64) (models.Usuario, error)
	Add(models.Usuario) (models.Usuario, error)
	Update(ID uint64, name, password string) error
	Remove(ID uint64) error
}
