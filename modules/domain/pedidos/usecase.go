package pedidos

import "main/models"

type UseCase interface {
	FetchPedidos() ([]models.Pedido, error)
	GetPedido(ID uint64) (models.Pedido, error)
	GetUsuarioByPedido(pedidoId uint64) (models.Usuario, error)
	GetUsuario(usuarioID uint64) (models.Usuario, error)
	AddPedido(name string, quantidade uint, usuario models.Usuario) (models.Pedido, error)
	UpdatePedido(id uint64, name string, quantidade uint) (models.Pedido, error)
	RemovePedido(id uint64) error
}
