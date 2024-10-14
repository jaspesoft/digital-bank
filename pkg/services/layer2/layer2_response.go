package layer2

const (
	individual  layer2ApplicationType = "INDIVIDUAL"
	corporation layer2ApplicationType = "CORPORATION"
)

const (
	INCOMPLETE           layer2ApplicationStatus = "INCOMPLETE"
	READY_FOR_SUBMISSION layer2ApplicationStatus = "READY_FOR_SUBMISSION"
	SUBMITTED            layer2ApplicationStatus = "SUBMITTED"
	CHANGES_REQUESTED    layer2ApplicationStatus = "CHANGES_REQUESTED"
	APPROVED             layer2ApplicationStatus = "APPROVED"
	REJECTED             layer2ApplicationStatus = "REJECTED"
	COMPLETED            layer2ApplicationStatus = "COMPLETED"
)

const (
	PASSPORT                         Layer2DocumentType = "PASSPORT"
	DRIVERS_LICENCE_FRONT            Layer2DocumentType = "DRIVERS_LICENCE_FRONT"
	DRIVERS_LICENCE_BACK             Layer2DocumentType = "DRIVERS_LICENCE_BACK"
	LAYER2_ACCOUNT_AGREEMENT         Layer2DocumentType = "LAYER2_ACCOUNT_AGREEMENT"
	ACCOUNT_AGREEMENT                Layer2DocumentType = "ACCOUNT_AGREEMENT"
	LEASE_AGREEMENT                  Layer2DocumentType = "LEASE_AGREEMENT"
	INCORPORATION_DOCUMENTS          Layer2DocumentType = "INCORPORATION_DOCUMENTS"
	ARTICLES_OF_INCORPORATION        Layer2DocumentType = "ARTICLES_OF_INCORPORATION"
	BENEFICIAL_OWNERSHIP_CERTIFICATE Layer2DocumentType = "BENEFICIAL_OWNERSHIP_CERTIFICATE"
	IDENTITY_CARD_FRONT              Layer2DocumentType = "IDENTITY_CARD_FRONT"
	IDENTITY_CARD_BACK               Layer2DocumentType = "IDENTITY_CARD_BACK"
	UTILITY_BILL                     Layer2DocumentType = "UTILITY_BILL"
	BANK_STATEMENT                   Layer2DocumentType = "BANK_STATEMENT"
	PAYSLIP                          Layer2DocumentType = "PAYSLIP"
)

type (
	layer2ApplicationType string

	layer2ApplicationStatus string

	Layer2DocumentType string

	applicationResponse struct {
		data struct {
			id string `json:"id"`
		} `json:"data"`
	}

	layer2ApplicationStatusResponse struct {
		id                          string                        `json:"id"`
		status                      layer2ApplicationStatus       `json:"status"`
		termsAndConditionsAccepted  bool                          `json:"terms_and_conditions_accepted"`
		customerID                  string                        `json:"customer_id"`
		applicationErrors           []string                      `json:"application_errors,omitempty"`
		applicationValidationErrors []validationError             `json:"application_validation_errors"`
		applicationDocumentErrors   []ApplicationDocumentError    `json:"application_document_errors"`
		individualErrors            []individualError             `json:"individual_errors,omitempty"`
		applicationIndividualErrors []applicationIndividualErrors `json:"application_individual_errors,omitempty"`
		validationErrors            []validationError             `json:"validation_errors"`
	}

	validationError struct {
		fieldName   string `json:"field_name"`
		fieldStatus string `json:"field_status"`
		description string `json:"description"`
	}

	ApplicationDocumentError struct {
		DocumentID  string             `json:"document_id"`
		Status      string             `json:"status"`
		Description string             `json:"description,omitempty"`
		Document    Layer2DocumentType `json:"document"`
	}

	applicationIndividualErrors struct {
		individualID     string                      `json:"individual_id"`
		validationErrors []individualValidationError `json:"validation_errors"`
		documentErrors   []ApplicationDocumentError  `json:"document_errors"`
	}

	individualError struct {
		id                         string                     `json:"id"`
		individualValidationErrors []validationError          `json:"individual_validation_errors"`
		individualDocumentErrors   []ApplicationDocumentError `json:"individual_document_errors"`
	}

	individualValidationError struct {
		fieldName   string `json:"field_name"`
		fieldStatus string `json:"field_status"`
		description string `json:"description"`
	}
)
