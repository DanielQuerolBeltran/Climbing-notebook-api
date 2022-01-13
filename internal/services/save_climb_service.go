package services

import (
	"context"

	entities "github.com/DanielQuerolBeltran/Climbing-notebook-api/internal/platform"
)

type SaveClimbService struct {
	climbRepository entities.ClimbRepository
}


func NewSaveClimbService(climbRepository entities.ClimbRepository) SaveClimbService {
	return SaveClimbService{
		climbRepository: climbRepository,
	}
}

func (s SaveClimbService) Execute(ctx context.Context, date string, grade string, description string, area string) (entities.Climb, error) {
	climb, err := entities.NewClimb(
		entities.NewClimbId().String(),
		date,
		grade,
		description,
		area,
	)
	if err != nil {
		return entities.Climb{}, err
	}

	err = s.climbRepository.Save(ctx, climb)

	if err != nil {
		return entities.Climb{}, err
	}

	return climb, nil
}