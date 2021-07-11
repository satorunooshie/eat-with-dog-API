package errors

var (
	UserAlreadyExists    = newBadRequest("user_registration_error", "すでに登録済みです")
	UserNotFound         = newBadRequest("user_not_found", "ユーザーが見つかりません")
	UserNotActive        = newBadRequest("user_not_active", "ユーザーが見つかりません")
	UserNotPaid          = newUnauthorized("user_not_paid", "お支払い状況をご確認ください")
	UserParseDateOfBirth = newInternalServerError("user_parse_date_of_birth", "生年月日が不正です")
)
