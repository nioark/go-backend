package usecase

import (
	"errors"
	"main/models"
	"main/modules/domain/usuarios"
)

type usecase struct {
	repo usuarios.Repository
}

func New(repo usuarios.Repository) usuarios.UseCase {
	return &usecase{repo: repo}
}

func (u usecase) Fetch() ([]models.Usuario, error) {
	return u.repo.Fetch()
}

func (u usecase) Get(ID int64) (models.Usuario, error) {
	//todo seid for vazio ja volta erro

	return u.repo.Get(ID)
}

func (u usecase) AddUser(username, password string) (models.Usuario, error) {
	if username == "" || password == "" {
		return models.Usuario{}, errors.New("usuario e senha invalido")
	}

	usuario := models.Usuario{Name: username, Password: password}

	var err error
	usuario, err = u.repo.Add(usuario)
	if err != nil {
		return models.Usuario{}, err
	}

	return u.Get(usuario.Id)
}

func (u usecase) UpdateUser(id int64, username, password string) (models.Usuario, error) {
	if username == "" || password == "" || id == 0 {
		return models.Usuario{}, errors.New("usuario, senha ou id invalido")
	}

	//usuario := models.Usuario{Id: id, Name: username, Password: password}

	if err := u.repo.Update(id, username, password); err != nil {
		return models.Usuario{}, err
	}

	return u.Get(id)
}

func (u usecase) RemoveUser(id int64) error {
	if id == 0 {
		return errors.New("id invalido")
	}

	return u.repo.Remove(id)
}
