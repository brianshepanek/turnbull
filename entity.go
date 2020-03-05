package turnbull

import(
	"github.com/dave/jennifer/jen"
)

func buildScaffoldEntityFile(entity entity) (*jen.File, error){

	// Entities
	f := jen.NewFile(packageName())

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

		if field.Primary {
			entityGetter, err := fieldToPrimaryGetter(entity, field)
			if err != nil {
				return nil, err
			}
			f.Add(entityGetter)
		}
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

		entityMarshalJSON, err := entityMarshalJSON(entity)
		if err != nil {
			return nil, err
		}
		f.Add(entityMarshalJSON)

		entityUnmarshalJSON, err := entityUnmarshalJSON(entity)
		if err != nil {
			return nil, err
		}
		f.Add(entityUnmarshalJSON)
	}

	return f, nil
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
		code, err := fieldToStructField(entity, field)
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

func fieldToStructField(entity entity, field field) (jen.Code, error){

	// Vars
	var resp jen.Statement

	// ID
	if field.Name != "" {
		resp.Id(structFieldId(field))
	}

	// Op
	if field.Op != "" {
		resp.Op(field.Op)
	}

	// Slice
	if field.Slice {
		resp.Index()
	}

	// Qual
	if field.Type != "" {
		if field.Type == "self" {
			resp.Qual(scaffoldEntitiesFilePath(), interfaceId(entity))
		} else {
			resp.Qual(field.Package, field.Type)
		}
	} 

	return &resp, nil
}

func fieldToJSONStructField(entity entity, field field) (jen.Code, error){

	// Vars
	var resp jen.Statement

	// ID
	if field.Name != "" {
		resp.Id(getterId(field))
	}

	// Op
	if field.Op != "" {
		resp.Op(field.Op)
	}

	// Slice
	if field.Slice {
		resp.Index()
	}

	// Qual
	if field.Type != "" {
		if field.Type == "self" {
			resp.Qual(scaffoldEntitiesFilePath(), interfaceId(entity))
		} else {
			resp.Qual(field.Package, field.Type)
		}
	} 

	// Tags
	if field.Name != "" {
		resp.Tag(map[string]string{"json": tagId(field)})
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
		
		// Primary
		if field.Primary {

			code, err := fieldToPrimaryInterfaceGetter(field)
			if err != nil {
				return nil, err
			}

			// Append
			methods = append(methods, code)

		}

		
		
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
		methods = append(methods, jen.Id("UnmarshalJSON").Params(jen.Id("data").Index().Byte()).Parens(jen.List(
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

func fieldToPrimaryInterfaceGetter(field field) (jen.Code, error){

	// Vars
	var resp jen.Statement

	// ID
	if field.Name != "" {
		resp.Id("Primary")
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
			jen.Id("m").Op("*").Qual("", structId(entity)),
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
			jen.Id("m" + "." + structFieldId(field)),
		),
	)

	

	return &resp, nil
}


func fieldToPrimaryGetter(entity entity, field field) (jen.Code, error){

	// Vars
	var resp jen.Statement

	// Func
	resp.Func()

	// ID
	if field.Name != "" {
		resp.Params(
			jen.Id("m").Op("*").Qual("", structId(entity)),
		).Id("Primary").Params()
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
			jen.Id("m" + "." + structFieldId(field)),
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
			jen.Id("m").Op("*").Qual("", structId(entity)),
		).Id(setterId(field))
		args.Id(structFieldId(field))
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
		jen.Id("m" + "." + structFieldId(field)).Op("=").Id(structFieldId(field)),
	)
	

	return &resp, nil
}

func entityMarshalJSON(entity entity) (jen.Code, error){


	// Vars
	var resp jen.Statement

	// Func
	resp.Func()

	// ID
	resp.Params(
		jen.Id("m").Op("*").Qual("", structId(entity)),
	).
	Id("MarshalJSON").
	Params().
	Parens(
		jen.List(
			jen.Index().Byte(),
			jen.Error(),
		),
	)	

	var jsonStruct []jen.Code
	jsonStructDict := make(jen.Dict)
	for _, field := range entity.Fields {
		code, err := fieldToJSONStructField(entity, field)
		if err != nil {
			return nil, err
		}
		jsonStruct = append(jsonStruct, code)
		jsonStructDict[jen.Id(getterId(field))] = jen.Id("m").Dot(getterId(field)).Call()
		
	}

	// Block
	resp.Block(

		jen.Type().Id("jsonStructPrivate").Struct(jsonStruct...),

		jen.Id("jsonStruct").
		Op(":=").
		Qual("", "jsonStructPrivate").
		Values(jsonStructDict),

		jen.Return().
		Qual("encoding/json", "Marshal").
		Call(
			jen.Op("&").
			Id("jsonStruct"),
		),
	)

	return &resp, nil
}

func entityUnmarshalJSON(entity entity) (jen.Code, error){


	// Vars
	var resp jen.Statement

	// Func
	resp.Func()

	// ID
	resp.Params(
		jen.Id("m").Op("*").Qual("", structId(entity)),
	).
	Id("UnmarshalJSON").
	Params(
		jen.Id("data").Index().Byte(),
	).
	Parens(
		jen.List(
			jen.Error(),
		),
	)	

	var jsonStruct []jen.Code
	var jsonSetterFunctions []jen.Code
	// jsonStructDict := make(jen.Dict)
	for _, field := range entity.Fields {
		code, err := fieldToJSONStructField(entity, field)
		if err != nil {
			return nil, err
		}
		jsonStruct = append(jsonStruct, code)

		jsonSetterFunctions = append(jsonSetterFunctions, jen.Id("m").
			Dot(setterId(field)).
			Call(
				jen.Id("jsonStruct").
				Dot(getterId(field)),
			),
		)
		// jsonStructDict[jen.Id(getterId(field))] = jen.Id("m").Dot(getterId(field)).Call()
		
	}

	// Block
	var block []jen.Code


	block = append(block, 
		jen.Type().
		Id("jsonStructPrivate").
		Struct(jsonStruct...),
	)
	block = append(block, 
		jen.Id("jsonStruct").
		Op(":=").
		Qual("", "jsonStructPrivate").
		Values(),
	)
	block = append(block, 
		jen.Err().
		Op(":=").
		Qual("encoding/json", "Unmarshal").
		Call(
			jen.Id("data"),
			jen.Op("&").
			Id("jsonStruct"),
		),
	)
	block = append(block, 
		jen.If(
			jen.Err().
			Op("!=").
			Nil(),
		).
		Block(
			jen.Return(
				jen.List(
					jen.Err(),
				),
			),
		),
	)
	for _, jsonSetterFunction := range jsonSetterFunctions {
		block = append(block, jsonSetterFunction)
	}
	block = append(block, jen.Return(
		jen.List(
			jen.Nil(),
		),
	),)

	resp.Block(
		block...,
	)

	return &resp, nil
}