hello-temporal.io
---

This is sample of [Temporal.io](https://temporal.io/), a simple, scalable open source way to write and run reliable cloud applications.

## Description

This is a sample implementation of a money transfer application with [Temporal.io](https://temporal.io/), based on an [official tutorial](https://docs.temporal.io/go/run-your-first-app-tutorial).

![architecture](https://github.com/hyorimitsu/hello-temporal.io/blob/main/doc/img/architecture.png)

## Structure

### Used language, tools, and other components

|language/tools|description|
|--------------|-----------|
|[Temporal.io](https://temporal.io/) |tool for a way to write and run reliable cloud applications|
|[Go](https://github.com/golang/go)  |programming language|
|[Skaffold](https://skaffold.dev/)   |tool for building, pushing and deploying your application|
|[Kubernetes](https://kubernetes.io/)|container orchestrator|

### Directories

```
.
├── app               # => worker & workflow application using temporal
│   ├── cmd
│   │   ├── worker
│   │   └── workflow
│   ├── pkg
│   │   ├── activity
│   │   ├── domain
│   │   └── workflow
│   └── (some omitted)
├── k8s               # => kubernetes manifests for worker & workflow application and temporal
│   ├── app
│   │   ├── worker
│   │   └── workflow
│   ├── temporal
│   │   ├── admintools
│   │   ├── cassandra
│   │   ├── server
│   │   └── web
│   └── (some omitted)
└── skaffold.yaml
```

## Usage

### Run and setup the application

1. Run the application in minikube.

    ```shell script
    # start minikube
    minikube start --profile hell-temporalio
    
    # treat any context as local
    skaffold config set --global local-cluster true
    
    # change the destination to the docker in minikube
    eval $(minikube -p hell-temporalio docker-env)
    
    # run application
    skaffold dev
    ```

2. Register "default" as namespace for temporal via admin-tools.

    ```shell script
    kubectl exec -it services/hello-temporalio-temporal-admintools --namespace temporal /bin/bash -- tctl namespace register default
    ```

### Run the worker

1. Run the worker.

    ```shell script
    curl -X POST http://localhost:9000/worker
    ```

2. Access the following URL and confirm that the page is displayed.

    http://localhost:8088

### Run the workflow

1. Run the workflow.

    ```shell script
    curl -X POST http://localhost:9001/workflow
    ```

2. Access the following URL again and confirm that the workflow is displayed.

    http://localhost:8088

### Simulate the activity bugs

1. Switch comment out [here](https://github.com/hyorimitsu/hello-temporal.io/blob/main/app/pkg/activity/activity.go#L29) and wait for the rebuild to complete.

2. Run the worker.

    ```shell script
    curl -X POST http://localhost:9000/worker
    ```

3. Run the workflow.

    ```shell script
    curl -X POST http://localhost:9001/workflow
    ```

    You will see that the workflow is pending and that the deposit is repeatedly executed according to the retry settings.

4. Revert comment out [here](https://github.com/hyorimitsu/hello-temporal.io/blob/main/app/pkg/activity/activity.go#L29) and wait for the rebuild to complete.

5. Run the worker.

    ```shell script
    curl -X POST http://localhost:9000/worker
    ```

    You will see that the failed activity is automatically executed.
    As this shows, bugs can be fixed "on the fly" without losing workflow status.

### Stop and delete the application

1. Stop and delete the application

    ```shell script
    # revert the destination to the local docker
    eval $(minikube -p hell-temporalio docker-env -u)
    
    # stop minikube
    minikube stop --profile hell-temporalio
    
    # delete minikube
    minikube delete --profile hell-temporalio
    ```

## Pending

The following are not yet supported.

- Elasticsearch: Integration with Temporal.io will allow [Advanced Visibility](https://docs.temporal.io/concepts/what-is-advanced-visibility/).
- Helm chart: Easier management and others of kubernetes. [Here](https://github.com/temporalio/helm-charts) is official sample.
