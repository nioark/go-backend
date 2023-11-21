package usecase

import (
	"errors"
	"main/models"
	"main/modules/domain/pedidos"
	"main/modules/domain/usuarios"
)

type usecase struct {
	usuarios usuarios.UseCase
	repo     pedidos.Repository
}

func New(usuarios usuarios.UseCase, repo pedidos.Repository) pedidos.UseCase {
	return &usecase{usuarios: usuarios, repo: repo}
}

func (u usecase) FetchPedidos() ([]models.Pedido, error) {
	return u.repo.FetchPedidos()
}

func (u usecase) GetPedido(ID uint64) (models.Pedido, error) {
	return u.repo.GetPedido(ID)
}

func (u usecase) GetUsuarioByPedido(pedidoId uint64) (models.Usuario, error) {
	return u.repo.GetUsuarioByPedido(pedidoId)
}

func (u usecase) AddPedido(usuarioID uint64, name string, quantidade uint) (models.Pedido, error) {
	if name == "" || quantidade == 0 || usuarioID == 0 {
		return models.Pedido{}, errors.New("valores de entradas invalido")
	}

	usuario, err := u.usuarios.GetUsuario(usuarioID)

	if err != nil {
		return models.Pedido{}, err
	}

	pedido := models.Pedido{
		Name:       name,
		Quantidade: quantidade,
		UsuarioID:  usuario.Id,
	}

	id, err := u.repo.Add(pedido)
	if err != nil {
		return models.Pedido{}, err
	}

	return u.GetPedido(id)
}
func (u usecase) UpdatePedido(id uint64, name string, quantidade uint) (models.Pedido, error) {
	if id == 0 || name == "" || quantidade == 0 {
		return models.Pedido{}, errors.New("valores de entradas invalido")
	}

	if err := u.repo.Update(id, name, quantidade); err != nil {
		return models.Pedido{}, err
	}

	return u.GetPedido(id)
}

func (u usecase) RemovePedido(id uint64) error {
	if id == 0 {
		return errors.New("id invalido")
	}

	return u.repo.Remove(id)
}
