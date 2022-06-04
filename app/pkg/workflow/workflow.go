package workflow

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"

	"github.com/hyorimitsu/hello-temporal.io/app/pkg/activity"
	"github.com/hyorimitsu/hello-temporal.io/app/pkg/domain"
)

func TransferMoney(ctx workflow.Context, transferDetails domain.TransferDetails) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
			MaximumInterval:    time.Minute,
			MaximumAttempts:    500,
		},
	})
	if err := workflow.ExecuteActivity(ctx, activity.Withdraw, transferDetails).Get(ctx, nil); err != nil {
		return err
	}
	if err := workflow.ExecuteActivity(ctx, activity.Deposit, transferDetails).Get(ctx, nil); err != nil {
		return err
	}
	return nil
}
