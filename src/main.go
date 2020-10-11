package main

// Azure SDK
// go get -u -d github.com/Azure/azure-sdk-for-go/...
// Event Hub (kafka)
// go get -u -d github.com/Azure/azure-event-hubs-go

import (
	"fmt"
	"net/http"
)

// import
// "github.com/Azure/go-autorest/autorest/to" // not sure what this does yet

func main() {
	fmt.Println("Starting go server on port 80")
	http.HandleFunc("/", handler)
	//http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil)
	http.ListenAndServe(":80", nil)
}

func handler(responseWriter http.ResponseWriter, request *http.Request) {

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
