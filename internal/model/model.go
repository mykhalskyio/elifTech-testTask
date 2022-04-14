package model

type Bank struct {
	Id             int     `json:"id"`
	BankName       string  `json:"bankname"`
	InterestRate   int     `json:"interestrate"`
	MaxLoan        int     `json:"maxloan"`
	MinDownPayment int     `json:"mindownpayment"`
	LoanTerm       float64 `json:"loanterm"`
}

type Mortgage struct {
	Id             int     `json:"id"`
	InitialLoan    int     `json:"initialloan"`
	DownPayment    int     `json:"downpayment"`
	MonthlyPayment float64 `json:"monthlypayment"`
	BankId         int     `json:"bankid"`
}
