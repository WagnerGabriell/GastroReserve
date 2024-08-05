package usecases

import (
	dto "GastroReserve/internal/DTO"
	"GastroReserve/internal/entities"
	"GastroReserve/internal/infra/repositories"
)

type CreateReserveUseCase struct {
	Repository repositories.IReserveRepositoryMysql
}

func NewCreateReserveUseCase(repository repositories.IReserveRepositoryMysql) *CreateReserveUseCase {
	return &CreateReserveUseCase{Repository: repository}
}

func (u *CreateReserveUseCase) Execute(reserveInputDTO *dto.ReserveInputDTO, UId string) (dto.ReserveOutputDTO, error) {
	newReserve := entities.NewReserve(reserveInputDTO.TableId, UId, reserveInputDTO.Data)
	err := u.Repository.CreateReserve(newReserve)
	if err != nil {
		return dto.ReserveOutputDTO{}, err
	}
	return dto.ReserveOutputDTO{
		Id:      newReserve.Id,
		TableId: newReserve.TableId,
		UserId:  newReserve.UserId,
		Data:    newReserve.Data,
	}, nil
}
