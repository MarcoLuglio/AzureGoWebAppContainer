package main

// Azure SDK
// go get -u -d github.com/Azure/azure-sdk-for-go/...
// Event Hub (kafka)
// go get -u -d github.com/Azure/azure-event-hubs-go

import (
	"fmt"
	"net/http"
	"time"
)

// import
// "github.com/Azure/go-autorest/autorest/to" // not sure what this does yet

func main() {

	fmt.Println("Starting go server on port 80")

	// check https://medium.com/@simonfrey/go-as-in-golang-standard-net-http-config-will-break-your-production-environment-1360871cb72b
	handler := timeoutHandler{}
	timeoutServer := &http.Server{
		ReadHeaderTimeout: 20 * time.Second,
		ReadTimeout:       1 * time.Minute,
		WriteTimeout:      2 * time.Minute,
		Handler:           handler,
		Addr:              ":80",
	}
	timeoutServer.ListenAndServe()

}

// #region timeoutHandler struct

type timeoutHandler struct{}

func (handler timeoutHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {

	defer request.Body.Close()

	switch request.RequestURI {
	case "/":
		// OK
	default:
		// TODO 404
		return
	}

	/*
		// this is how to do a varible timeout to read the request, but it will also eat part of the write timeout
		timer := time.AfterFunc(5*time.Second, func() {
			request.Body.Close()
		})

		// if more time is needed to write the response, reset the timer here, os in another method
		timer.Reset(1 * time.Second)
	*/

	var err error

	switch request.Method {
	case "OPTIONS":
		err = handleOptions(responseWriter, request)
	case "GET":
		err = handleGet(responseWriter, request)
		/*case "POST":
			err = handleVideoAPIPost(writer, request, video)
		case "PUT":
			err = handleVideoAPIPut(writer, request, video)
		case "DELETE":
			err = handleVideoAPIDelete(writer, request, video)*/
	}

	if err != nil {
		//util.CheckError(err)
		//return
	}

}

// #endregion

func handleOptions(responseWriter http.ResponseWriter, request *http.Request) (err error) {
	responseWriter.Header().Set("access-control-allow-origin", "*") // posso usar *, mas não se requerir autenticação com bearer:
	responseWriter.Header().Set("access-control-allow-methods", "options, get, post, put, delete")
	responseWriter.Header().Set("access-control-allow-headers", "content-type") // posso usar *
	return
}

func handleGet(responseWriter http.ResponseWriter, request *http.Request) (err error) {
	responseWriter.Header().Set("content-type", "application/json")
	responseWriter.WriteHeader(200)
	responseWriter.Write([]byte(`{"hello":"world"}`))
	return
}
