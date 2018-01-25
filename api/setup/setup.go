package setup

import (
	"context"

	loggingapi "github.com/rancher/logging-api/api/logging"
	"github.com/rancher/norman/store/crd"
	"github.com/rancher/norman/types"
	loggingschema "github.com/rancher/types/apis/logging.cattle.io/v3/schema"
	loggingclient "github.com/rancher/types/client/logging/v3"
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
	for _, schema := range schemas.SchemasForVersion(loggingschema.Version) {
		crdSchemas = append(crdSchemas, schema)
	}

	if err := crdStore.AddSchemas(ctx, crdSchemas...); err != nil {
		return err
	}
	return nil
}

func Logging(schemas *types.Schemas) {
	schema := schemas.Schema(&loggingschema.Version, loggingclient.LoggingType)
	schema.Validator = loggingapi.Validator
}

func ProjectLogging(schemas *types.Schemas) {
	schema := schemas.Schema(&loggingschema.Version, loggingclient.ProjectLoggingType)
	schema.Validator = loggingapi.Validator
}
