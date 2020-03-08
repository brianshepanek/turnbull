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
	f.Line()

	// Slice Struct
	entitySliceStruct, err := entityToSliceStruct(entity)
	if err != nil {
		return nil, err
	}
	f.Add(entitySliceStruct)
	f.Line()

	// Interface
	entityInterface, err := entityToInterface(entity)
	if err != nil {
		return nil, err
	}
	f.Add(entityInterface)

	// Slice Interface
	entitySliceInterface, err := entityToSliceInterface(entity)
	if err != nil {
		return nil, err
	}
	f.Add(entitySliceInterface)

	// Interface Constructor Function
	interfaceConstructorFunction, err := interfaceConstructorFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(interfaceConstructorFunction)
	f.Line()

	// Slice Interface Constructor Function
	sliceInterfaceConstructorFunction, err := sliceInterfaceConstructorFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(sliceInterfaceConstructorFunction)
	f.Line()

	// Len
	lenFunction, err := lenFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(lenFunction)

	// Append
	appendFunction, err := appendFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(appendFunction)

	// Elements
	elementsFunction, err := elementsFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(elementsFunction)

	// Getters
	for _, field := range entity.Fields {
		entityGetter, err := fieldToGetter(entity, field)
		if err != nil {
			return nil, err
		}
		f.Add(entityGetter)
		f.Line()

		if field.Primary {
			entityGetter, err := fieldToPrimaryGetter(entity, field)
			if err != nil {
				return nil, err
			}
			f.Add(entityGetter)
			f.Line()
		}
	}
	
	// Setters
	for _, field := range entity.Fields {
		entitySetter, err := fieldToSetter(entity, field)
		if err != nil {
			return nil, err
		}
		f.Add(entitySetter)
		f.Line()
	}

	// Set All Setter
	setAllSetter, err := setAllSetter(entity)
	if err != nil {
		return nil, err
	}
	f.Add(setAllSetter)
	f.Line()

	// JSON
	if entity.JSON {

		entityMarshalJSON, err := entityMarshalJSON(entity)
		if err != nil {
			return nil, err
		}
		f.Add(entityMarshalJSON)
		f.Line()

		entityUnmarshalJSON, err := entityUnmarshalJSON(entity)
		if err != nil {
			return nil, err
		}
		f.Add(entityUnmarshalJSON)
		f.Line()
	}

	return f, nil
}

func entityToSliceStruct(entity entity) (jen.Code, error){

	// Vars
	var resp jen.Statement

	// Struct
	resp.Type()

	// Type
	if entity.Name != "" {
		resp.Id(sliceStructId(entity))
	}

	// Index
	resp.Index()

	// Type
	resp.Qual("", interfaceId(entity))

	return &resp, nil
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
	if field.Type != "self" {
		if field.Slice {
			resp.Index()
		}
	}
	

	// Qual
	if field.Type != "" {
		if field.Type == "self" {
			if field.Slice {
				resp.Qual(scaffoldEntitiesFilePath(), sliceInterfaceId(entity))
			} else {
				resp.Qual(scaffoldEntitiesFilePath(), interfaceId(entity))
			}
			
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

	// Set All
	code, err := setAllInterfaceSetter(entity)
	if err != nil {
		return nil, err
	}

	// Append
	methods = append(methods, code)

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

func entityToSliceInterface(entity entity) (jen.Code, error){

	// Vars
	var resp jen.Statement

	// Interface
	resp.Type()

	// Methods
	var methods []jen.Code

	// Len
	methods = append(methods, jen.Id("Len").Params().Parens(jen.List(
		jen.Int(),
	)))

	// Append
	methods = append(methods, jen.Id("Append").Params(jen.Id("req").Qual("", interfaceId(entity))))

	// Elements
	methods = append(methods, jen.Id("Elements").Params().Parens(jen.List(
		jen.Index().
		Qual("", interfaceId(entity)),
	)))

	// Type
	if entity.Name != "" {
		resp.Id(sliceInterfaceId(entity))
	}

	// Methods
	resp.Interface(methods...)

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

func setAllInterfaceSetter(entity entity) (jen.Code, error){

	// Vars
	var resp jen.Statement

	// ID
	resp.Id("SetAll").
	Params(
		jen.Id("req").
		Qual("", interfaceId(entity)),
	)
	
	return &resp, nil
}

func interfaceConstructorFunction(entity entity) (jen.Code, error){

	// Vars
	var resp jen.Statement

	// Func
	resp.Func().
	Id(interfaceConstructorId(entity)).
	Params().
	List(
		jen.Id(interfaceId(entity)),
	).
	Block(
		jen.Return(
			jen.Op("&").
			Id(structId(entity)).
			Values(),
		),
	)

	return &resp, nil

}

func sliceInterfaceConstructorFunction(entity entity) (jen.Code, error){

	// Vars
	var resp jen.Statement

	// Func
	resp.Func().
	Id(sliceInterfaceConstructorId(entity)).
	Params().
	List(
		jen.Id(sliceInterfaceId(entity)),
	).
	Block(
		jen.Return(
			jen.Op("&").
			Id(sliceStructId(entity)).
			Values(),
		),
	)

	return &resp, nil

}

func lenFunction(entity entity) (jen.Code, error){

	// Vars
	var resp jen.Statement

	// Func
	resp.Func()


	resp.Params(
		jen.Id("m").
		Op("*").
		Qual("", sliceStructId(entity)),
	).
	Id("Len").
	Params().
	List(
		jen.Int(),
	).
	Block(
		jen.If(
			jen.Id("m").
			Op("!=").
			Nil(),
		).
		Block(
			jen.Return(
				jen.Len(
					jen.Op("*").
					Id("m"),
				),
			),
		),
		jen.Return(
			jen.Lit(0),
		),
	)

	return &resp, nil

}

func appendFunction(entity entity) (jen.Code, error){

	// Vars
	var resp jen.Statement

	// Func
	resp.Func()


	resp.Params(
		jen.Id("m").
		Op("*").
		Qual("", sliceStructId(entity)),
	).
	Id("Append").
	Params(
		jen.Id("req").
		Qual("", interfaceId(entity)),
	).
	Block(
		jen.If(
			jen.Id("m").
			Op("!=").
			Nil(),
		).
		Block(
			jen.Op("*").
			Id("m").
			Op("=").
			Append(
				jen.Op("*").
				Id("m"),
				jen.Id("req"),
			),
		),
	)

	return &resp, nil

}

func elementsFunction(entity entity) (jen.Code, error){

	// Vars
	var resp jen.Statement

	// Func
	resp.Func()


	resp.Params(
		jen.Id("m").
		Op("*").
		Qual("", sliceStructId(entity)),
	).
	Id("Elements").
	Params().
	List(
		jen.Index().
		Id(interfaceId(entity)),
	).
	Block(
		jen.Return(
			jen.Op("*").
			Id("m"),
		),
	)

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

func setAllSetter(entity entity) (jen.Code, error){

	// Vars
	var resp jen.Statement
	var block []jen.Code

	// Func
	resp.Func()

	// ID
	resp.Params(
		jen.Id("m").Op("*").Qual("", structId(entity)),
	).
	Id("SetAll").
	Params(
		jen.Id("req").
		Qual("", interfaceId(entity)),
	)
	
	// Block
	for _, field := range entity.Fields {
	
		block = append(block,
			jen.Id("m").
			Dot(setterId(field)).
			Call(
				jen.Id("req").
				Dot(getterId(field)).
				Call(),
			),
		)
		
	}
	
	// Block
	resp.Block(
		block...,
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