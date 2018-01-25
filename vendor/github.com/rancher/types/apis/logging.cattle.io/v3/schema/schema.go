package schema

import (
	"github.com/rancher/norman/types"
	m "github.com/rancher/norman/types/mapper"
	"github.com/rancher/types/apis/logging.cattle.io/v3"
	"github.com/rancher/types/factory"
	"github.com/rancher/types/mapper"
)

var (
	Version = types.APIVersion{
		Version: "v3",
		Group:   "logging.cattle.io",
		Path:    "/v3",
	}

	Schemas = factory.Schemas(&Version).
		Init(loggingTypes)
	minPort = int64(1)
	maxPort = int64(65535)
)

func loggingTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v3.Logging{}, m.DisplayName{}).
		AddMapperForType(&Version, v3.ProjectLogging{}, m.DisplayName{}).
		AddMapperForType(&Version, v3.ProjectLogging{}, &mapper.NamespaceIDMapper{}).
		MustImportAndCustomize(&Version, v3.ElasticsearchConfig{}, func(schema *types.Schema) {
			schema.MustCustomizeField("host", func(field types.Field) types.Field {
				field.Required = true
				field.Nullable = false
				return field
			})
			schema.MustCustomizeField("port", func(field types.Field) types.Field {
				field.Required = true
				field.Nullable = false
				field.Default = 9200
				field.Min = &minPort
				field.Max = &maxPort
				return field
			})
			schema.MustCustomizeField("indexPrefix", func(field types.Field) types.Field {
				field.Required = true
				field.Nullable = false
				return field
			})
			schema.MustCustomizeField("dateFormat", func(field types.Field) types.Field {
				field.Required = true
				field.Nullable = false
				field.Type = "enum"
				field.Options = []string{"YYYY-MM-DD", "YYYY-MM", "YYYY"}
				field.Default = "YYYY-MM-DD"
				return field
			})
		}).
		MustImportAndCustomize(&Version, v3.SplunkConfig{}, func(schema *types.Schema) {
			schema.MustCustomizeField("host", func(field types.Field) types.Field {
				field.Required = true
				field.Nullable = false
				return field
			})
			schema.MustCustomizeField("port", func(field types.Field) types.Field {
				field.Required = true
				field.Nullable = false
				field.Default = 8088
				field.Min = &minPort
				field.Max = &maxPort
				return field
			})
			schema.MustCustomizeField("protocol", func(field types.Field) types.Field {
				field.Type = "enum"
				field.Options = []string{"http", "tls"}
				field.Default = "http"
				return field
			})
			schema.MustCustomizeField("token", func(field types.Field) types.Field {
				field.Required = true
				field.Nullable = false
				return field
			})
		}).
		MustImportAndCustomize(&Version, v3.KafkaConfig{}, func(schema *types.Schema) {
			schema.MustCustomizeField("topic", func(field types.Field) types.Field {
				field.Required = true
				field.Nullable = false
				return field
			})
			schema.MustCustomizeField("dataType", func(field types.Field) types.Field {
				field.Required = true
				field.Nullable = false
				field.Default = "json"
				return field
			})
			schema.MustCustomizeField("maxSendRetries", func(field types.Field) types.Field {
				field.Required = true
				field.Nullable = false
				field.Default = 1
				max := int64(5)
				min := int64(0)
				field.Max = &max
				field.Min = &min
				return field
			})
			schema.MustCustomizeField("maxSendRetries", func(field types.Field) types.Field {
				field.Default = 1
				return field
			})
			schema.MustCustomizeField("dataType", func(field types.Field) types.Field {
				field.Type = "enum"
				field.Options = []string{"json", "ltsv", "msgpack"}
				field.Default = "json"
				return field
			})
		}).
		MustImportAndCustomize(&Version, v3.Zookeeper{}, func(schema *types.Schema) {
			schema.MustCustomizeField("host", func(field types.Field) types.Field {
				field.Required = true
				field.Nullable = false
				return field
			})
			schema.MustCustomizeField("port", func(field types.Field) types.Field {
				field.Required = true
				field.Nullable = false
				field.Default = 2181
				field.Min = &minPort
				field.Max = &maxPort
				return field
			})
		}).
		MustImportAndCustomize(&Version, v3.BrokerList{}, func(schema *types.Schema) {
			schema.MustCustomizeField("brokerList", func(field types.Field) types.Field {
				field.Required = true
				field.Nullable = false
				maxLen := int64(100)
				field.MaxLength = &maxLen
				return field
			})
		}).
		MustImportAndCustomize(&Version, v3.SyslogConfig{}, func(schema *types.Schema) {
			schema.MustCustomizeField("port", func(field types.Field) types.Field {
				field.Default = 51400
				return field
			})
			schema.MustCustomizeField("severity", func(field types.Field) types.Field {
				field.Type = "enum"
				field.Options = []string{"emerg", "alert", "crit", "err", "warning", "notice", "info", "debug"}
				field.Default = "notice"
				return field
			})
		}).
		MustImportAndCustomize(&Version, v3.EmbeddedConfig{}, func(schema *types.Schema) {
			schema.MustCustomizeField("indexPrefix", func(field types.Field) types.Field {
				field.Required = true
				field.Nullable = false
				return field
			})
			schema.MustCustomizeField("dateFormat", func(field types.Field) types.Field {
				field.Required = true
				field.Nullable = false
				field.Type = "enum"
				field.Options = []string{"YYYY-MM-DD", "YYYY-MM", "YYYY"}
				field.Default = "YYYY-MM-DD"
				return field
			})
		}).
		MustImport(&Version, v3.Logging{}).
		MustImport(&Version, v3.ProjectLogging{})
}
