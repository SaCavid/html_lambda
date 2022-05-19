/*
	http://localhost:9000/2015-03-31/functions/function/invocations
*/

package main

import (
	"context"
	"github.com/SaCavid/router"
	"github.com/SaCavid/router/models"
	"html_lambda/handler"
	"log"
)

func HandleRequest(ctx context.Context, event models.LambdaRequest) (models.LambdaResponse, error) {

	r := router.NewLambdaRouter(ctx, &event)
	r.AllowedMethods("GET", "POST")

	srv := handler.Handler{
		Router: &r,
	}

	r.Handler("GET", "/path", srv.Meth)
	r.Handler("GET", "/html", srv.HTML)

	r.W.Set("Content-type", "application/json")

	return r.Run()
}

func main() {
	//lambda.Start(HandleRequest)
	event := models.LambdaRequest{}
	event.RequestContext.Http.Method = "GET"
	event.RequestContext.Http.Path = "/html"
	str, err := HandleRequest(context.Background(), event)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(str)
}
