package repository

import (
	"YellowPepper-FundsTransfers/core/model"
	"database/sql"
	"errors"
	"fmt"
)

/* TYPES DEFINITION */

type TransferRepository interface {
	Create(transfer *model.Transfer, tx *sql.Tx) error
	FindById(id int64, tx *sql.Tx) (*model.Transfer, error)
	FindAll(tx *sql.Tx) (*[]model.Transfer, error)
}

type DefaultTransferRepository struct {
	statementCreate   string
	statementFindById string
	statementFind     string
}

/* TYPES CONSTRUCTOR */

func NewDefaultTransferRepository() *DefaultTransferRepository {
	return &DefaultTransferRepository{
		statementCreate:   "insert into transfer (origin_account, destination_account, amount, date, status) values (?, ?, ?, ?, ?)",
		statementFindById: "select id, origin_account, destination_account, amount, date, status from transfer where id = ?",
		statementFind:     "select id, origin_account, destination_account, amount, date, status from transfer",
	}
}

func (repository *DefaultTransferRepository) Create(transfer *model.Transfer, tx *sql.Tx) error {
	statement, err := tx.Prepare(repository.statementCreate)
	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(transfer.OriginAccount, transfer.DestinationAccount, transfer.Amount, transfer.Date, transfer.Status)
	if err != nil {
		return err
	}

	transfer.Id, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (repository *DefaultTransferRepository) FindById(id int64, tx *sql.Tx) (*model.Transfer, error) {
	statement, err := tx.Prepare(repository.statementFindById)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	row := statement.QueryRow(id)

	var transfer model.Transfer
	if err := row.Scan(&transfer.Id, &transfer.OriginAccount, &transfer.DestinationAccount, &transfer.Amount, &transfer.Date, &transfer.Status); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New(fmt.Sprintf("transfer with id %d not found", id))
		}
		return nil, err
	}

	return &transfer, nil
}

func (repository *DefaultTransferRepository) FindAll(tx *sql.Tx) (*[]model.Transfer, error) {
	statement, err := tx.Prepare(repository.statementFind)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.Query()
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	transfers := make([]model.Transfer, 0)
	for rows.Next() {

		var transfer model.Transfer
		if err := rows.Scan(&transfer.Id, &transfer.OriginAccount, &transfer.DestinationAccount, &transfer.Amount, &transfer.Date, &transfer.Status); err != nil {
			return nil, err
		}

		transfers = append(transfers, transfer)
	}

	return &transfers, nil
}
