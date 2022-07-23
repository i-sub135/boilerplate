package getmerchant

import "api-majoo/shared/model"

type (
	ResponseGenerator struct {
		Merchant Merchant  `json:"Merchant"`
		Outlets  []Outlets `json:"outlets"`
	}
	Merchant struct {
		UserID       int    `json:"UserID"`
		MerchantName string `json:"MerchantName"`
	}
	Outlets struct {
		MerchantID   int            `json:"MerchantID"`
		OutletName   string         `json:"OutletName"`
		Transactions []Transactions `json:"Transactions"`
	}
	Transactions struct {
		Date      string `json:"Date"`
		Qty       int    `json:"Qty"`
		BillTotal int    `json:"BillTotal"`
	}
)

func (h *Handler) toResponse(merchant *model.MerchantModel, outlet *[]model.OutletsModel) ResponseGenerator {

	var out []Outlets

	for _, outletsModel := range *outlet {
		out = append(out, Outlets{
			MerchantID:   outletsModel.MerchantID,
			OutletName:   outletsModel.OutletName,
			Transactions: h.toGetTransactions(outletsModel.MerchantID, outletsModel.ID),
		})
	}

	return ResponseGenerator{
		Merchant: Merchant{
			UserID:       merchant.UserID,
			MerchantName: merchant.MerchantName,
		},
		Outlets: out,
	}
}

func (h *Handler) toGetTransactions(merID, outID int) []Transactions {
	trans, _ := h.getTransactiopnsOutlete(merID, outID)
	var t []Transactions
	for _, trx := range *trans {
		t = append(t, Transactions{
			Date:      trx.Dates,
			Qty:       trx.Qty,
			BillTotal: trx.BillTotal,
		})
	}

	return t
}
