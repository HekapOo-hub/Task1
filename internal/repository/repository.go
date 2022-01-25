package repository

import (
	"context"
	"fmt"
	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repo interface {
	Create(context.Context, model.Human) error
	Get(context.Context, int) (*model.Human, error)
	Update(context.Context, int, model.Human) error
	Delete(context.Context, int) error
}
type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(p *pgxpool.Pool) *Repository {
	return &Repository{db: p}
}

func (r *Repository) Create(ctx context.Context, h model.Human) error {
	query := "insert into people (name,male,age) values ($1,$2,$3)"
	_, err := r.db.Exec(ctx, query, h.Name, h.Male, h.Age)
	if err != nil {
		return fmt.Errorf("can't create human %w", err)
	}
	return nil
}

func (r *Repository) Get(ctx context.Context, id int) (*model.Human, error) {
	h := model.Human{}
	row := r.db.QueryRow(ctx, "select * from people where id=$1", id)
	err := row.Scan(&h.Id, &h.Name, &h.Male, &h.Age)
	if err != nil {
		return nil, fmt.Errorf("can't get human info %w", err)
	} else {
		return &h, nil
	}
}
func (r *Repository) Update(ctx context.Context, id int, h model.Human) error {
	query := "update people set  name=$1,male=$2,age=$3 where id=$4"
	_, err := r.db.Exec(ctx, query, h.Name, h.Male, h.Age, id)
	if err != nil {
		return fmt.Errorf("can't update human %w", err)
	}
	return nil
}
func (r *Repository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, "delete from people where id=$1", id)
	if err != nil {
		return fmt.Errorf("can't delete info %w", err)
	}
	return nil
}
