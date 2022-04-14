package handler

import (
	"log"
	"math"
	"net/http"
	"strconv"
	"text/template"

	"github.com/mykhalskyio/elifTech-testTask/internal/db"
	"github.com/mykhalskyio/elifTech-testTask/internal/model"
)

type Handler struct {
	db *db.Postgres
}

func NewHandler(db *db.Postgres) *Handler {
	return &Handler{db}
}

type Data struct {
	Bank      *model.Bank
	Banks     *[]model.Bank
	Mortgages *[]model.Mortgage
}

// Banks
func (h *Handler) Banks(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/bank/home.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	banks, err := h.db.GetBanks()
	if err != nil {
		log.Println(err)
		return
	}

	err = ts.Execute(w, Data{Banks: banks})
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
	}
}

func (h *Handler) DeleteBank(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println(err)
		return
	}
	err = h.db.DeleteBank(id)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/banks", 301)
}

func (h *Handler) EditBankPage(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/bank/edit.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println(err)
	}
	bank, err := h.db.GetBank(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, Data{Bank: bank})
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
	}
}

func (h *Handler) EditBank(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue("Id"))
	if err != nil {
		log.Println(err)
	}
	BankName := r.FormValue("bankName")
	InterestRate, err := strconv.Atoi(r.FormValue("InterestRate"))
	if err != nil {
		log.Println(err)
	}
	MaximumLoan, err := strconv.Atoi(r.FormValue("MaximumLoan"))
	if err != nil {
		log.Println(err)
	}
	Minimumdownpayment, err := strconv.Atoi(r.FormValue("Minimumdownpayment"))
	if err != nil {
		log.Println(err)
	}
	Loanterm, err := strconv.ParseFloat(r.FormValue("Loanterm"), 64)
	if err != nil {
		log.Println(err)
	}
	err = h.db.EditBank(&model.Bank{Id: id, BankName: BankName, InterestRate: InterestRate, MaxLoan: MaximumLoan, MinDownPayment: Minimumdownpayment, LoanTerm: Loanterm})
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
	}
	http.Redirect(w, r, "/banks", 301)
}

func (h *Handler) AddBank(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/bank/add.html")
	if err != nil {
		log.Println(err)
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
	}
}

func (h *Handler) AddBankdb(w http.ResponseWriter, r *http.Request) {
	BankName := r.FormValue("bankName")
	InterestRate, err := strconv.Atoi(r.FormValue("InterestRate"))
	if err != nil {
		log.Println(err)
	}
	MaximumLoan, err := strconv.Atoi(r.FormValue("MaximumLoan"))
	if err != nil {
		log.Println(err)
	}
	Minimumdownpayment, err := strconv.Atoi(r.FormValue("Minimumdownpayment"))
	if err != nil {
		log.Println(err)
	}
	Loanterm, err := strconv.ParseFloat(r.FormValue("Loanterm"), 64)
	if err != nil {
		log.Println(err)
	}
	err = h.db.CreateBank(&model.Bank{BankName: BankName, InterestRate: InterestRate, MaxLoan: MaximumLoan, MinDownPayment: Minimumdownpayment, LoanTerm: Loanterm})
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		w.Write([]byte("Error write"))
	}
	http.Redirect(w, r, "/banks", 301)
}

// Mortgages
func (h *Handler) Mortgages(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/mortgage/home.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	mortgages, err := h.db.GetMortgages()
	if err != nil {
		log.Println(err)
		return
	}
	banks, err := h.db.GetBanks()
	if err != nil {
		log.Println(err)
		return
	}
	err = ts.Execute(w, Data{Mortgages: mortgages, Banks: banks})
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
	}
}

func (h *Handler) DeleteMortage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println(err)
		return
	}
	err = h.db.DeleteMortage(id)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/mortgages", 301)
}

func (h *Handler) AddMortgage(w http.ResponseWriter, r *http.Request) {
	InitialLoan, err := strconv.Atoi(r.FormValue("InitialLoan"))
	if err != nil {
		log.Println(err)
		return
	}
	DownPayment, err := strconv.Atoi(r.FormValue("DownPayment"))
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error"))
		return
	}
	BankId, err := strconv.Atoi(r.FormValue("BankId"))
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error"))
		return
	}
	bank, err := h.db.GetBank(BankId)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error"))
		return
	}
	if bank.MaxLoan < InitialLoan {
		log.Println(err)
		w.Write([]byte("Error"))
		return
	}
	if bank.MinDownPayment > DownPayment {
		log.Println(err)
		w.Write([]byte("Error"))
		return
	}
	R := float64(bank.InterestRate)
	R = (R / 100)
	N := float64(12) * 10
	var monthlypayment float64
	monthlypayment = (float64(InitialLoan) * (float64(R) / 12) * math.Pow(1+(float64(R)/12), N)) / (math.Pow(1+(float64(R)/12), N) - 1)
	err = h.db.CreateMortage(&model.Mortgage{InitialLoan: InitialLoan, DownPayment: DownPayment, MonthlyPayment: monthlypayment, BankId: BankId})
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	http.Redirect(w, r, "/mortgages", 301)
}
