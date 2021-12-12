package service

import (
	"context"

	"github.com/Mague/neo4j-example/graph/model"
	"github.com/Mague/neo4j-example/models"
	"github.com/Mague/neo4j-example/repository"
)

type Service struct {
	repository repository.Repository
}

func NewService(r repository.Repository) Service {
	return Service{
		repository: r,
	}
}

func (s *Service) FindMovieByUUID(ctx context.Context, uuid string) (*models.Movie, error) {
	return s.repository.FindMovieByUUID(ctx, uuid)
}

func (s *Service) FindDirectorsByMovieUUID(ctx context.Context, uuid string) ([]*models.Person, error) {
	return s.repository.FindPersonByMovieUUID(ctx, "DIRECTED", uuid)
}

func (s *Service) FindWriterByUUID(ctx context.Context, uuid string) ([]*models.Person, error) {
	return s.repository.FindPersonByMovieUUID(ctx, "WROTER", uuid)
}

func (s *Service) FindCastByMovieUUID(ctx context.Context, uuid string) ([]*models.Person, error) {
	return s.repository.FindPersonByMovieUUID(ctx, "ACTED_IN", uuid)
}

func (s *Service) FindMovieParticipationsByPersonUUID(ctx context.Context, uuid string) ([]*model.Participation, error) {
	return s.repository.FindMovieParticipationsByPersonUUID(ctx, uuid)
}
