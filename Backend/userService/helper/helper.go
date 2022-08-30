package helper

import (
	"context"
	"encoding/json"
	"userService/poststore"
	tracer "userService/tracer"

	//"userService/poststore"
	"io"
	"net/http"
)

// renderJSON renders 'v' as JSON and writes it as a response into w.
func RenderJSON(ctx context.Context, w http.ResponseWriter, v interface{}) {
	span := tracer.StartSpanFromContext(ctx, "renderJSON")
	defer span.Finish()

	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func DecodeBody(ctx context.Context, r io.Reader) (*poststore.RequestPost, error) {
	span := tracer.StartSpanFromContext(ctx, "decodeBody")
	defer span.Finish()

	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	var rt poststore.RequestPost
	if err := dec.Decode(&rt); err != nil {
		return nil, err
	}
	return &rt, nil
}
