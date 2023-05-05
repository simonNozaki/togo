package error

// ApplicationError はエラー発生時に返却するエラーレスポンスの定義です
type ApplicationError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
