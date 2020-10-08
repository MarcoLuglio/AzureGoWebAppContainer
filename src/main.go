package azurego101

// Azure SDK
// go get -u -d github.com/Azure/azure-sdk-for-go/...
// Event Hub (kafka)
// go get -u -d github.com/Azure/azure-event-hubs-go

import (
	"fmt"
	"net/http"
	"os"
)

// import
// "github.com/Azure/go-autorest/autorest/to" // not sure what this does yet

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil)
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
	responseWriter.WriteHeader(500)
	responseWriter.Header().Set("Content-Type", "application/json")
	fmt.Print(responseWriter, `{"hello":"world"}`)
	return
}
