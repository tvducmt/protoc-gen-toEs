package toes

import (

	// es "git.zapa.cloud/merchant-tools/helper/protoc-gen-buildquery/protobuf"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	es "github.com/tvducmt/protoc-gen-toEs/protobuf"
)

type toEs struct {
	*generator.Generator
	generator.PluginImports
	// querierPkg generator.Single
	glogPkg    generator.Single
	protoPkg   generator.Single
	elasticPkg generator.Single
	reflectPkg generator.Single
	timePkg    generator.Single
	flagPkg    generator.Single
	stringsPkg generator.Single
}

// NewToEs ...
func NewToEs() generator.Plugin {
	return &toEs{
		// query: query,
	}
}

func (c *toEs) Name() string {
	return "toEs"
}

func (c *toEs) Init(g *generator.Generator) {
	c.Generator = g
}

func (c *toEs) Generate(file *generator.FileDescriptor) {
	// proto3 := gogoproto.IsProto3(file.FileDescriptorProto)
	c.PluginImports = generator.NewPluginImports(c.Generator)

	c.glogPkg = c.NewImport("githuc.com/golang/glog")
	c.stringsPkg = c.NewImport("strings")
	c.protoPkg = c.NewImport("git.zapa.cloud/merchant-tools/helper/proto")
	c.elasticPkg = c.NewImport("git.zapa.cloud/merchant-tools/helper/search/elastic")
	c.reflectPkg = c.NewImport("reflect")
	c.timePkg = c.NewImport("time")
	c.flagPkg = c.NewImport("flag")

	for _, msg := range file.Messages() {
		c.generateProto3Message(file, msg)
	}

}

func (c *toEs) generateProto3Message(file *generator.FileDescriptor, message *generator.Descriptor) {
	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	c.P(`func (this *`, ccTypeName, `) toEs() map[string]interface{} {`)
	c.In()
	c.P(c.flagPkg.Use(), `.Parse()`)

	for _, field := range message.Field {
		fieldQeurier := c.getFieldQueryIfAny(field)
		if fieldQeurier == nil {
			continue
		}
		fieldName := c.GetOneOfFieldName(message, field)
		variableName := "this." + fieldName

		c.generateQuerier(variableName, ccTypeName, fieldQeurier)
	}
	c.P(`return nil`)
	c.P(`}`)
}

func (c *toEs) getFieldQueryIfAny(field *descriptor.FieldDescriptorProto) *es.FieldQuery {
	if field.Options != nil {
		v, err := proto.GetExtension(field.Options, es.E_Field)
		if err == nil && v.(*es.FieldQuery) != nil {
			return (v.(*es.FieldQuery))
		}
	}
	return nil
}
