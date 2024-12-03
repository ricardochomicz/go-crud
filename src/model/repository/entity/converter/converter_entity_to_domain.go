package converter

import (
	"github.com/ricardochomicz/go-crud/src/model"
	"github.com/ricardochomicz/go-crud/src/model/repository/entity"
)

func ConverterEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	domain.SetID(entity.ID.Hex())

	return domain
}

func ConverterEntitiesToDomain(
	entities []entity.UserEntity,
) []model.UserDomainInterface {
	var domains []model.UserDomainInterface
	for _, entity := range entities {
		domains = append(domains, ConverterEntityToDomain(entity))
	}
	return domains
}
