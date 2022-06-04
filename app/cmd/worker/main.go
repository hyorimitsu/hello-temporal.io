package main

import (
	"log"
	"net/http"
	"os"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/hyorimitsu/hello-temporal.io/app/constants"
	"github.com/hyorimitsu/hello-temporal.io/app/pkg/activity"
	"github.com/hyorimitsu/hello-temporal.io/app/pkg/workflow"
)

func main() {
	http.HandleFunc("/worker", workerHandler)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func workerHandler(_ http.ResponseWriter, _ *http.Request) {
	c, err := client.NewClient(client.Options{
		HostPort: os.Getenv("TEMPORAL_GRPC_ENDPOINT"),
	})
	if err != nil {
		log.Fatalln("unable to create temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, constants.TransferMoneyTaskQueueName, worker.Options{})
	w.RegisterWorkflow(workflow.TransferMoney)
	w.RegisterActivity(activity.Withdraw)
	w.RegisterActivity(activity.Deposit)

	if err = w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("unable to run worker", err)
	}
}
