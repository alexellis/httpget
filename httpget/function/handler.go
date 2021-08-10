package function

import (
	"fmt"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {

	if r.Body != nil {
		defer r.Body.Close()

	}
	q := r.URL.Query().Get("q")
	if len(q) == 0 {
		http.Error(w, "Provide a URL with ?q=https://", http.StatusBadRequest)
		return
	}
	req, err := http.NewRequest(r.Method, q, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Status: %s", res.Status)))
}
