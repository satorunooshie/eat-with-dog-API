package errors

const defaultErrorMessage = "エラーが発生しました"

var (
	SystemPanic      = newInternalServerError("system_panic", defaultErrorMessage)
	SystemDefault    = newInternalServerError("system_default", defaultErrorMessage)
	SystemUnknown    = newInternalServerError("system_unknown", defaultErrorMessage)
	InvalidParameter = newBadRequest("invalid_parameter", defaultErrorMessage)

	NotFound = newNotFound("not_found", "ページが見つかりませんでした")
)
