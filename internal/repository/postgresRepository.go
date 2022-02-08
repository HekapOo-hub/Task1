package repository

import (
	"context"
	"fmt"

	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
	uuid "github.com/satori/go.uuid"
)

// Repository is a crud interface for working with db where people info stored
type Repository interface {
	Create(ctx context.Context, human model.Human) error
	Get(ctx context.Context, name string) (*model.Human, error)
	Update(ctx context.Context, name string, human model.Human) error
	Delete(ctx context.Context, name string) error
}

// PostgresRepository implements crud interface with human entity
type PostgresRepository struct {
	db *pgxpool.Pool
}

// NewRepository returns new PostgresRepository
func NewRepository(p *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: p}
}

// Create is used for creating human info in db
func (r *PostgresRepository) Create(ctx context.Context, h model.Human) error {
	h.ID = uuid.NewV1().String()
	query := "insert into people (id,name,male,age) values ($1,$2,$3,$4)"
	_, err := r.db.Exec(ctx, query, h.ID, h.Name, h.Male, h.Age)
	if err != nil {
		return fmt.Errorf("postgres  creation error %w", err)
	}
	return nil
}

// Get is used for getting human info from db
func (r *PostgresRepository) Get(ctx context.Context, name string) (*model.Human, error) {
	h := model.Human{}
	row := r.db.QueryRow(ctx, `select * from people where name=$1`, name)
	err := row.Scan(&h.ID, &h.Name, &h.Male, &h.Age)
	if err != nil {
		return nil, fmt.Errorf("postgres get error %w", err)
	}
	return &h, nil
}

// Update is used for updating human info in db
func (r *PostgresRepository) Update(ctx context.Context, name string, h model.Human) error {
	query := "update people set  name=$1,male=$2,age=$3 where name=$4"
	_, err := r.db.Exec(ctx, query, h.Name, h.Male, h.Age, name)
	if err != nil {
		return fmt.Errorf("postgres update error %w", err)
	}
	return nil
}

// Delete is used for deleting human info from db
func (r *PostgresRepository) Delete(ctx context.Context, name string) error {
	rowsAffected, err := r.db.Exec(ctx, "delete from people where id=$1", name)
	if err != nil {
		return fmt.Errorf("postgres delete error %w", err)
	}
	if rowsAffected.RowsAffected() == 0 {
		return fmt.Errorf("no such human in db")
	}
	return nil
}
