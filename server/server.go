package server

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/aiwantaozi/tools-api/api/setup"
	normanapi "github.com/rancher/norman/api"
	"github.com/rancher/norman/parse"
	"github.com/rancher/norman/types"
	toolsSchema "github.com/rancher/types/apis/tools.cattle.io/v4/schema"
	"k8s.io/client-go/rest"
)

func New(ctx context.Context, config *rest.Config) (http.Handler, error) {
	schemas := types.NewSchemas().
		AddSchemas(toolsSchema.Schemas)

	if err := setup.Schemas(ctx, config, schemas); err != nil {
		return nil, err
	}

	server := normanapi.NewAPIServer()
	server.URLParser = func(schemas *types.Schemas, url *url.URL) (parse.ParsedURL, error) {
		fmt.Println("here: %s", url)
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
