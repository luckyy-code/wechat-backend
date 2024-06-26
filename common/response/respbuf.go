package response

import (
	"encoding/json"
	"net/http"
)

func Ok(w http.ResponseWriter, status int) {
	m := make(map[string]interface{})
	m["code"] = SUCCESS
	m["message"] = SUCCESS_MESSAGE
	buf, _ := json.Marshal(m)

	w.WriteHeader(status)
	w.Write(buf)
}

func Fail(code int, message string, w http.ResponseWriter, status int) {
	m := make(map[string]interface{})
	m["code"] = code
	m["message"] = message
	buf, _ := json.Marshal(m)

	w.WriteHeader(status)
	w.Write(buf)
}

func DefaultFail(w http.ResponseWriter, status int) {
	Fail(FAIL, SYSTEM_FAIL_MESSAGE, w, status)
}
