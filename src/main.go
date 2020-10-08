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

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You just browsed page (if blank you're at the root): %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+os.Getenv("HTTP_PLATFORM_PORT"), nil)
}
