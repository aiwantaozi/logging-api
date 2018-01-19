package setup

import (
	"context"

	"github.com/rancher/norman/store/crd"
	"github.com/rancher/norman/types"
	toolsschema "github.com/rancher/types/apis/tools.cattle.io/v4/schema"
	toolsclient "github.com/rancher/types/client/tools/v4"
	"k8s.io/client-go/rest"
)

func Schemas(ctx context.Context, config *rest.Config, schemas *types.Schemas) error {
	Logging(schemas)
	ProjectLogging(schemas)

	crdStore, err := crd.NewCRDStoreFromConfig(*config)
	if err != nil {
		return err
	}

	var crdSchemas []*types.Schema
	for _, schema := range schemas.SchemasForVersion(toolsschema.Version) {
		crdSchemas = append(crdSchemas, schema)
	}

	if err := crdStore.AddSchemas(ctx, crdSchemas...); err != nil {
		return err
	}
	return nil
}

func Logging(schemas *types.Schemas) {
	schemas.Schema(&toolsschema.Version, toolsclient.LoggingType)
}

func ProjectLogging(schemas *types.Schemas) {
	schemas.Schema(&toolsschema.Version, toolsclient.ProjectLoggingType)
}
