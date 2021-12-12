package repository

import (
	"context"

	"github.com/Mague/neo4j-example/graph/model"
	"github.com/Mague/neo4j-example/models"
)

type Repository interface {
	FindMovieByUUID(ctx context.Context, uuid string) (*models.Movie, error)
	FindMovies(ctx context.Context, title *string, actor *string) ([]*models.Movie, error)
	FindMovieParticipationsByPersonUUID(ctx context.Context, uuid string) ([]*model.Participation, error)
	FindPersonByMovieUUID(ctx context.Context, role string, uuid string) ([]*models.Person, error)
}
