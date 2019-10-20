package client

// Account struct
type Account struct {
	ID          string `json:"id,omitempty"`
	IBAN        string `json:"iban,omitempty"`
	MSISDN      string `json:"msisdn,omitempty"`
	Currency    string `json:"currency,omitempty"`
	Name        string `json:"name,omitempty"`
	AccountType string `json:"accountType,omitempty"`
	BIC         string `json:"bic,omitempty"`
}

// Aspsp struct
type Aspsp struct {
	ID           string `json:"id,omitempty"`
	BIC          string `json:"bic,omitempty"`
	BankCode     string `json:"bank-code,omitempty"`
	ASPSPCDE     string `json:"aspsp-cde,omitempty"`
	Name         string `json:"name,omitempty"`
	LogoLocation string `json:"logoLocation,omitempty"`
	APIList      []struct {
		Consents                   []string `json:"consents,omitempty"`
		Payments                   []string `json:"payments,omitempty"`
		Accounts                   []string `json:"accounts,omitempty"`
		Balances                   []string `json:"balances,omitempty"`
		Transaction                []string `json:"transaction,omitempty"`
		FundsConfirmations         []string `json:"funds-confirmations,omitempty"`
		PeriodicPayments           []string `json:"periodic-payments,omitempty"`
		MultibancoPayments         []string `json:"multibanco-payments,omitempty"`
		PeriodicMultibancoPayments []string `json:"periodic-multibanco-payments,omitempty"`
	} `json:"api-list,omitempty"`
}

// Balance struct
type Balance struct {
	Ammout struct {
		Currency string `json:"currency,omitempty"`
		Content  string `json:"content,omitempty"`
	} `json:"ammout,omitempty"`
	LastActionDateTime string `json:"lastActionDateTime,omitempty"`
	Date               string `json:"date,omitempty"`
}

// Transaction struct
type Transaction struct {
	TransactionId string `json:"transactionId,omitempty"`
	CreditorName  string `json:"creditorName,omitempty"`
	Amount        struct {
		Currency string `json:"currency,omitempty"`
		Content  string `json:"content,omitempty"`
	} `json:"amount,omitempty"`
	BookingDate                       string `json:"bookingDate,omitempty"`
	ValueDate                         string `json:"valueDate,omitempty"`
	RemittanceInformationUnstructured string `json:"remittanceInformationUnstructured,omitempty"`
}

// ConsentPayload struct
type ConsentPayload struct {
	Access struct {
		Accounts          []AccountReference `json:"accounts,omitempty"`
		Balances          []AccountReference `json:"balances,omitempty"`
		Transactions      []AccountReference `json:"transactions,omitempty"`
		AvailableAccounts string             `json:"availableAccounts,omitempty"`
		AllPSD2           string             `json:"allPsd2,omitempty"`
	} `json:"access,omitempty"`
	RecurringIndicator       bool   `json:"recurringIndicator"`
	ValidUntil               string `json:"validUntil"`
	FrequencyPerDay          int    `json:"frequencyPerDay"`
	CombinedServiceIndicator bool   `json:"combinedServiceIndicator"`
}

// AccountReference struct
type AccountReference struct {
	IBAN      string `json:"iban,omitempty"`
	BBAN      string `json:"bban,omitempty"`
	PAN       string `json:"pan,omitempty"`
	MaskedPan string `json:"maskedPan,omitempty"`
	MSISDN    string `json:"msisdn,omitempty"`
	Currency  string `json:"currency,omitempty"`
}

// ConsentResponseResource struct
type ConsentResponseResource struct {
	TransactionStatus string `json:"transactionStatus,omitempty"`
	ConsentID         string `json:"consentId,omitempty"`
	// ScaMethods        []interface{} `json:"scaMethods,omitempty"`
	// ChosenScaMethod interface{} `json:"chosenScaMethod,omitempty"`
	// ChallengeData   interface{}      `json:"challengeData,omitempty"`
	Links struct {
		Redirect                   string `json:"redirect,omitempty"`
		UpdatePsuIdentification    string `json:"updatePsuIdentification,omitempty"`
		UpdatePsuAuthenication     string `json:"updatePsuAuthenication,omitempty"`
		SelectAuthenticationMethod string `json:"selectAuthenticationMethod,omitempty"`
		AuthoriseTransaction       string `json:"authoriseTransaction,omitempty"`
		Status                     string `json:"status,omitempty"`
	} `json:"_links,omitempty"`
	PsuMessage string `json:"psuMessage,omitempty"`
}
