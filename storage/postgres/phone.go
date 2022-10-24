package postgres

import (
	"context"
	"fmt"
	"task/rest_service/storage"

	restservice "task/rest_service"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type phoneRepo struct {
	db *pgxpool.Pool
}

func NewPhoneRepo(db *pgxpool.Pool) storage.PhoneRepoI {
	return &phoneRepo{
		db: db,
	}
}

func (r *phoneRepo) Create(ctx context.Context, entity restservice.Phone) (id string, err error) {
	id = uuid.NewString()

	query := `INSERT INTO phone (id, number) VALUES ($1, $2)`
	_, err = r.db.Exec(ctx, query, id, entity.Number)
	if err != nil {
		return "", fmt.Errorf("error while inserting row: %w", err)
	}

	return id, nil
}

func (r *phoneRepo) GetByPK(ctx context.Context, id string) (phone restservice.Phone, err error) {
	query := `SELECT id, number FROM phone WHERE id = $1`
	err = r.db.QueryRow(ctx, query, id).Scan(&phone.ID, &phone.Number)
	if err != nil {
		if err == pgx.ErrNoRows {
			return phone, storage.ErrorNotFound
		}

		return phone, fmt.Errorf("error while getting row: %w", err)
	}

	return phone, nil
}

func (r *phoneRepo) GetList(ctx context.Context, limit, offset int) (res []restservice.Phone, count int32, err error) {

	cQuery := `SELECT count(1) FROM phone `

	row := r.db.QueryRow(ctx, cQuery)
	err = row.Scan(&count)
	if err != nil {
		return nil, 0, fmt.Errorf("error while getting count of rows: %w", err)
	}

	query := `SELECT id, number FROM phone ORDER BY id LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("error while getting list of rows: %w", err)
	}

	for rows.Next() {
		var phone restservice.Phone
		err = rows.Scan(&phone.ID, &phone.Number)
		if err != nil {
			return nil, 0, fmt.Errorf("error while scanning row: %w", err)
		}

		res = append(res, phone)
	}

	return res, count, nil
}

func (r *phoneRepo) Update(ctx context.Context, entity restservice.Phone) (rowsAffected int64, err error) {
	query := `UPDATE phone SET number = $1 WHERE id = $2`

	res, err := r.db.Exec(ctx, query, entity.Number, entity.ID)
	if err != nil {
		return 0, fmt.Errorf("error while updating row: %w", err)
	}

	return res.RowsAffected(), nil
}

func (r *phoneRepo) Delete(ctx context.Context, id string) (rowsAffected int64, err error) {
	query := `DELETE FROM phone WHERE id = $1`

	res, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return 0, fmt.Errorf("error while deleting row: %w", err)
	}

	return res.RowsAffected(), nil
}
