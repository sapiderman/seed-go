package handlers

import (
	"net/http"

	"encoding/json"

	"github.com/sapiderman/seed-go/internal/helpers"
)

// CallbackPayload is the callback payload
type CallbackPayload struct {
	Amount            int     `json:"amount"`
	TransactionStatus string  `json:"transaction_status"`
	OrderID           string  `json:"order_id"`
	Message           string  `json:"message"`
	ShippingAddr      Address `json:"shipping_address"`
	TransactionTime   string  `json:"transaction_time"`
	TransactionID     string  `json:"transaction_id"`
	SignatureKey      string  `json:"signature_key"`
}

// Address data struct
type Address struct {
	Name        string
	Address     string
	City        string
	PostalCode  string
	Phone       string
	CountryCode string
}

// CallBackKredini implements kredini call back
func (h *Handlers) CallBackKredini(w http.ResponseWriter, r *http.Request) {
	logf := hLog.WithField("fn", "CallBackKredini()")

	callback := CallbackPayload{}
	err := json.NewDecoder(r.Body).Decode(&callback)
	if err != nil {
		logf.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = helpers.ValidateInput(r.Context(), callback)
	if err != nil {
		logf.Error(err)
		return
	}

	logf.Info("*** received: ", callback)

	helpers.HTTPResponseBuilder(r.Context(), w, r, http.StatusOK, "Kredini Callback", nil)
}

type allowed struct {
	Eligible        bool   `json:"eligible"`
	RejectionReason string `json:"rejection_reason"`
}

// TransactionEngine mocking transaction engine
func (h *Handlers) TransactionEngine(w http.ResponseWriter, r *http.Request) {
	logf := hLog.WithField("fn", "TransactionEngine()")

	callback := CallbackPayload{}
	err := json.NewDecoder(r.Body).Decode(&callback)
	if err != nil {
		logf.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// err = helpers.ValidateInput(r.Context(), callback)
	// if err != nil {
	// 	logf.Error(err)
	// 	return
	// }

	logf.Info("*** received: ", callback)

	data := allowed{
		Eligible:        true,
		RejectionReason: "too good",
	}

	helpers.HTTPResponseBuildPlain(r.Context(), w, r, http.StatusOK, "transaction inquiry", data)
}

// TransationEnquiry payload
type transationEnquiry struct {
	Consumer    ConsumerDetails `json:"consumer"`
	Checkout    CheckoutDetails `json:"checkout"`
	CreditLimit CreditDetail    `json:"credit_limit"`
	OSB         int             `json:"outstanding_balance"`
}

// CheckoutDetails describes what the checkout is for eligibility inquiry
type CheckoutDetails struct {
	Amount int    `json:"amount"`
	Type   string `json:"type"`
}

// CreditDetail for eligibility inquiry
type CreditDetail struct {
	MaxLimit       int `json:"maximum_limit"`
	RemainingLimit int `json:"remaining_limit"`
}
type ConsumerDetails struct {
	ConsumerID         int              `json:"id"`
	Name               string           `json:"name" validate:"required"`
	Phone              string           `json:"phone"`
	Email              string           `json:"email" validate:"required,email"`
	Gender             string           `json:"gender"`
	Birthdate          string           `json:"birthdate"`
	MaritalStatus      int              `json:"marital_status" validate:"required,numeric,min=1,max=4"`
	NumberOfDependents int              `json:"number_of_dependents" validate:"numeric,min=0,max=4"`
	MotherMaidenName   string           `json:"mother_maiden_name" validate:"required"`
	Subscription       int              `json:"subscription"`
	Address            string           `json:"address"`
	Occupation         Occupation       `json:"occupation" validate:"required"`
	EducationLevel     int              `json:"education_level" validate:"numeric,min=1,max=7"`
	EmergencyContact   EmergencyContact `json:"emergency_contact" validate:"required"`
	KycState           string           `json:"kyc_state"`
	NIK                string           `json:"nik"`
	Status             string           `json:"status"`
}

// Occupation contains occupation information
type Occupation struct {
	EmploymentType     int    `json:"employment_type" validate:"min=1,max=7"`
	CompanyName        string `json:"company_name"`
	Industry           string `json:"industry" validate:"oneof='Hospitality/Pelayanan' 'Kesehatan' 'Keuangan' 'Konstruksi' 'Pabrik' 'Pariwisata' 'Pertahanan/Militer' 'Pertambangan' 'Pertanian/Peternakan/Kehutanan/Perikanan' 'Properti' 'Teknologi Informasi' 'Transportasi' 'Lainnya'"`
	WorkAddress        string `json:"work_address"`
	EmploymentStatus   int    `json:"employment_status" validate:"min=1,max=2"`
	MonthlyIncomeRange int    `json:"monthly_income_range" validate:"min=1,max=8"`
	WorkExperience     int    `json:"work_experience" validate:"min=1,max=4"`
}

// EmergencyContact contains emergency contact information
type EmergencyContact struct {
	Name         string `json:"name" validate:"required"`
	Relationship int    `json:"relationship" validate:"required,numeric,min=1,max=4"`
	Phone        string `json:"phone" validate:"required"`
	Address      string `json:"address" validate:"required"`
}

// Onboarding mocking engine
func (h *Handlers) Onboarding(w http.ResponseWriter, r *http.Request) {
	logf := hLog.WithField("fn", "Onboarding()")

	payload := transationEnquiry{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		logf.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// err = helpers.ValidateInput(r.Context(), callback)
	// if err != nil {
	// 	logf.Error(err)
	// 	return

	logf.Info("*** received: ", payload)

	data := make(map[string]interface{})
	data["enabled"] = true
	data["credit_limit"] = map[string]int{"maximum_limit": 50000}

	helpers.HTTPResponseBuildPlain(r.Context(), w, r, http.StatusOK, "Onboarding OK", data)
}
