package service

import (
	"errors"
	"main/internal/model"
	"main/internal/repository"
)

type AuthorService interface {
	GetAll() ([]model.Authors, error)
	GetByID(id string) ([]model.Authors, error)
	Create(req *model.CreateAuthor) error
	Replace(id string, req *model.UpdateAuthor) error
	Update(id string, req *model.UpdateAuthor) error
	Delete(id string) error
}

type authorService struct {
	repo repository.AuthorRepository
}

func NewAuthorService(repo repository.AuthorRepository) AuthorService {
	return &authorService{repo: repo}
}

func (s *authorService) GetAll() ([]model.Authors, error) {
	return s.repo.GetAll()
}

func (s *authorService) GetByID(id string) ([]model.Authors, error) {
	authors, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	if len(authors) == 0 {
		return nil, errors.New("author not found")
	}
	return authors, nil
}

func (s *authorService) Create(req *model.CreateAuthor) error {
	// Бизнес-валидация
	if req.FirstName == nil || *req.FirstName == "" {
		return errors.New("first name is required")
	}
	if req.LastName == nil || *req.LastName == "" {
		return errors.New("last name is required")
	}
	return s.repo.Create(req)
}

func (s *authorService) Replace(id string, req *model.UpdateAuthor) error {
	// Проверка на обязательные поля для PUT
	if req.FirstName == nil || *req.FirstName == "" {
		return errors.New("first name is required for replace")
	}
	if req.LastName == nil || *req.LastName == "" {
		return errors.New("last name is required for replace")
	}

	rowsAffected, err := s.repo.Replace(id, req)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("author not found")
	}
	return nil
}

func (s *authorService) Update(id string, req *model.UpdateAuthor) error {
	rowsAffected, err := s.repo.Update(id, req)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("author not found")
	}
	return nil
}

func (s *authorService) Delete(id string) error {
	rowsAffected, err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("author not found")
	}
	return nil
}
