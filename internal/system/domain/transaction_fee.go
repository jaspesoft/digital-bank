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
		} `json:"fedWire" bson:"fedWire"`
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
		DomesticUSA DomesticUSA `json:"domesticUsa" bson:"domesticUsa"`
		SwiftUSA    SwiftUSA    `json:"swiftUsa" bson:"swiftUsa"`
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
		"domesticUsa": f.DomesticUSA,
		"swiftUsa":    f.SwiftUSA,
		"swap":        f.Swap,
	}
}

func TransactionFeeFromPrimitive(t map[string]interface{}) *TransactionFee {
	achUSAMap := t["domesticUsa"].(map[string]interface{})["ach"].(map[string]interface{})
	domesticUSAMap := t["domesticUsa"].(map[string]interface{})["fedWire"].(map[string]interface{})
	swiftUsaMap := t["swiftUsa"].(map[string]interface{})
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
