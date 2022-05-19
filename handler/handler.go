package handler

import (
	"context"
	"github.com/SaCavid/router"
	"github.com/SaCavid/router/models"
	"net/http"
)

type Handler struct {
	Router *router.Router
}

func (h *Handler) Meth(ctx context.Context, event models.LambdaRequest) (models.LambdaResponse, error) {

	return models.LambdaResponse{
		StatusCode: http.StatusCreated,
		Body:       "Hello from path url",
	}, nil
}

func (h *Handler) HTML(ctx context.Context, event models.LambdaRequest) (models.LambdaResponse, error) {
	body, err := h.Router.Execute("login.html", "./assets/login.html", nil)
	if err != nil {
		return *h.Router.W, err
	}

	h.Router.W.Body = body
	h.Router.W.StatusCode = http.StatusOK
	h.Router.W.Headers["Content-type"] = "text/html"
	return *h.Router.W, nil
}
