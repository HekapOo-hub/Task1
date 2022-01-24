package repository

import (
	"Task1/internal/config"
	"Task1/internal/model"
	"context"
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

func NewRepository(ctx context.Context, try int) (*Repository, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.Connect(ctx, cfg.GetURL())
	if err != nil {
		return nil, err
	}
	if err = pool.Ping(ctx); err != nil {
		return NewRepository(ctx, try+1)
	}
	r := &Repository{db: pool}
	return r, nil
}

func (r *Repository) Create(ctx context.Context, h model.Human) error {
	query := "insert into people (name,male,age) values ($1,$2,$3)"
	_, err := r.db.Exec(ctx, query, h.Name, h.Male, h.Age)
	return err
}

func (r *Repository) Get(ctx context.Context, id int) (*model.Human, error) {
	h := model.Human{}
	row := r.db.QueryRow(ctx, "select * from people where id=$1", id)
	err := row.Scan(&h.Id, &h.Name, &h.Male, &h.Age)
	if err != nil {
		return nil, err
	} else {
		return &h, nil
	}
}
func (r *Repository) Update(ctx context.Context, id int, h model.Human) error {
	query := "update people set  name=$1,male=$2,age=$3 where id=$4"
	_, err := r.db.Exec(ctx, query, h.Name, h.Male, h.Age, id)
	return err
}
func (r *Repository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, "delete from people where id=$1", id)
	return err
}
