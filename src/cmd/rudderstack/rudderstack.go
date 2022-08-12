// package main

// import (
// 	analytics "github.com/segmentio/analytics-go/v3"
// )

// const (
// 	segmentWriteKey = "4yZtXCQmhm1taH7mZ5MTjp8v9zOI5kGk"
// )

// func main() {
// 	client := analytics.New(segmentWriteKey)
// 	defer client.Close()

// 	client.Enqueue(analytics.Identify{
// 		UserId: "019mr8mf4r",
// 		Traits: analytics.NewTraits().
// 			SetName("Michael Bolton").
// 			SetEmail("mbolton@example.com").
// 			Set("plan", "Enterprise").
// 			Set("friends", 42),
// 	})
// 	client.Enqueue(analytics.Track{
// 		UserId: "019mr8mf4r",
// 		Event:  "Signed Up",
// 		Properties: analytics.NewProperties().
// 			Set("plan", "Enterprise"),
// 	})
// }

package main

import (
	"github.com/rudderlabs/analytics-go"
)

func main() {
	// Instantiates a client to use send messages to the Rudder API.
	client := analytics.New("<REMOVED", "<REMOVED>")

	// complex := map[string]int{
	// 	"numBackups", 1
	// }

	// Enqueues a track event that will be sent asynchronously.
	client.Enqueue(analytics.Track{
		UserId:     "test-user",
		Event:      "test-snippet3",
		Properties: analytics.NewProperties().Set("simple", "backupHappened").Set("complex", map[string]int{"numBackups": 1}),
	})

	// Flushes any queued messages and closes the client.
	client.Close()
}
