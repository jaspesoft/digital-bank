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
		documentId   string
		patch        string
		documentType DocumentType
		documentSide DocumentSide
	}

	Stogare interface {
		UploadFile(file multipart.File, header *multipart.FileHeader) (string, error)
		DownloadFile(fileName string) (string, error)
	}
)

func NewDocument(documentId string, patch string, documentType DocumentType) *Document {
	return &Document{
		documentId:   documentId,
		patch:        patch,
		documentType: documentType,
	}
}

func (d *Document) GetDocumentId() string {
	return d.documentId
}

func (d *Document) GetPatch() string {
	return d.patch
}

func (d *Document) GetDocumentType() DocumentType {
	return d.documentType
}

func (d *Document) GetDocumentSide() DocumentSide {
	return d.documentSide
}

func (d *Document) UpdateDocument(doc Document) {
	d.documentSide = doc.documentSide
	d.patch = doc.patch
	d.documentType = doc.documentType
	d.documentSide = doc.documentSide

}

func (d *Document) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"documentId":   d.documentId,
		"patch":        d.patch,
		"documentType": d.documentType,
		"documentSide": d.documentSide,
	}
}
