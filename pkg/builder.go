package recon

import (
	"fmt"
	"strings"

	"github.com/yndc/recon/pkg/schema"
	"github.com/yndc/recon/pkg/utils"
)

type Builder struct {
	loaders           []Loader
	config            *Config
	onValidationError func(key string, value interface{}, err error)
	onLoaded          func(key string, value interface{})
}

// create a new config builder
func New(sample interface{}) *Builder {
	return &Builder{
		loaders: make([]Loader, 0),
		config: &Config{
			value:  sample,
			values: make(map[string]interface{}),
		},
	}
}

// add a file loader to the config builder
func (b *Builder) FromFile(path string) *Builder {
	b.loaders = append(b.loaders, NewFileLoader(path, b.config.LoadMap))
	return b
}

// build into a config container
func (b *Builder) Build() (*Config, error) {
	// build the schema
	schema, err := schema.NewSchema(b.config.value)
	if err != nil {
		return nil, err
	}
	b.config.schema = *schema

	// add hooks to the loaders
	if b.onValidationError == nil {
		b.onValidationError = func(key string, value interface{}, err error) {}
	}
	if b.onLoaded == nil {
		b.onLoaded = func(key string, value interface{}) {}
	}

	// load all values from all loaders
	for _, loader := range b.loaders {
		loader.Load()
	}

	// ensure that the required fields are all filled
	requiredFields := utils.NewSetWithValues(b.config.schema.GetRequiredFieldString()...)
	for key := range b.config.values {
		requiredFields.Remove(key)
	}

	if requiredFields.Count() > 0 {
		return nil, fmt.Errorf("failed to build config, not all required fields are set: %s", strings.Join(requiredFields.Values(), ", "))
	}

	return b.config, nil
}