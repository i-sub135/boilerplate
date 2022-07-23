package login

import "api-majoo/shared/model"

func (h *Handler) getUserByUsernamePassword(username, password string) (data *model.UsersModel, err error) {
	err = h.db.Table("Users").
		Where("user_name = ?", username).
		Where("password = ?", password).
		Find(&data).Error
	return data, err
}
