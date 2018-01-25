package server

import (
	"context"
	"net/http"
	"net/url"

	"github.com/rancher/logging-api/api/setup"
	normanapi "github.com/rancher/norman/api"
	"github.com/rancher/norman/parse"
	"github.com/rancher/norman/types"
	loggingSchema "github.com/rancher/types/apis/logging.cattle.io/v3/schema"
	"k8s.io/client-go/rest"
)

func New(ctx context.Context, config *rest.Config) (http.Handler, error) {
	schemas := types.NewSchemas().
		AddSchemas(loggingSchema.Schemas)

	if err := setup.Schemas(ctx, config, schemas); err != nil {
		return nil, err
	}
	setup.AddDate()

	server := normanapi.NewAPIServer()
	server.URLParser = func(schemas *types.Schemas, url *url.URL) (parse.ParsedURL, error) {
		return URLParser(schemas, url)
	}
	// server.AccessControl = rbac.NewAccessControl(cluster.RBAC)
	// server.URLParser = func(schemas *types.Schemas, url *url.URL) (parse.ParsedURL, error) {
	// 	return URLParser(cluster.ClusterName, schemas, url)
	// }
	// server.Resolver = NewResolver(server.Resolver)
	// server.StoreWrapper = store.ProjectSetter(server.StoreWrapper)

	if err := server.AddSchemas(schemas); err != nil {
		return nil, err
	}

	return server, nil
}
