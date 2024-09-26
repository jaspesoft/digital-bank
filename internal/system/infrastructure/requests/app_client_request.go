package systemreq

import systemdomain "digital-bank/internal/system/domain"

type (
	Commissions struct {
		DomesticUSA systemdomain.DomesticUSA `json:"domestic_usa" binding:"required"`
		SwiftUSA    systemdomain.SwiftUSA    `json:"swift_usa" binding:"required"`
		Swap        systemdomain.Swap        `json:"swap" binding:"required"`
	}

	AppClientRequest struct {
		Name        string      `json:"companyName" binding:"required"`
		Email       string      `json:"email" binding:"required"`
		PhoneNumber string      `json:"phoneNumber" binding:"required"`
		Commissions Commissions `json:"commissions" binding:"required"`
	}
)
