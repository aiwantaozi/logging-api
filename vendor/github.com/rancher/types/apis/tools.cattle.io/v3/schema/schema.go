package schema

import (
	"github.com/rancher/norman/types"
	"github.com/rancher/types/apis/tools.cattle.io/v3"
	"github.com/rancher/types/factory"
)

var (
	Version = types.APIVersion{
		Version: "v3",
		Group:   "tools.cattle.io",
		Path:    "/v3",
	}

	Schemas = factory.Schemas(&Version).
		Init(loggingTypes)
)

func loggingTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		MustImport(&Version, v3.Logging{}).
		MustImport(&Version, v3.ProjectLogging{})
}
