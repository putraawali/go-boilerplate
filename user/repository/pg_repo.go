package repository

import (
	"context"
	"database/sql"
	"service-notif/model"
)

type UserRepo struct {
	DB *sql.DB
}

func UserConnection(db *sql.DB) UserRepositoryContract {
	return &UserRepo{
		DB: db,
	}
}

func (repo *UserRepo) GetByID(ctx context.Context, user_id int) (model.UserModel, error) {
	script := "select id,  transporter_id, shipper_id from mt_master_user mmu where id = $1 limit 1"

	cursor, err := repo.DB.QueryContext(ctx, script, user_id)
	defer cursor.Close()
	if err != nil {
		return model.UserModel{}, err
	}

	var userData model.UserModel

	for cursor.Next() {
		err = cursor.Scan(
			&userData.Id,
			&userData.TransporterId,
			&userData.ShipperId,
		)
	}

	return userData, nil
}
