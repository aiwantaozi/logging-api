package main

import (
	"context"
	"os"

	"fmt"
	"net/http"

	"github.com/rancher/logging-api/server"
	"github.com/rancher/types/config"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	if err != nil {
		return err
	}

	logCtx, err := config.NewLoggingContext(*kubeConfig)
	if err != nil {
		return err
	}

	handler, err := server.New(context.Background(), *logCtx, kubeConfig)
	if err != nil {
		return err
	}

	fmt.Println("Listening on 0.0.0.0:1235")
	return http.ListenAndServe("0.0.0.0:1235", handler)
}
