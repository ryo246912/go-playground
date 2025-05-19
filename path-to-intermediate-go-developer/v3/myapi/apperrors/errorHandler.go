package apperrors

import (
	"encoding/json"
	"errors"
	"net/http"
)

// エラーが発生したときのレスポンス処理をここで一括で行う
func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	var appErr *MyAppError
	// 重要なポイントとしては、「エラーハンドラの第三引数で受け取っている err の型は error インターフェースであって、
	// 具体型である MyAppError 構造体ではない」ということです。
	// そのため、第三引数で受け取った err を MyAppError 構造体とみなして、
	// 内部の ErrCode フィールド等を取り出したいのであれば、
	// エラーインターフェース型である変数 err を MyAppError 型に変換してやる必要

	// errors.As 関数で引数の err を MyAppError 型の appErr に変換する
	if !errors.As(err, &appErr) {
		// もし変換に失敗したら Unknown エラーを変数 appErr に手動で格納
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	var statusCode int
	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)

}
