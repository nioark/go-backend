package usecase

import (
	"errors"
	"main/models"
	"main/modules/domain/pedidos"
)

type usecase struct {
	repo pedidos.Repository
}

func New(repo pedidos.Repository) pedidos.UseCase {
	return &usecase{repo: repo}
}

func (u usecase) FetchPedidos() ([]models.Pedido, error) {
	return u.repo.FetchPedidos()
}

func (u usecase) GetPedido(ID uint64) (models.Pedido, error) {
	return u.repo.GetPedido(ID)
}

func (u usecase) GetUsuario(usuarioID uint64) (models.Usuario, error) {
	return u.repo.GetUsuario(usuarioID)
}

func (u usecase) GetUsuarioByPedido(pedidoId uint64) (models.Usuario, error) {
	return u.repo.GetUsuarioByPedido(pedidoId)
}

func (u usecase) AddPedido(name string, quantidade uint, usuario models.Usuario) (models.Pedido, error) {
	if name == "" || quantidade == 0 || usuario.Id == 0 || usuario.Name == "" || usuario.Password == "" {
		return models.Pedido{}, errors.New("valores de entradas invalido")
	}

	pedido := models.Pedido{
		Name:       name,
		Quantidade: quantidade,
		UsuarioID:  usuario.Id,
	}

	var err error
	pedido, err = u.repo.Add(pedido)
	if err != nil {
		return models.Pedido{}, err
	}

	return u.GetPedido(uint64(pedido.ID))
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
