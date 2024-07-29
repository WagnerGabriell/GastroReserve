package usecases

import (
	dto "GastroReserve/internal/DTO"
	"GastroReserve/internal/infra/repositories"
)

type GetAllReserveUseCase struct {
	Repository repositories.IReserveRepositoryMysql
}

func NewGetAllReserveUseCase(repository repositories.IReserveRepositoryMysql) *GetAllReserveUseCase {
	return &GetAllReserveUseCase{Repository: repository}
}

func (u *GetAllReserveUseCase) Execute() ([]*dto.ReserveOutputDTO, error) {
	reserves, err := u.Repository.GetReserve()
	if err != nil {
		return nil, err
	}
	var reserveOutputDTO []*dto.ReserveOutputDTO
	for _, r := range reserves {
		reserveOutputDTO = append(reserveOutputDTO, &dto.ReserveOutputDTO{
			Id:       r.Id,
			Custumer: r.Custumer,
			TableId:  r.TableId,
			Data:     r.Data,
		})
	}
	return reserveOutputDTO, nil
}
