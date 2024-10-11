package accountdomain

import "mime/multipart"

const (
	SELFIE                           DocumentType = "selfie"
	PASSPORT                         DocumentType = "passport"
	DRIVER_LICENSE                   DocumentType = "driver_license"
	GOVERNMENT_ID                    DocumentType = "government_id"
	UTILITY_BILL                     DocumentType = "utility_bill"
	BANK_STATEMENT                   DocumentType = "bank_statement"
	ACCOUNT_AGREEMENT                DocumentType = "account_agreement"
	BENEFICIAL_OWNERSHIP_CERTIFICATE DocumentType = "beneficial_ownership_certificate"
	INCORPORATION_DOCUMENT           DocumentType = "incorporation_document"
	ARTICLE_OF_INCORPORATION         DocumentType = "article_of_incorporation"
	W2                               DocumentType = "w2"
	PAYSLIP                          DocumentType = "payslip"

	FRONT DocumentSide = "front"
	BACK  DocumentSide = "back"
)

type (
	DocumentType string
	DocumentSide string

	Document struct {
		AccountID    string       `json:"accountId"`
		Patch        string       `json:"patch"`
		DocumentType DocumentType `json:"documentType"`
		DocumentSide DocumentSide `json:"documentSide"`
	}

	Stogare interface {
		UploadFile(file multipart.File, header *multipart.FileHeader) (string, error)
		DownloadFile(fileName string) (string, error)
	}
)

func NewDocument(accountID string, patch string, docType DocumentType, docSide DocumentSide) *Document {
	return &Document{
		AccountID:    accountID,
		Patch:        patch,
		DocumentType: docType,
		DocumentSide: docSide,
	}
}

func (d *Document) GetPatch() string {
	return d.Patch
}

func (d *Document) GetDocumentType() DocumentType {
	return d.DocumentType
}

func (d *Document) GetDocumentSide() DocumentSide {
	return d.DocumentSide
}

func (d *Document) UpdateDocument(doc Document) {
	d.DocumentSide = doc.DocumentSide
	d.Patch = doc.Patch
	d.DocumentType = doc.DocumentType
	d.DocumentSide = doc.DocumentSide

}

func (d *Document) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"accountId":    d.AccountID,
		"patch":        d.Patch,
		"documentType": d.DocumentType,
		"documentSide": d.DocumentSide,
	}
}

func DocumentFromPrimitives(data map[string]interface{}) *Document {
	return &Document{
		AccountID:    data["accountId"].(string),
		Patch:        data["patch"].(string),
		DocumentType: data["documentType"].(DocumentType),
		DocumentSide: data["documentSide"].(DocumentSide),
	}
}
