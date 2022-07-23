package getmerchant

import (
	"api-majoo/shared/model"
)

func (h *Handler) getMerchantByUserID(id int) (data *model.MerchantModel, err error) {
	err = h.db.Table("Merchants").
		Where("user_id = ?", id).
		Find(&data).Error
	return data, err
}

func (h *Handler) getOutletByMerchantID(mercID int) (data *[]model.OutletsModel, err error) {
	err = h.db.Table("Outlets").
		Where("merchant_id = ?", mercID).
		Find(&data).Error

	return data, err
}

func (h *Handler) getTransactiopnsOutlete(mercID, outID int) (data *[]model.TransactionModels, err error) {

	err = h.db.Table("Transactions").
		Select("DATE_FORMAT(created_at, '%Y-%m-%d')as dates , count(bill_total) as qty, sum(bill_total)as bill_total").
		Where("merchant_id = ?", mercID).
		Where("outlet_id =?", outID).
		Group("DATE_FORMAT(created_at, '%Y-%m-%d')").
		Find(&data).Error

	return data, err
}
