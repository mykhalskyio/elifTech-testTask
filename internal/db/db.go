package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/mykhalskyio/elifTech-testTask/internal/config"
	_ "github.com/mykhalskyio/elifTech-testTask/internal/migrations"
	"github.com/mykhalskyio/elifTech-testTask/internal/model"
	"github.com/pressly/goose/v3"
)

type Postgres struct {
	*sql.DB
}

func NewConnect(cfg *config.Config) (*Postgres, error) {
	connect, err := sql.Open("postgres", fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=%s",
		cfg.Postgres.Name,
		cfg.Postgres.User,
		cfg.Postgres.Pass,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Sslmode))
	if err != nil {
		return nil, err
	}
	if err = connect.Ping(); err != nil {
		return nil, err
	}
	return &Postgres{connect}, nil
}

func (pg *Postgres) MigrationInitUp() error {
	err := goose.Up(pg.DB, ".")
	if err != nil {
		return err
	}
	return nil
}

func (pg *Postgres) MigrationInitDown() error {
	err := goose.Down(pg.DB, ".")
	if err != nil {
		return err
	}
	return nil
}

// Bank
func (pg *Postgres) GetBanks() (*[]model.Bank, error) {
	rows, err := pg.Query("SELECT * FROM banks;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		id             int
		bankName       string
		interestRate   int
		maxLoan        int
		minDownPayment int
		loanTerm       float64
	)
	banks := []model.Bank{}
	for rows.Next() {
		err := rows.Scan(&id, &bankName, &interestRate, &maxLoan, &minDownPayment, &loanTerm)
		if err != nil {
			return nil, err
		}
		banks = append(banks, model.Bank{Id: id, BankName: bankName, InterestRate: interestRate, MaxLoan: maxLoan, MinDownPayment: minDownPayment, LoanTerm: loanTerm})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &banks, nil
}

func (pg *Postgres) GetBank(id int) (*model.Bank, error) {
	var (
		idb            int
		bankName       string
		interestRate   int
		maxLoan        int
		minDownPayment int
		loanTerm       float64
	)
	row := pg.QueryRow("SELECT * FROM banks WHERE id=$1", id)
	err := row.Scan(&idb, &bankName, &interestRate, &maxLoan, &minDownPayment, &loanTerm)
	if err != nil {
		return nil, err
	}
	return &model.Bank{Id: idb, BankName: bankName, InterestRate: interestRate, MaxLoan: maxLoan, MinDownPayment: minDownPayment, LoanTerm: loanTerm}, nil
}

func (pg *Postgres) CreateBank(bank *model.Bank) error {
	_, err := pg.Exec("INSERT INTO banks(bank_name, interest_rate, max_loan, min_down_payment, loan_term) VALUES($1, $2, $3, $4, $5);", bank.BankName, bank.InterestRate, bank.MaxLoan, bank.MinDownPayment, bank.LoanTerm)
	if err != nil {
		return err
	}
	return nil
}

func (pg *Postgres) EditBank(bank *model.Bank) error {
	_, err := pg.Exec("UPDATE banks SET bank_name = $1, interest_rate = $2, max_loan = $3, min_down_payment = $4, loan_term = $5 WHERE id = $6", bank.BankName, bank.InterestRate, bank.MaxLoan, bank.MinDownPayment, bank.LoanTerm, bank.Id)
	if err != nil {
		return err
	}
	return nil
}

func (pg *Postgres) DeleteBank(id int) error {
	_, err := pg.Exec("DELETE FROM banks WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

// Mortgage
func (pg *Postgres) GetMortgages() (*[]model.Mortgage, error) {
	rows, err := pg.Query("SELECT * FROM mortgages;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		id             int
		initialLoan    int
		downPayment    int
		monthlyPayment float64
		bankId         int
	)
	mortgages := []model.Mortgage{}
	for rows.Next() {
		err := rows.Scan(&id, &initialLoan, &downPayment, &monthlyPayment, &bankId)
		if err != nil {
			return nil, err
		}
		mortgages = append(mortgages, model.Mortgage{Id: id, InitialLoan: initialLoan, DownPayment: downPayment, MonthlyPayment: monthlyPayment, BankId: bankId})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &mortgages, nil
}

func (pg *Postgres) CreateMortage(mortgage *model.Mortgage) error {
	_, err := pg.Exec("INSERT INTO mortgages(initial_loan, down_payment, monthly_payment, bank_id) VALUES($1, $2, $3, $4)", mortgage.InitialLoan, mortgage.DownPayment, mortgage.MonthlyPayment, mortgage.BankId)
	if err != nil {
		return err
	}
	return nil
}

func (pg *Postgres) EditMortage(mortgage *model.Mortgage) error {
	_, err := pg.Exec("UPDATE mortgages SET initial_loan = $1, down_payment = $2, monthly_payment = $3, bank_id = $4 WHERE id = $5", mortgage.InitialLoan, mortgage.DownPayment, mortgage.MonthlyPayment, mortgage.BankId, mortgage.Id)
	if err != nil {
		return err
	}
	return nil
}

func (pg *Postgres) DeleteMortage(id int) error {
	_, err := pg.Exec("DELETE FROM mortgages WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
