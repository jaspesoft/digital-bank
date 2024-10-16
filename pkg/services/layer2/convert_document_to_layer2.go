package layer2

import (
	accountdomain "digital-bank/internal/account/domain"
)

func ConvertDocumentTypeToLayer2(docType accountdomain.DocumentType, side accountdomain.DocumentSide) Layer2DocumentType {

	switch docType {
	case accountdomain.UTILITY_BILL:
		return UTILITY_BILL

	case accountdomain.INCORPORATION_DOCUMENT:
		return INCORPORATION_DOCUMENTS

	case accountdomain.ARTICLE_OF_INCORPORATION:
		return ARTICLES_OF_INCORPORATION

	case accountdomain.ACCOUNT_AGREEMENT:
		if side == accountdomain.FRONT {
			return ACCOUNT_AGREEMENT
		}
		return LAYER2_ACCOUNT_AGREEMENT

	case accountdomain.BANK_STATEMENT:
		return BANK_STATEMENT

	case accountdomain.PAYSLIP:
		return PAYSLIP

	case accountdomain.BENEFICIAL_OWNERSHIP_CERTIFICATE:
		return BENEFICIAL_OWNERSHIP_CERTIFICATE

	case accountdomain.GOVERNMENT_ID:
		if side == accountdomain.FRONT {
			return IDENTITY_CARD_FRONT
		}
		return IDENTITY_CARD_BACK

	case accountdomain.PASSPORT:
		return PASSPORT

	case accountdomain.DRIVER_LICENSE:
		if side == accountdomain.FRONT {
			return DRIVERS_LICENCE_FRONT
		}
		return DRIVERS_LICENCE_BACK

	default:
		return ""
	}

}

func LookLayer2DocumentID(layer2Documents []ApplicationDocumentError, layer2DocType Layer2DocumentType) *string {
	for _, doc := range layer2Documents {
		if doc.Document == layer2DocType {
			return &doc.DocumentID
		}
	}

	return nil
}
