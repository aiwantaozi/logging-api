package setup

import (
	"context"

	"github.com/rancher/norman/store/crd"
	"github.com/rancher/norman/types"
	toolsschema "github.com/rancher/types/apis/tools.cattle.io/v3/schema"
	"github.com/rancher/types/client/management/v3"
	"k8s.io/client-go/rest"
)

// var (
// 	crdVersions = []*types.APIVersion{
// 		&toolsschema.Version,
// 	}
// )

func Schemas(ctx context.Context, config rest.Config, schemas *types.Schemas) error {
	// subscribe.Register(&toolsschema.Version, schemas)
	Logging(schemas)
	ProjectLogging(schemas)

	crdStore, err := crd.NewCRDStoreFromConfig(config)
	if err != nil {
		return err
	}
	var crdSchemas []*types.Schema
	if err := crdStore.AddSchemas(ctx, schemas); err != nil {
		return err
	}
}

func Logging(schemas *types.Schemas) {
	schema := schemas.Schema(&toolsschema.Version, client.Logging)
}

func ProjectLogging(schemas *types.Schemas) {
	schema := schemas.Schema(&toolsschema.Version, client.ProjectLogging)
}
