package repository

import (
	"context"
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
	uuid "github.com/satori/go.uuid"
)

type Repository interface {
	Create(context.Context, model.Human) error
	Get(context.Context, string) (*model.Human, error)
	Update(context.Context, string, model.Human) error
	Delete(context.Context, string) error
}
type PostgresRepository struct {
	db *pgxpool.Pool
}

func NewRepository(p *pgxpool.Pool) Repository {
	return &PostgresRepository{db: p}
}

func (r *PostgresRepository) Create(ctx context.Context, h model.Human) error {
	h.Id = uuid.NewV1().String()
	query := "insert into people (id,name,male,age) values ($1,$2,$3,$4)"
	_, err := r.db.Exec(ctx, query, h.Id, h.Name, h.Male, h.Age)
	if err != nil {
		return fmt.Errorf("postgres  creation error %w", err)
	}
	return nil
}

func (r *PostgresRepository) Get(ctx context.Context, name string) (*model.Human, error) {
	h := model.Human{}
	row := r.db.QueryRow(ctx, `select * from people where name=$1`, name)
	err := row.Scan(&h.Id, &h.Name, &h.Male, &h.Age)
	if err != nil {
		return nil, fmt.Errorf("postgres get error %w", err)
	} else {
		return &h, nil
	}
}
func (r *PostgresRepository) Update(ctx context.Context, id string, h model.Human) error {
	query := "update people set  name=$1,male=$2,age=$3 where id=$4"
	_, err := r.db.Exec(ctx, query, h.Name, h.Male, h.Age, id)
	if err != nil {
		return fmt.Errorf("postgres update error %w", err)
	}
	return nil
}
func (r *PostgresRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, "delete from people where id=$1", id)
	if err != nil {
		return fmt.Errorf("postgres delete error %w", err)
	}
	return nil
}