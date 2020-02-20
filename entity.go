package turnbull

import(
	"strings"
	"github.com/dave/jennifer/jen"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

func buildEntityFile(entity entity) (*jen.File, error){

	// Entities
	pluralize := pluralize.NewClient()
	f := jen.NewFile(pluralize.Plural(entityName))

	// Struct
	entityStruct, err := entityToStruct(entity)
	if err != nil {
		return nil, err
	}
	f.Add(entityStruct)

	// Interface
	entityInterface, err := entityToInterface(entity)
	if err != nil {
		return nil, err
	}
	f.Add(entityInterface)

	// Getters
	for _, field := range entity.Fields {
		entityGetter, err := fieldToGetter(entity, field)
		if err != nil {
			return nil, err
		}
		f.Add(entityGetter)
	}
	
	// Setters
	for _, field := range entity.Fields {
		entitySetter, err := fieldToSetter(entity, field)
		if err != nil {
			return nil, err
		}
		f.Add(entitySetter)
	}

	// JSON
	if entity.JSON {

	}

	return f, nil
}

func structId(entity entity)(string){
	return strcase.ToLowerCamel(strings.Join([]string{entity.Name, entityName}, stringSeparator))
}

func interfaceId(entity entity)(string){
	return strcase.ToCamel(strings.Join([]string{entity.Name, entityName}, stringSeparator))
}

func getterId(field field)(string){
	return strcase.ToCamel(field.Name)
}

func setterId(field field)(string){
	return strcase.ToCamel(strings.Join([]string{setterVerb, field.Name}, stringSeparator))
}

func entityToStruct(entity entity) (jen.Code, error){

	// Vars
	var resp jen.Statement

	// Struct
	resp.Type()

	// Fields
	var fields []jen.Code
	for _, field := range entity.Fields {

		// Field to Struct Field
		code, err := fieldToStructField(field)
		if err != nil {
			return nil, err
		}

		// Append
		fields = append(fields, code)
	}

	// Type
	if entity.Name != "" {
		resp.Id(structId(entity))
	}

	// Fields
	if len(fields) > 0 {
		resp.Struct(fields...)
	}

	return &resp, nil

}

func fieldToStructField(field field) (jen.Code, error){

	// Vars
	var resp jen.Statement

	// ID
	if field.Name != "" {
		resp.Id(field.Name)
	}

	// Slice
	if field.Slice {
		resp.Index()
	}

	// Qual
	if field.Type != "" {
		resp.Qual(field.Package, field.Type)
	} 

	return &resp, nil
}

func entityToInterface(entity entity) (jen.Code, error){

	// Vars
	var resp jen.Statement

	// Interface
	resp.Type()

	// Methods
	var methods []jen.Code

	// Getters
	for _, field := range entity.Fields {

		// Field to Interface Getter
		code, err := fieldToInterfaceGetter(field)
		if err != nil {
			return nil, err
		}

		// Append
		methods = append(methods, code)
		
	}

	// Setters
	for _, field := range entity.Fields {

		// Field to Interface Setter
		code, err := fieldToInterfaceSetter(field)
		if err != nil {
			return nil, err
		}

		// Append
		methods = append(methods, code)
	}

	// Marshal JSON
	if entity.JSON{
		methods = append(methods, jen.Id("MarshalJSON").Params().Parens(jen.List(
			jen.Index().Byte(),
			jen.Error(),
		)))
	}
	

	// Type
	if entity.Name != "" {
		resp.Id(interfaceId(entity))
	}

	// Methods
	if len(methods) > 0 {
		resp.Interface(methods...)
	}

	return &resp, nil

}

func fieldToInterfaceGetter(field field) (jen.Code, error){

	// Vars
	var resp jen.Statement

	// ID
	if field.Name != "" {
		resp.Id(getterId(field))
		resp.Params()
	}

	// Slice
	if field.Slice {
		resp.Index()
	}

	// Qual
	if field.Type != "" {
		resp.Qual(field.Package, field.Type)
	} 

	return &resp, nil
}

func fieldToInterfaceSetter(field field) (jen.Code, error){

	// Vars
	var resp, args jen.Statement

	// ID
	if field.Name != "" {
		resp.Id(setterId(field))
		args.Id(field.Name)
	}

	// Slice
	if field.Slice {
		args.Index()
	}

	// Qual
	if field.Type != "" {
		args.Qual(field.Package, field.Type)
	} 

	// Params
	resp.Params(&args)

	return &resp, nil
}

func fieldToGetter(entity entity, field field) (jen.Code, error){

	// Vars
	var resp jen.Statement

	// Func
	resp.Func()

	// ID
	if field.Name != "" {
		resp.Params(
			jen.Id(structId(entity)).Op("*").Qual("", structId(entity)),
		).Id(getterId(field)).Params()
	}

	// Slice
	if field.Slice {
		resp.Index()
	}

	// Qual
	if field.Type != "" {
		resp.Qual(field.Package, field.Type)
	} 

	// Return
	resp.Block(
		jen.Return(
			jen.Id(structId(entity) + "." + field.Name),
		),
	)

	

	return &resp, nil
}

func fieldToSetter(entity entity, field field) (jen.Code, error){

	// Vars
	var resp, args jen.Statement

	// Func
	resp.Func()

	// ID
	if field.Name != "" {
		resp.Params(
			jen.Id(structId(entity)).Op("*").Qual("", structId(entity)),
		).Id(setterId(field))
		args.Id(field.Name)
	}

	// Slice
	if field.Slice {
		args.Index()
	}

	// Qual
	if field.Type != "" {
		args.Qual(field.Package, field.Type)
	} 

	// Params
	resp.Params(&args)

	// Block
	resp.Block(
		jen.Id(structId(entity) + "." + field.Name).Op("=").Id(field.Name),
	)
	

	return &resp, nil
}