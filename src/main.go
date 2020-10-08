package main

// Azure SDK
// go get -u -d github.com/Azure/azure-sdk-for-go/...
// Event Hub (kafka)
// go get -u -d github.com/Azure/azure-event-hubs-go

import (
	"net/http"
)

// import
// "github.com/Azure/go-autorest/autorest/to" // not sure what this does yet

func main() {
	http.HandleFunc("/", handler)
	//http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil)
	http.ListenAndServe(":8080", nil)
}

func handler(responseWriter http.ResponseWriter, request *http.Request) {

	var err error

	switch request.Method {
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

func handleGet(responseWriter http.ResponseWriter, request *http.Request) (err error) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(200)
	responseWriter.Write([]byte(`{"hello":"world"}`))
	return
}
