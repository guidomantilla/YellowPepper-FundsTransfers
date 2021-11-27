package repository

import (
	"YellowPepper-FundsTransfers/pkg/app/core/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

/* TYPES DEFINITION */

type TransferRepository interface {
	Create(context context.Context, tx *sql.Tx, transfer *model.Transfer) error
	FindById(context context.Context, tx *sql.Tx, id int64) (*model.Transfer, error)
	FindAll(context context.Context, tx *sql.Tx) (*[]model.Transfer, error)
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

func (repository *DefaultTransferRepository) Create(context context.Context, tx *sql.Tx, transfer *model.Transfer) error {
	statement, err := tx.Prepare(repository.statementCreate)
	if err != nil {
		return err
	}
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			log.Println("Error closing the statement")
		}
	}(statement)

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

func (repository *DefaultTransferRepository) FindById(context context.Context, tx *sql.Tx, id int64) (*model.Transfer, error) {
	statement, err := tx.Prepare(repository.statementFindById)
	if err != nil {
		return nil, err
	}
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			log.Println("Error closing the statement")
		}
	}(statement)

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

func (repository *DefaultTransferRepository) FindAll(context context.Context, tx *sql.Tx) (*[]model.Transfer, error) {
	statement, err := tx.Prepare(repository.statementFind)
	if err != nil {
		return nil, err
	}
	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			log.Println("Error closing the statement")
		}
	}(statement)

	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Error closing the result set")
		}
	}(rows)

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
