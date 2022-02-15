package dockertest

import (
	"context"
	"testing"

	"github.com/HekapOo-hub/Task1/internal/model"
	"github.com/HekapOo-hub/Task1/internal/repository"
	"github.com/jackc/pgx/v4/pgxpool"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

var postgresDB *pgxpool.Pool

func TestHumanCreate(t *testing.T) {
	ctx := context.Background()
	repo := repository.NewRepository(postgresDB)
	expected := model.Human{Name: "create", Age: 11, Male: false}
	expected.ID = uuid.NewV1().String()
	err := repo.Create(ctx, expected)
	require.NoError(t, err)
	var actual model.Human
	row := postgresDB.QueryRow(ctx, `select * from people where name=$1`, expected.Name)
	err = row.Scan(&actual.ID, &actual.Name, &actual.Male, &actual.Age)
	require.NoError(t, err)
	require.Equal(t, expected.ID, actual.ID)
	require.Equal(t, expected.Male, actual.Male)
	require.Equal(t, expected.Age, actual.Age)
	_, err = postgresDB.Exec(ctx, "truncate people")
	require.NoError(t, err)
}

func TestHumanGet(t *testing.T) {
	ctx := context.Background()
	repo := repository.NewRepository(postgresDB)
	expected := model.Human{ID: uuid.NewV4().String(), Name: "get", Age: 123, Male: true}
	_, err := postgresDB.Exec(ctx, "insert into people (id,name,male,age) values ($1,$2,$3,$4)",
		expected.ID, expected.Name, expected.Male, expected.Age)
	require.NoError(t, err)
	actual, err := repo.Get(ctx, expected.Name)
	require.NoError(t, err)
	require.Equal(t, expected.Age, actual.Age)
	require.Equal(t, expected.Male, actual.Male)
	require.Equal(t, expected.ID, actual.ID)
	_, err = postgresDB.Exec(ctx, "truncate people")
	require.NoError(t, err)
}

func TestHumanUpdate(t *testing.T) {
	ctx := context.Background()
	repo := repository.NewRepository(postgresDB)
	expected := model.Human{ID: uuid.NewV4().String(), Name: "updated", Male: true, Age: 228}
	_, err := postgresDB.Exec(ctx, "insert into people (id,name,male,age) values ($1,$2,$3,$4)",
		expected.ID, "update", false, 2)
	require.NoError(t, err)

	err = repo.Update(ctx, "update", expected)
	require.NoError(t, err)
	var actual model.Human
	row := postgresDB.QueryRow(ctx, `select * from people where name=$1`, "updated")
	err = row.Scan(&actual.ID, &actual.Name, &actual.Male, &actual.Age)
	require.NoError(t, err)
	require.Equal(t, expected.ID, actual.ID)
	require.Equal(t, expected.Male, actual.Male)
	require.Equal(t, expected.Age, actual.Age)
	_, err = postgresDB.Exec(ctx, "truncate people")
	require.NoError(t, err)
}

func TestHumanDelete(t *testing.T) {
	ctx := context.Background()
	repo := repository.NewRepository(postgresDB)
	_, err := postgresDB.Exec(ctx, "insert into people (id,name,male,age) values ($1,$2,$3,$4)",
		uuid.NewV4().String(), "delete", false, 2)
	require.NoError(t, err)
	err = repo.Delete(ctx, "delete")
	require.NoError(t, err)
	var deleted model.Human
	row := postgresDB.QueryRow(ctx, `select * from people where name=$1`, "delete")
	err = row.Scan(&deleted.ID, &deleted.Name, &deleted.Male, &deleted.Age)
	row = postgresDB.QueryRow(ctx, `select * from people where name=$1`, "not existing")
	expectedError := row.Scan(&deleted.ID, &deleted.Name, &deleted.Male, &deleted.Age)
	require.Equal(t, expectedError, err)
}
