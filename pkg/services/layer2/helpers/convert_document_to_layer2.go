package layer2helpers

import (
	accountdomain "digital-bank/internal/account/domain"
	layer2 "digital-bank/pkg/services/layer2"
)

func ConvertDocumentTypeToLayer2(docType accountdomain.DocumentType, side accountdomain.DocumentSide) layer2.Layer2DocumentType {

	switch docType {
	case accountdomain.UTILITY_BILL:
		return layer2.UTILITY_BILL

	case accountdomain.INCORPORATION_DOCUMENT:
		return layer2.INCORPORATION_DOCUMENTS

	case accountdomain.ARTICLE_OF_INCORPORATION:
		return layer2.ARTICLES_OF_INCORPORATION

	case accountdomain.ACCOUNT_AGREEMENT:
		if side == accountdomain.FRONT {
			return layer2.ACCOUNT_AGREEMENT
		}
		return layer2.LAYER2_ACCOUNT_AGREEMENT

	case accountdomain.BANK_STATEMENT:
		return layer2.BANK_STATEMENT

	case accountdomain.PAYSLIP:
		return layer2.PAYSLIP

	case accountdomain.BENEFICIAL_OWNERSHIP_CERTIFICATE:
		return layer2.BENEFICIAL_OWNERSHIP_CERTIFICATE

	case accountdomain.GOVERNMENT_ID:
		if side == accountdomain.FRONT {
			return layer2.IDENTITY_CARD_FRONT
		}
		return layer2.IDENTITY_CARD_BACK

	case accountdomain.PASSPORT:
		return layer2.PASSPORT

	case accountdomain.DRIVER_LICENSE:
		if side == accountdomain.FRONT {
			return layer2.DRIVERS_LICENCE_FRONT
		}
		return layer2.DRIVERS_LICENCE_BACK

	default:
		return ""
	}

}

func LookLayer2DocumentID(layer2Documents []layer2.ApplicationDocumentError, layer2DocType layer2.Layer2DocumentType) *string {
	for _, doc := range layer2Documents {
		if doc.Document == layer2DocType {
			return &doc.DocumentID
		}
	}

	return nil
}
