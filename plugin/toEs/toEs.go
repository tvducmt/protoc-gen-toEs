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
	pbTimePkg  generator.Single
	ptypesPkg  generator.Single
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
	c.PluginImports = generator.NewPluginImports(c.Generator)
	c.glogPkg = c.NewImport("github.com/golang/glog")
	c.protoPkg = c.NewImport("git.zapa.cloud/merchant-tools/helper/proto")
	c.elasticPkg = c.NewImport("git.zapa.cloud/merchant-tools/helper/search/elastic")
	c.reflectPkg = c.NewImport("reflect")
	c.timePkg = c.NewImport("time")
	c.flagPkg = c.NewImport("flag")
	c.pbTimePkg = c.NewImport("github.com/golang/protobuf/ptypes/timestamp")
	c.ptypesPkg = c.NewImport("github.com/golang/protobuf/ptypes")

	c.P(`func checkNull(field interface{}) bool {`)
	c.P(`zero := `, c.reflectPkg.Use(), `.Zero(`, c.reflectPkg.Use(), `.TypeOf(field)).Interface()	`)
	c.P(`if `, c.reflectPkg.Use(), `.DeepEqual(field, zero) {`)
	c.P(`return true`)
	c.P(`}`)
	c.P(`return false`)
	c.P(`}`)

	c.P(`func checkTimestampType(field interface{}) (*`, c.pbTimePkg.Use(), `.Timestamp, bool) {`)
	c.P(`if ts, ok := field.(*`, c.pbTimePkg.Use(), `.Timestamp); ok {`)
	c.P(`return ts, true`)
	c.P(`}`)
	c.P(`return nil, false`)
	c.P(`}`)

	c.P(`func checkDateType(field interface{}) (*`, c.protoPkg.Use(), `.Date, bool) {`)
	c.P(`if date, ok := field.(*`, c.protoPkg.Use(), `.Date); ok {`)
	c.P(`return date, true`)
	c.P(`}`)
	c.P(`return nil, false`)
	c.P(`}`)

	c.P(`func makeKeyMap(m *map[string]interface{}, key string) *map[string]interface{} {`)
	c.P(`if t, ok := (*m)[key]; ok {`)
	c.P(`if t, ok := t.(*map[string]interface{}); ok {`)
	c.P(`return t`)
	c.P(`}`)
	c.P(`}`)
	c.P(`t := &map[string]interface{}{}`)
	c.P(`(*m)[key] = t`)
	c.P(`return t`)
	c.P(`}`)

	for _, msg := range file.Messages() {
		if msg.DescriptorProto.GetOptions().GetMapEntry() {
			continue
		}
		c.generateProto3Message(file, msg)
	}

}

func (c *toEs) generateProto3Message(file *generator.FileDescriptor, message *generator.Descriptor) {
	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	c.P(`func (this *`, ccTypeName, `) GetEsMap(esMap *map[string]interface{}) {`)
	c.In()
	c.P(c.flagPkg.Use(), `.Parse()`)
	// c.P(`esMap := map[string]interface{}{}`)

	for _, field := range message.Field {
		fieldEs := c.getFieldQueryIfAny(field)
		if fieldEs == nil {
			continue
		}
		fieldName := c.GetOneOfFieldName(message, field)
		variableName := "this." + fieldName

		//} else if field.IsEnum() {
		//	c.generateEnumEs(field, variableName, ccTypeName, fieldName, fieldValidator)

		if field.IsMessage() {
			c.generatePtrAndStructEs(variableName, ccTypeName, fieldName, fieldEs)
		} else {
			c.generateFieldEs(variableName, ccTypeName, fieldEs)
		}

	}
	c.P(`}`)
}

func (c *toEs) generatePtrAndStructEs(variableName string, ccTypeName string, fieldName string, fv *es.FieldEs) {
	tag := fv.GetEs()
	if tag != "" {
		c.P(`if !checkNull( ` + variableName + `){`)
		c.P(`if ts, ok := checkTimestampType(`, variableName, `); ok {`)
		c.P(`if ts != nil {`)
		c.P(`tm, err := `, c.ptypesPkg.Use(), `.Timestamp(ts)`)
		c.P(`if err != nil {`)
		c.P(c.glogPkg.Use(), `.Errorln(err)`)
		c.P(`} else {`)
		c.P(`(*esMap)["`, tag, `"] = tm.UnixNano() / int64(`, c.timePkg.Use(), `.Millisecond)`)

		c.P(`}`)
		c.P(`}`)
		c.P(`} else if date, ok := checkDateType(`, variableName, `); ok {`)
		c.P(`if date != nil {`)
		c.P(`tm := `, c.protoPkg.Use(), `.DateToTimeSearch(date)`)
		c.P(`(*esMap)["`, tag, `"] = tm.UnixNano() / int64(`, c.timePkg.Use(), `.Millisecond)`)

		c.P(`}`)
		c.P(`} else {`)
		c.P(`this.Get`, fieldName, `().GetEsMap(makeKeyMap(esMap, "`, tag, `"))`)
		c.P(`}`)

		c.P(`}`)
	}
}

func (c *toEs) generateFieldEs(variableName string, ccTypeName string, fv *es.FieldEs) {
	tag := fv.GetEs()
	if tag != "" {
		c.P(`if !checkNull( ` + variableName + `){`)
		c.P(`(*esMap)["`, tag, `"] = `, variableName)
		c.P(`}`)

	}
}

func (c *toEs) getFieldQueryIfAny(field *descriptor.FieldDescriptorProto) *es.FieldEs {
	if field.Options != nil {
		v, err := proto.GetExtension(field.Options, es.E_Field)
		if err == nil && v.(*es.FieldEs) != nil {
			return (v.(*es.FieldEs))
		}
	}
	return nil
}
