package main

import (
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"

	"github.com/hyorimitsu/hello-temporal.io/app/constants"
	"github.com/hyorimitsu/hello-temporal.io/app/pkg/domain"
	"github.com/hyorimitsu/hello-temporal.io/app/pkg/workflow"
)

func main() {
	http.HandleFunc("/workflow", workflowHandler)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func workflowHandler(_ http.ResponseWriter, req *http.Request) {
	c, err := client.NewClient(client.Options{
		HostPort: os.Getenv("TEMPORAL_GRPC_ENDPOINT"),
	})
	if err != nil {
		log.Fatalln("unable to create temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "transfer-money-workflow",
		TaskQueue: constants.TransferMoneyTaskQueueName,
	}
	data := domain.TransferDetails{
		ID:     uuid.NewString(),
		Amount: 100,
		From:   "AAA",
		To:     "BBB",
	}

	wr, err := c.ExecuteWorkflow(req.Context(), options, workflow.TransferMoney, data)
	if err != nil {
		log.Fatalln("unable to execute workflow", err)
	}
	log.Printf("Workflow executed: [ID: %s, RunID: %s]\n", wr.GetID(), wr.GetRunID())
}
