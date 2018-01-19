package schema

import (
	"github.com/rancher/norman/types"
	"github.com/rancher/types/apis/tools.cattle.io/v4"
	"github.com/rancher/types/factory"
	"github.com/rancher/types/mapper"
)

var (
	Version = types.APIVersion{
		Version: "v4",
		Group:   "tools.cattle.io",
		Path:    "/v4",
	}

	Schemas = factory.Schemas(&Version).
		Init(loggingTypes)
)

func loggingTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v4.ProjectLogging{}, &mapper.NamespaceIDMapper{}).
		MustImportAndCustomize(&Version, v4.ElasticsearchConfig{}, func(schema *types.Schema) {
			schema.MustCustomizeField("logstashDateformat", func(field types.Field) types.Field {
				field.Type = "enum"
				field.Options = []string{"YYYY.MM.DD", "YYYY.MM", "YYYY"}
				field.Required = true
				field.Default = "YYYY.MM.DD"
				field.Nullable = false
				return field
			})
		}).
		MustImportAndCustomize(&Version, v4.SplunkConfig{}, func(schema *types.Schema) {
			schema.MustCustomizeField("protocol", func(field types.Field) types.Field {
				field.Type = "enum"
				field.Options = []string{"http", "https"}
				field.Required = true
				field.Default = "http"
				field.Nullable = false
				return field
			})
			schema.MustCustomizeField("timeFormat", func(field types.Field) types.Field {
				field.Type = "enum"
				field.Options = []string{"unixtime", "localtime", "none"}
				field.Required = true
				field.Default = "unixtime"
				field.Nullable = false
				return field
			})
		}).
		MustImportAndCustomize(&Version, v4.ElasticsearchConfig{}, func(schema *types.Schema) {
			schema.MustCustomizeField("logstashDateformat", func(field types.Field) types.Field {
				field.Type = "enum"
				field.Options = []string{"YYYY.MM.DD", "YYYY.MM", "YYYY"}
				field.Required = true
				field.Default = "YYYY.MM.DD"
				field.Nullable = false
				return field
			})
		}).
		MustImportAndCustomize(&Version, v4.SplunkConfig{}, func(schema *types.Schema) {
			schema.MustCustomizeField("protocol", func(field types.Field) types.Field {
				field.Type = "enum"
				field.Options = []string{"http", "https"}
				field.Required = true
				field.Default = "http"
				field.Nullable = false
				return field
			})
			schema.MustCustomizeField("timeFormat", func(field types.Field) types.Field {
				field.Type = "enum"
				field.Options = []string{"unixtime", "localtime", "none"}
				field.Required = true
				field.Default = "unixtime"
				field.Nullable = false
				return field
			})
		}).
		MustImportAndCustomize(&Version, v4.KafkaConfig{}, func(schema *types.Schema) {
			schema.MustCustomizeField("brokerType", func(field types.Field) types.Field {
				field.Type = "enum"
				field.Options = []string{"broker", "zookeeper"}
				field.Required = true
				field.Default = "zookeeper"
				field.Nullable = false
				return field
			})
			schema.MustCustomizeField("outputDataType", func(field types.Field) types.Field {
				field.Type = "enum"
				field.Options = []string{"json", "ltsv", "msgpack"}
				field.Required = true
				field.Default = "json"
				field.Nullable = false
				return field
			})
		}).
		MustImport(&Version, v4.Logging{}).
		MustImport(&Version, v4.ProjectLogging{})
}
