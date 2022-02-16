package repository

import (
	"context"
	"fmt"

	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

// HumanRepository is a crud interface for working with db where people info stored
type HumanRepository interface {
	Create(ctx context.Context, human model.Human) error
	Get(ctx context.Context, name string) (*model.Human, error)
	Update(ctx context.Context, name string, human model.Human) error
	Delete(ctx context.Context, name string) error
}

// HumanPostgresRepository implements crud interface with human entity
type HumanPostgresRepository struct {
	db *pgxpool.Pool
}

// NewHumanRepository returns new PostgresRepository
func NewHumanRepository(p *pgxpool.Pool) *HumanPostgresRepository {
	return &HumanPostgresRepository{db: p}
}

// Create is used for creating human info in db
func (r *HumanPostgresRepository) Create(ctx context.Context, h model.Human) error {
	query := "insert into people (id,name,male,age) values ($1,$2,$3,$4)"
	_, err := r.db.Exec(ctx, query, h.ID, h.Name, h.Male, h.Age)
	if err != nil {
		return fmt.Errorf("postgres create error %w", err)
	}
	return nil
}

// Get is used for getting human info from db
func (r *HumanPostgresRepository) Get(ctx context.Context, name string) (*model.Human, error) {
	h := model.Human{}
	row := r.db.QueryRow(ctx, `select * from people where name=$1`, name)
	err := row.Scan(&h.ID, &h.Name, &h.Male, &h.Age)
	if err != nil {
		return nil, fmt.Errorf("postgres get error %w", err)
	}
	return &h, nil
}

// Update is used for updating human info in db
func (r *HumanPostgresRepository) Update(ctx context.Context, name string, h model.Human) error {
	query := "update people set  name=$1,male=$2,age=$3 where name=$4"
	_, err := r.db.Exec(ctx, query, h.Name, h.Male, h.Age, name)
	if err != nil {
		return fmt.Errorf("postgres update error %w", err)
	}
	return nil
}

// Delete is used for deleting human info from db
func (r *HumanPostgresRepository) Delete(ctx context.Context, name string) error {
	rowsAffected, err := r.db.Exec(ctx, "delete from people where name=$1", name)
	if err != nil {
		return fmt.Errorf("postgres delete error %w", err)
	}
	if rowsAffected.RowsAffected() == 0 {
		return fmt.Errorf("no such human in db")
	}
	return nil
}
