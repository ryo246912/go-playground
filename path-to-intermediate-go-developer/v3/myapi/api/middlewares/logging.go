package middlewares

import (
	"log"
	"net/http"
)

// 自作 ResponseWriter を作る
type resLoggingWriter struct {
	http.ResponseWriter // interfaceの埋め込み
	code                int
}

// コンストラクタを作る
func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

// WriteHeader メソッドを作る
func (rsw *resLoggingWriter) WriteHeader(code int) {
	rsw.code = code
	rsw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	// インターフェース型ということは、決められたメソッドさえ持てばどんな型でも渡すことができるということ
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// リクエスト情報をロギング
		log.Println(req.RequestURI, req.Method)
		// 自作の ResponseWriter を作って
		rlw := NewResLoggingWriter(w)

		next.ServeHTTP(rlw, req)

		log.Println("res: ", rlw.code)
	})
}
