package repositories

import (
	"database/sql"
	"fmt"

	"github.com/wptest/configs"
	"github.com/wptest/internal/models"
	"github.com/wptest/pkg/mysql"
)

type DeviceRepository struct {
	db     mysql.MySqlFactory
	config *configs.Config
}

type IDeviceRepository interface {
	Store(deviceRequest models.Device) error
	GetById(id uint64) (models.Device, error)
	GetAll() ([]models.Device, error)
}

func NewDeviceRepository(mysqlDB mysql.MySqlFactory, cfg *configs.Config) *DeviceRepository {
	return &DeviceRepository{
		db:     mysqlDB,
		config: cfg,
	}
}

func (d *DeviceRepository) GetAll() (devices []models.Device, err error) {
	db, err := d.db.GetDB()
	if err != nil {
		return
	}

	rows, err := db.Query(fmt.Sprintf(`select id, device,  value, updated_at from deviceMaster`))

	if err != nil {
		return
	}

	for rows.Next() {
		var device models.Device
		err = rows.Scan(
			&device.ID,
			&device.Device,
			&device.Value,
			&device.UpdatedAt,
		)

		devices = append(devices, device)
	}

	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func (d *DeviceRepository) GetById(id uint64) (device models.Device, err error) {
	db, err := d.db.GetDB()
	if err != nil {
		return
	}

	rows, err := db.Query(fmt.Sprintf(`
		select id, device, value, updated_at from deviceMaster
		WHERE id=?
	`), id)

	if err != nil {
		return
	}

	for rows.Next() {
		err = rows.Scan(
			&device.ID,
			&device.Device,
			&device.Value,
			&device.UpdatedAt,
		)
	}

	if err != nil && err != sql.ErrNoRows {
		return
	}
	return
}

func (d *DeviceRepository) Store(device models.Device) (err error) {
	db, err := d.db.GetDB()
	if err != nil {
		return
	}

	tx, err := db.Begin()
	if err != nil {
		return
	}

	stmtUptd, err := tx.Prepare(`update deviceMaster set value =?, updated_at=? where device=?`)
	res, err := stmtUptd.Exec(device.Value, device.UpdatedAt, device.Device)

	if err != nil {
		if errTx := tx.Rollback(); errTx != nil {
			return errTx
		}
		return
	}
	rowaffect, err := res.RowsAffected()
	err = tx.Commit()
	if rowaffect == 0 {
		tx, err := db.Begin()
		if err != nil {
			return err
		}

		stmt, err := tx.Prepare("INSERT INTO deviceMaster (device, value, updated_at) VALUES (?,?,?)")
		if err != nil {
			return err
		}
		res, err = stmt.Exec(device.Device, device.Value, device.UpdatedAt)
		if err != nil {
			if errTx := tx.Rollback(); errTx != nil {
				return errTx
			}
			return err
		}

		_, err = res.LastInsertId()
		if err != nil {
			if errTx := tx.Rollback(); errTx != nil {
				return errTx
			}
			return err
		}
		tx.Commit()
	}
	return
}
