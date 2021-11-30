package graph

import (
	"log"
	"os"

	"github.com/northwesternmutual/grammes"
	"github.com/northwesternmutual/grammes/query/traversal"
)

var url string = "wss://" + os.Getenv("NEPTUNE_URL") + "/gremlin"
var client *grammes.Client
var g traversal.String

// Creates a new client
func InitialiseGraphClient() {
	gClient, err := grammes.DialWithWebSocket(url)
	client = gClient
	if err != nil {
		log.Fatalf("Error while creating client: %s\n", err.Error())
	}
	g = grammes.Traversal()
}

func testGremlin() {
	// Executing a basic query to assure that the client is working.
	res, err := client.ExecuteStringQuery("g.V().groupCount().by(label).unfold()")
	if err != nil {
		log.Fatalf("Querying error: %s\n", err.Error())
	}

	// Print out the result as a string
	for _, r := range res {
		log.Println(string(r))
	}
}
