package systemdomain

type (
	DomesticUSA struct {
		ACH struct {
			IN  float64
			OUT float64
		}
		FedWire struct {
			IN  float64
			OUT float64
		}
	}

	SwiftUSA struct {
		IN  float64
		OUT float64
	}

	Swap struct {
		Buy  float64
		Sell float64
	}

	TransactionFee struct {
		domesticUSA DomesticUSA `json:"domesticUSA"`
		swiftUSA    SwiftUSA    `json:"swiftUSA"`
		swap        Swap
	}
)

func NewTransactionFee(domesticUSA DomesticUSA, swift SwiftUSA, swap Swap) *TransactionFee {
	return &TransactionFee{
		domesticUSA: domesticUSA,
		swiftUSA:    swift,
		swap:        swap,
	}
}

func (f *TransactionFee) GetFeeAchUSA() struct {
	IN  float64
	OUT float64
} {
	return f.domesticUSA.ACH
}

func (f *TransactionFee) GetFeeFedWire() struct {
	IN  float64
	OUT float64
} {
	return f.domesticUSA.FedWire
}

func (f *TransactionFee) GetSwiftUSA() struct {
	IN  float64
	OUT float64
} {
	return f.swiftUSA
}

func (f *TransactionFee) GetSwapFeeForBuy() float64 {
	return f.swap.Buy
}

func (f *TransactionFee) GetSwapFeeForSell() float64 {
	return f.swap.Sell
}

func (f *TransactionFee) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"domesticUSA": f.domesticUSA,
		"swiftUSA":    f.swiftUSA,
		"swap":        f.swap,
	}
}
