package usuarios

import "main/models"

type UseCase interface {
	FetchUsuarios() ([]models.Usuario, error)
	FetchPedidos(user_id uint64) ([]models.Pedido, error)
	GetUsuario(ID uint64) (models.Usuario, error)
	AddUser(username, password string) (models.Usuario, error)
	UpdateUser(id uint64, username, password string) (models.Usuario, error)
	RemoveUser(id uint64) error
}
