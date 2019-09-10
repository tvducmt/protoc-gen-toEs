package toes

import (

	// querier "git.zapa.cloud/merchant-tools/helper/protoc-gen-buildquery/protobuf"

	pb_svc "go-plugin-demo/client-test/protobuf/copy/report-service"
	"sync"

	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
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

// NewCopy ...
func NewCopy() generator.Plugin {
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

	c.P(`func indirectType(reflectType reflect.Type) reflect.Type {`)
	c.P(`for reflectType.Kind() == reflect.Ptr || reflectType.Kind() == reflect.Slice {`)
	c.P(`reflectType = reflectType.Elem()`)
	c.P(`}`)
	c.P(`return reflectType`)
	c.P(`}`)

	// for _, msg := range file.Messages() {

	// 	c.generateProto3Message(file, msg)
	// 	// }
	// }
	// for _, msg := range file.Messages() {

	// 	c.generateSetFieldMethod(file, msg)
	// 	// }
	// }
}

func (c *copy) generateSetFieldMethod(file *generator.FileDescriptor, message *generator.Descriptor) {
	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	for _, field := range message.Field {
		fieldName := c.GetOneOfFieldName(message, field)
		c.P(`func (this *`, ccTypeName, `) Set`, fieldName, ` (resp interface{})bool {`)
		c.P(`if this != nil {`)
		c.P(`str := fmt.Sprintf("%v", resp)`)
		c.P(`this.`, fieldName, ` = str`)
		c.P(`return true`)

		c.P(`}`)
		c.P(`return false`)
		c.P(`}`)
	}

}
func (c *copy) generateField(reqCoreType interface{}, ccTypeName, fieldName string) {
	c.P(`if _, ok := reqCoreServiceVal.Type().FieldByName("`, fieldName, `"); ok {`)
	c.P(`if k, ok := resp.(interface {`)
	c.P(`Set`, fieldName, `(v interface{}) bool`)
	c.P(`}); ok {`)
	c.P(`k.Set`, fieldName, `(this.`, fieldName, `)`)
	c.P(`} else {`)
	c.P(`}`)
	//c.P(reqCoreType, `.`, fieldName, ` = this.`, fieldName)
	c.P(`}`)
}

func (c *copy) generateProto3Message(file *generator.FileDescriptor, message *generator.Descriptor) {
	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	c.P(`func (this *`, ccTypeName, `) Copy(resp interface{}) {`)
	c.In()
	c.P(c.flagPkg.Use(), `.Parse()`)
	once3 := &sync.Once{}

	reqCoreServiceVal := func() {
		c.P(`reqCoreServiceVal := reflect.Indirect(reflect.ValueOf(resp))`)
	}
	reqCoreType := &pb_svc.ListCITransactionsRequest{} // reqCoreServiceVal.Type()
	for _, field := range message.Field {
		once3.Do(reqCoreServiceVal)
		fieldName := c.GetOneOfFieldName(message, field)
		c.generateField(reqCoreType, ccTypeName, fieldName)
	}

	// c.P(`reqVal := reflect.Indirect(reflect.ValueOf(`, ccTypeName, `{}))`)
	// c.P(`reqType := indirectType(reqVal.Type())`)
	// c.P(`reqCoreServiceVal := reflect.Indirect(reflect.ValueOf(resp))`)
	// c.P(`fields:= `, fields)
	// c.P(`for i,v := range  fields {`)
	// // c.P(`v := reqType.Field(i)`)
	// // c.P(`fmt.Println(":reqCoreServiceVal.Type()", reqCoreServiceVal.Type().Name())`)
	// c.P(`if !strings.HasPrefix(v.Name, "XXX_") {`)
	// c.P(`if f, ok := reqCoreServiceVal.Type().FieldByName(v.Name); ok {`)
	// c.P(`nameField := v.Name`)
	// // reqCoreService.ZpTransID = l.GetZpTransID()
	// // c.P(`reqCoreServiceVal.Type().Name().`, nameField, ` = this.`, nameField)
	// c.P(`fmt.Println("into here  v.Index", v.Name)`)
	// //c.P(`fmt.Println("into here  f.Index", f.Name)`)
	// // fmt.Println("into here  v.Index", v.Index)
	// // fields = append(fields, &searchField{indexFrom: v.Index, indexTo: f.Index})
	// c.P(`}`)

	// c.P(`}`)
	// c.P(`}`)

	// c.Out()
	c.P(`}`)
}
