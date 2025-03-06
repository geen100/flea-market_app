package services

import (
	"gin_fleamarket/dto"
	"gin_fleamarket/models"
	"gin_fleamarket/repository"
)

type IItemServices interface {
	FindAll() (*[]models.Item, error) 
	FindById(itemId uint)(*models.Item,error)
	Create(CreateItemInput dto.CreateItemInput)(*models.Item,error)
}

type ItemService struct {
	repository repository.IItemRepository
}

func NewItemServices(repository repository.IItemRepository) IItemServices { 
	return &ItemService{repository: repository}
}

func (s *ItemService) FindAll() (*[]models.Item, error) { 
	return s.repository.FindAll()
}

func (s * ItemService) FindById(itemId uint)(*models.Item,error){
	return s.repository.FindById(itemId)
}

func (s * ItemService) Create(CreateitemInput dto.CreateItemInput) (*models.Item,error){
	newItem := models.Item{
		Name: CreateitemInput.Name,
		Price: CreateitemInput.Price,
		Description: CreateitemInput.Description,
		SoldOut: false,
	}
	return s.repository.Create(newItem)
}
