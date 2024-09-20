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

func TransactionFeeFromPrimitive(t map[string]interface{}) *TransactionFee {

	return &TransactionFee{
		domesticUSA: DomesticUSA{
			ACH: struct {
				IN  float64
				OUT float64
			}{
				IN:  t["domesticUSA"].(map[string]interface{})["ACH"].(map[string]interface{})["IN"].(float64),
				OUT: t["domesticUSA"].(map[string]interface{})["ACH"].(map[string]interface{})["OUT"].(float64),
			},
			FedWire: struct {
				IN  float64
				OUT float64
			}{
				IN:  t["domesticUSA"].(map[string]interface{})["FedWire"].(map[string]interface{})["IN"].(float64),
				OUT: t["domesticUSA"].(map[string]interface{})["FedWire"].(map[string]interface{})["OUT"].(float64),
			},
		},
		swiftUSA: SwiftUSA{
			IN:  t["swiftUSA"].(map[string]interface{})["IN"].(float64),
			OUT: t["swiftUSA"].(map[string]interface{})["OUT"].(float64),
		},
		swap: Swap{
			Buy:  t["swap"].(map[string]interface{})["Buy"].(float64),
			Sell: t["swap"].(map[string]interface{})["Sell"].(float64),
		},
	}
}
