package systemdomain

type (
	DomesticUSA struct {
		ACH struct {
			IN  float64 `json:"in"`
			OUT float64 `json:"out"`
		} `json:"ach"`
		FedWire struct {
			IN  float64 `json:"in"`
			OUT float64 `json:"out"`
		} `json:"fedWire"`
	}

	SwiftUSA struct {
		IN  float64 `json:"in"`
		OUT float64 `json:"out"`
	}

	Swap struct {
		Buy  float64 `json:"buy"`
		Sell float64 `json:"sell"`
	}

	TransactionFee struct {
		DomesticUSA DomesticUSA `json:"domestic_usa" bson:"domesticUsa"`
		SwiftUSA    SwiftUSA    `json:"swift_usa" bson:"swiftUsa"`
		Swap        Swap        `json:"swap" bson:"swap"`
	}
)

func NewTransactionFee(domesticUSA DomesticUSA, swift SwiftUSA, swap Swap) *TransactionFee {
	return &TransactionFee{
		DomesticUSA: domesticUSA,
		SwiftUSA:    swift,
		Swap:        swap,
	}
}

func (f *TransactionFee) GetFeeAchUSA() struct {
	IN  float64
	OUT float64
} {
	return struct {
		IN  float64
		OUT float64
	}(f.DomesticUSA.ACH)
}

func (f *TransactionFee) GetFeeFedWire() struct {
	IN  float64
	OUT float64
} {
	return struct {
		IN  float64
		OUT float64
	}(f.DomesticUSA.FedWire)
}

func (f *TransactionFee) GetSwiftUSA() struct {
	IN  float64
	OUT float64
} {
	return struct {
		IN  float64
		OUT float64
	}(f.SwiftUSA)
}

func (f *TransactionFee) GetSwapFeeForBuy() float64 {
	return f.Swap.Buy
}

func (f *TransactionFee) GetSwapFeeForSell() float64 {
	return f.Swap.Sell
}

func (f *TransactionFee) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"domestic_usa": f.DomesticUSA,
		"swift_usa":    f.SwiftUSA,
		"swap":         f.Swap,
	}
}

func TransactionFeeFromPrimitive(t map[string]interface{}) *TransactionFee {

	achUSAMap := t["domestic_usa"].(map[string]interface{})["ach"].(map[string]interface{})
	domesticUSAMap := t["domestic_usa"].(map[string]interface{})["fedWire"].(map[string]interface{})
	swiftUsaMap := t["swift_usa"].(map[string]interface{})
	swapMap := t["swap"].(map[string]interface{})

	return &TransactionFee{
		DomesticUSA: DomesticUSA{
			ACH: struct {
				IN  float64 `json:"in"`
				OUT float64 `json:"out"`
			}(struct {
				IN  float64
				OUT float64
			}{
				IN:  achUSAMap["in"].(float64),
				OUT: achUSAMap["out"].(float64),
			}),
			FedWire: struct {
				IN  float64 `json:"in"`
				OUT float64 `json:"out"`
			}(struct {
				IN  float64
				OUT float64
			}{
				IN:  domesticUSAMap["in"].(float64),
				OUT: domesticUSAMap["out"].(float64),
			}),
		},
		SwiftUSA: SwiftUSA{
			IN:  swiftUsaMap["in"].(float64),
			OUT: swiftUsaMap["out"].(float64),
		},
		Swap: Swap{
			Buy:  swapMap["buy"].(float64),
			Sell: swapMap["sell"].(float64),
		},
	}
}
