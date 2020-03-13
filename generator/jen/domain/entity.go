package domain

import(
	"github.com/dave/jennifer/jen"
	"github.com/brianshepanek/turnbull/domain/model"
	"github.com/brianshepanek/turnbull/formatter"
	"github.com/brianshepanek/turnbull/generator/jen/helper"
)

type entityGenerator struct{
	formatter formatter.Formatter
	helperGenerator helper.Generator
}

type EntityGenerator interface{

	// File
	ScaffoldFile(entity model.Entity) (*jen.File, error)

	// Scaffold Entity Struct
	scaffoldEntityStruct(entity model.Entity) (jen.Statement, error)
	scaffoldEntitySliceStruct(entity model.Entity) (jen.Statement, error)

	scaffoldEntityStructField(field model.Field, entity model.Entity) (jen.Code, error)

	// Scaffold Entity Interface
	scaffoldEntityInterface(entity model.Entity) (jen.Statement, error)

	// Scaffold Entity Interface Methods
	scaffoldEntityInterfaceGetter(field model.Field, entity model.Entity) (jen.Code, error)
	scaffoldEntityInterfaceSetter(field model.Field, entity model.Entity) (jen.Code, error)
	scaffoldEntityInterfaceSetAllSetter() (jen.Code, error)
	scaffoldEntityInterfaceMarshalJSON() (jen.Code, error)
	scaffoldEntityInterfaceUnmarshalJSON() (jen.Code, error)

	// Scaffold Entity Slice Interface
	scaffoldEntitySliceInterface(entity model.Entity) (jen.Statement, error)

	// Scaffold Entity Slice Interface Methods
	scaffoldEntitySliceInterfaceLen() (jen.Code, error)
	scaffoldEntitySliceInterfaceAppend(entity model.Entity) (jen.Code, error)
	scaffoldEntitySliceInterfaceElements(entity model.Entity) (jen.Code, error)

	// Scaffold Entity Interface Functions
	scaffoldEntityInterfaceConstructorFunction(entity model.Entity) (jen.Statement, error)
	scaffoldEntitySliceInterfaceConstructorFunction(entity model.Entity) (jen.Statement, error)
	scaffoldEntitySliceInterfaceLenFunction(entity model.Entity) (jen.Statement, error)
	scaffoldEntitySliceInterfaceAppendFunction(entity model.Entity) (jen.Statement, error)
	scaffoldEntitySliceInterfaceElementsFunction(entity model.Entity) (jen.Statement, error)
	scaffoldEntityInterfaceSetterFunction(field model.Field, entity model.Entity) (jen.Statement, error)
	scaffoldEntityInterfaceGetterFunction(field model.Field, entity model.Entity) (jen.Statement, error)
	scaffoldEntityInterfaceSetAllSetterFunction(entity model.Entity) (jen.Statement, error)
	scaffoldEntityInterfaceMarshalJSONFunction(entity model.Entity) (jen.Statement, error)
	scaffoldEntityInterfaceUnmarshalJSONFunction(entity model.Entity) (jen.Statement, error)
	scaffoldEntityJSONStructField(field model.Field, entity model.Entity) (jen.Code, error)

}

func NewEntityGenerator(formatter formatter.Formatter, helperGenerator helper.Generator) EntityGenerator {
	return &entityGenerator{
		formatter : formatter,
		helperGenerator : helperGenerator,
	}
}

func (entityGenerator *entityGenerator) ScaffoldFile(entity model.Entity) (*jen.File, error){
	
	
	// File
	packageName , err := entityGenerator.formatter.OutputScaffoldDomainEntityPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)

	// Struct
	entityStruct, err := entityGenerator.scaffoldEntityStruct(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&entityStruct)
	f.Line()

	// Slice Struct
	entitySliceStruct, err := entityGenerator.scaffoldEntitySliceStruct(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&entitySliceStruct)
	f.Line()

	// Interface
	entityInterface, err := entityGenerator.scaffoldEntityInterface(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&entityInterface)

	// Slice Interface
	entitySliceInterface, err := entityGenerator.scaffoldEntitySliceInterface(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&entitySliceInterface)

	// Interface Constructor Function
	interfaceConstructorFunction, err := entityGenerator.scaffoldEntityInterfaceConstructorFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&interfaceConstructorFunction)
	f.Line()

	// Slice Interface Constructor Function
	sliceInterfaceConstructorFunction, err :=  entityGenerator.scaffoldEntitySliceInterfaceConstructorFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&sliceInterfaceConstructorFunction)
	f.Line()

	// Len
	lenFunction, err := entityGenerator.scaffoldEntitySliceInterfaceLenFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&lenFunction)

	// Append
	appendFunction, err := entityGenerator.scaffoldEntitySliceInterfaceAppendFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&appendFunction)

	// Elements
	elementsFunction, err := entityGenerator.scaffoldEntitySliceInterfaceElementsFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&elementsFunction)

	// Getters
	for _, field := range entity.Fields {
		entityGetter, err := entityGenerator.scaffoldEntityInterfaceGetterFunction(field, entity)
		if err != nil {
			return nil, err
		}
		f.Add(&entityGetter)
		f.Line()

		// if field.Primary {
		// 	entityGetter, err := fieldToPrimaryGetter(entity, field)
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	f.Add(entityGetter)
		// 	f.Line()
		// }
	}
	
	// Setters
	for _, field := range entity.Fields {
		entitySetter, err := entityGenerator.scaffoldEntityInterfaceSetterFunction(field, entity)
		if err != nil {
			return nil, err
		}
		f.Add(&entitySetter)
		f.Line()
	}

	// Set All Setter
	setAllSetter, err := entityGenerator.scaffoldEntityInterfaceSetAllSetterFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&setAllSetter)
	f.Line()

	// JSON
	if entity.JSON {

		entityMarshalJSON, err := entityGenerator.scaffoldEntityInterfaceMarshalJSONFunction(entity)
		if err != nil {
			return nil, err
		}
		f.Add(&entityMarshalJSON)
		f.Line()

		entityUnmarshalJSON, err := entityGenerator.scaffoldEntityInterfaceUnmarshalJSONFunction(entity)
		if err != nil {
			return nil, err
		}
		f.Add(&entityUnmarshalJSON)
		f.Line()
	}
	
	
	return f, nil
}

func (entityGenerator *entityGenerator) scaffoldEntityStruct(entity model.Entity) (jen.Statement, error){
	
	// Vars
	var resp jen.Statement
	var fields []jen.Code

	// Struct
	resp.Type()

	// Type
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Fields
	for _, field := range entity.Fields {
		code, err := entityGenerator.scaffoldEntityStructField(field, entity)
		if err != nil {
			return nil, err
		}
		fields = append(fields, code)
	}

	// Struct
	resp.Struct(fields...)

	
	return resp, nil
}

func (entityGenerator *entityGenerator) scaffoldEntitySliceStruct(entity model.Entity) (jen.Statement, error){
	
	// Vars
	var resp jen.Statement

	// Struct
	resp.Type()

	// ID
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntitySliceStructId(entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Index
	resp.Index()

	// Interface ID
	interfaceId , err := entityGenerator.formatter.OutputScaffoldDomainEntityInterfaceId(entity)
	if err != nil {
		return nil, err
	}
	resp.Qual("", interfaceId)

	return resp, nil

}

func (entityGenerator *entityGenerator) scaffoldEntityInterface(entity model.Entity) (jen.Statement, error){
	
	// Vars
	var resp jen.Statement
	var fields []jen.Code

	// Struct
	resp.Type()

	// Type
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntityInterfaceId(entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Getter
	for _, field := range entity.Fields {

		// Getter
		code, err := entityGenerator.scaffoldEntityInterfaceGetter(field, entity)
		if err != nil {
			return nil, err
		}
		fields = append(fields, code)
	}

	// Setter
	for _, field := range entity.Fields {

		// Setter
		code, err := entityGenerator.scaffoldEntityInterfaceSetter(field, entity)
		if err != nil {
			return nil, err
		}
		fields = append(fields, code)
	}

	// Set All
	code, err := entityGenerator.scaffoldEntityInterfaceSetAllSetter()
	if err != nil {
		return nil, err
	}
	fields = append(fields, code)

	// JSON
	if entity.JSON {

		code, err := entityGenerator.scaffoldEntityInterfaceMarshalJSON()
		if err != nil {
			return nil, err
		}
		fields = append(fields, code)

	}

	// JSON
	if entity.JSON {

		code, err := entityGenerator.scaffoldEntityInterfaceUnmarshalJSON()
		if err != nil {
			return nil, err
		}
		fields = append(fields, code)

	}


	// Interface
	resp.Interface(fields...)

	
	return resp, nil
}

func (entityGenerator *entityGenerator) scaffoldEntitySliceInterface(entity model.Entity) (jen.Statement, error){
	
	// Vars
	var resp jen.Statement
	var fields []jen.Code

	// Struct
	resp.Type()

	// Type
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntitySliceInterfaceId(entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Len
	lenCode, err := entityGenerator.scaffoldEntitySliceInterfaceLen()
	if err != nil {
		return nil, err
	}
	fields = append(fields, lenCode)

	// Append
	appendCode, err := entityGenerator.scaffoldEntitySliceInterfaceAppend(entity)
	if err != nil {
		return nil, err
	}
	fields = append(fields, appendCode)

	// Elements
	elementsCode, err := entityGenerator.scaffoldEntitySliceInterfaceElements(entity)
	if err != nil {
		return nil, err
	}
	fields = append(fields, elementsCode)


	// Interface
	resp.Interface(fields...)

	
	return resp, nil
}



func (entityGenerator *entityGenerator) scaffoldEntityInterfaceGetter(field model.Field, entity model.Entity) (jen.Code, error){
	
	// ID
	var statement jen.Statement
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntityGetterId(field)
	if err != nil {
		return nil, err
	}

	// Set
	statement.Id(id)

	// Params
	statement.Params()

	// Field
	err = entityGenerator.helperGenerator.Field("", field, entity, &statement)
	if err != nil {
		return nil, err
	}

	return &statement, nil

}

func (entityGenerator *entityGenerator) scaffoldEntityInterfaceSetter(field model.Field, entity model.Entity) (jen.Code, error){
	
	// ID
	var statement, args jen.Statement
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntitySetterId(field)
	if err != nil {
		return nil, err
	}

	// Field ID
	fieldId , err := entityGenerator.formatter.OutputScaffoldDomainEntityFieldId(field)
	if err != nil {
		return nil, err
	}

	// Set
	statement.Id(id)
	args.Id(fieldId)
	
	// Args
	err = entityGenerator.helperGenerator.Field("", field, entity, &args)
	if err != nil {
		return nil, err
	}

	// Params
	statement.Params(&args)
	
	return &statement, nil
}


func (entityGenerator *entityGenerator) scaffoldEntityInterfaceSetAllSetter() (jen.Code, error){
	
	// ID
	var statement jen.Statement
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntitySetAllSetterId()
	if err != nil {
		return nil, err
	}


	// Set
	statement.Id(id)

	// Params
	statement.Params()
	
	return &statement, nil
}

func (entityGenerator *entityGenerator) scaffoldEntityInterfaceMarshalJSON() (jen.Code, error){
	
	// ID
	var statement jen.Statement

	// Set
	statement.Id("MarshalJSON")

	// Params
	statement.Params()

	// Parens
	statement.Parens(
		jen.List(
			jen.Index().
			Byte(),
			jen.Error(),
		),
	)
	
	return &statement, nil
}

func (entityGenerator *entityGenerator) scaffoldEntityInterfaceUnmarshalJSON() (jen.Code, error){
	
	// ID
	var statement jen.Statement

	// Set
	statement.Id("UnmarshalJSON")

	// Params
	statement.Params(
		jen.Id("data").
		Index().
		Byte(),
	)

	// Parens
	statement.Parens(
		jen.List(
			jen.Error(),
		),
	)
	
	return &statement, nil
}

func (entityGenerator *entityGenerator) scaffoldEntityStructField(field model.Field, entity model.Entity) (jen.Code, error){
	
	// ID
	var statement jen.Statement
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntityFieldId(field)
	if err != nil {
		return nil, err
	}

	// Set
	statement.Id(id)

	// Field
	err = entityGenerator.helperGenerator.Field("", field, entity, &statement)
	if err != nil {
		return nil, err
	}

	return &statement, nil
}

func (entityGenerator *entityGenerator) scaffoldEntitySliceInterfaceLen() (jen.Code, error){
	
	// ID
	var statement jen.Statement
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntityLenId()
	if err != nil {
		return nil, err
	}

	// Set
	statement.Id(id)

	// Params
	statement.Params()

	// Parens
	statement.Parens(jen.List(
		jen.Int(),
	))
	
	return &statement, nil
}

func (entityGenerator *entityGenerator) scaffoldEntitySliceInterfaceAppend(entity model.Entity) (jen.Code, error){
	
	// ID
	var statement jen.Statement
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntityAppendId()
	if err != nil {
		return nil, err
	}

	// Set
	statement.Id(id)

	// Interface ID
	interfaceId , err := entityGenerator.formatter.OutputScaffoldDomainEntityInterfaceId(entity)
	if err != nil {
		return nil, err
	}

	// Params
	statement.Params(
		jen.Id("req").
		Qual("", interfaceId),
	)

	
	return &statement, nil
}

func (entityGenerator *entityGenerator) scaffoldEntitySliceInterfaceElements(entity model.Entity) (jen.Code, error){
	
	// ID
	var statement jen.Statement
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntityElementsId()
	if err != nil {
		return nil, err
	}

	// Set
	statement.Id(id)

	// Interface ID
	interfaceId , err := entityGenerator.formatter.OutputScaffoldDomainEntityInterfaceId(entity)
	if err != nil {
		return nil, err
	}

	// Params
	statement.Params()

	// Parens
	statement.Parens(
		jen.List(
			jen.Index().
			Qual("", interfaceId),
		),
	)
	
	return &statement, nil
}

func (entityGenerator *entityGenerator) scaffoldEntityInterfaceConstructorFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntityInterfaceConstructorFunctionId(entity)
	if err != nil {
		return nil, err
	}
	
	// Interface ID
	interfaceId , err := entityGenerator.formatter.OutputScaffoldDomainEntityInterfaceId(entity)
	if err != nil {
		return nil, err
	}

	// Struct ID
	structId , err := entityGenerator.formatter.OutputScaffoldDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}

	// Func
	statement.Func()

	// ID
	statement.Id(id)

	// Params
	statement.Params()

	// List
	statement.List(
		jen.Id(interfaceId),
	)

	// Block
	statement.Block(
		jen.Return(
			jen.Op("&").
			Id(structId).
			Values(),
		),
	)
	
	
	return statement, nil
}

func (entityGenerator *entityGenerator) scaffoldEntitySliceInterfaceConstructorFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntitySliceInterfaceConstructorFunctionId(entity)
	if err != nil {
		return nil, err
	}
	
	// Interface ID
	interfaceId , err := entityGenerator.formatter.OutputScaffoldDomainEntitySliceInterfaceId(entity)
	if err != nil {
		return nil, err
	}

	// Slice Struct ID
	sliceStructId , err := entityGenerator.formatter.OutputScaffoldDomainEntitySliceStructId(entity)
	if err != nil {
		return nil, err
	}

	// Func
	statement.Func()

	// ID
	statement.Id(id)

	// Params
	statement.Params()

	// List
	statement.List(
		jen.Id(interfaceId),
	)

	// Block
	statement.Block(
		jen.Return(
			jen.Op("&").
			Id(sliceStructId).
			Values(),
		),
	)
	
	
	return statement, nil
}

func (entityGenerator *entityGenerator) scaffoldEntitySliceInterfaceLenFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntityLenId()
	if err != nil {
		return nil, err
	}

	// Slice Struct ID
	sliceStructId , err := entityGenerator.formatter.OutputScaffoldDomainEntitySliceStructId(entity)
	if err != nil {
		return nil, err
	}

	// Func
	statement.Func()

	// Params
	statement.Params(
		jen.Id("m").
		Op("*").
		Qual("", sliceStructId),
	)

	// ID
	statement.Id(id)

	// Params
	statement.Params()

	// List
	statement.List(
		jen.Int(),
	)

	// Block
	statement.Block(
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
	
	return statement, nil
}

func (entityGenerator *entityGenerator) scaffoldEntitySliceInterfaceAppendFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntityAppendId()
	if err != nil {
		return nil, err
	}

	// Interface ID
	interfaceId , err := entityGenerator.formatter.OutputScaffoldDomainEntitySliceInterfaceId(entity)
	if err != nil {
		return nil, err
	}

	// Slice Struct ID
	sliceStructId , err := entityGenerator.formatter.OutputScaffoldDomainEntitySliceStructId(entity)
	if err != nil {
		return nil, err
	}

	// Func
	statement.Func()

	// Params
	statement.Params(
		jen.Id("m").
		Op("*").
		Qual("", sliceStructId),
	)

	// ID
	statement.Id(id)

	// Params
	statement.Params(
		jen.Id("req").
		Qual("", interfaceId),
	)

	// List
	statement.List(
		jen.Int(),
	)

	// Block
	statement.Block(
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
	
	return statement, nil
}

func (entityGenerator *entityGenerator) scaffoldEntitySliceInterfaceElementsFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntityElementsId()
	if err != nil {
		return nil, err
	}

	// Interface ID
	interfaceId , err := entityGenerator.formatter.OutputScaffoldDomainEntitySliceInterfaceId(entity)
	if err != nil {
		return nil, err
	}

	// Slice Struct ID
	sliceStructId , err := entityGenerator.formatter.OutputScaffoldDomainEntitySliceStructId(entity)
	if err != nil {
		return nil, err
	}

	// Func
	statement.Func()

	// Params
	statement.Params(
		jen.Id("m").
		Op("*").
		Qual("", sliceStructId),
	)

	// ID
	statement.Id(id)

	// Params
	statement.Params()

	// List
	statement.List(
		jen.Index().
		Id(interfaceId),
	)

	// Block
	statement.Block(
		jen.Return(
			jen.Op("*").
			Id("m"),
		),	
	)
	
	return statement, nil
}

func (entityGenerator *entityGenerator) scaffoldEntityInterfaceSetterFunction(field model.Field, entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement, args jen.Statement
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntitySetterId(field)
	if err != nil {
		return nil, err
	}

	// Struct ID
	structId , err := entityGenerator.formatter.OutputScaffoldDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}

	// Field ID
	fieldId , err := entityGenerator.formatter.OutputScaffoldDomainEntityFieldId(field)
	if err != nil {
		return nil, err
	}

	// Func
	statement.Func()

	// Params
	statement.Params(
		jen.Id("m").
		Op("*").
		Qual("", structId),
	)

	// ID
	statement.Id(id)

	// Args
	args.Id(fieldId)
	err = entityGenerator.helperGenerator.Field("", field, entity, &args)
	if err != nil {
		return nil, err
	}

	// Params
	statement.Params(&args)

	// Return
	statement.Block(
		jen.Id("m").
		Dot(fieldId).
		Op("=").
		Id(fieldId),
	)
	
	return statement, nil
}

func (entityGenerator *entityGenerator) scaffoldEntityInterfaceSetAllSetterFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement
	var block []jen.Code
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntitySetAllSetterId()
	if err != nil {
		return nil, err
	}

	// Interface ID
	interfaceId , err := entityGenerator.formatter.OutputScaffoldDomainEntityInterfaceId(entity)
	if err != nil {
		return nil, err
	}

	// Struct ID
	structId , err := entityGenerator.formatter.OutputScaffoldDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}

	// Func
	statement.Func()

	// Params
	statement.Params(
		jen.Id("m").
		Op("*").
		Qual("", structId),
	)

	// ID
	statement.Id(id)

	// Params
	statement.Params(
		jen.Id("req").
		Qual("", interfaceId),
	)

	// Block
	for _, field := range entity.Fields {
		
		setterId , err := entityGenerator.formatter.OutputScaffoldDomainEntitySetterId(field)
		if err != nil {
			return nil, err
		}
		getterId , err := entityGenerator.formatter.OutputScaffoldDomainEntityGetterId(field)
		if err != nil {
			return nil, err
		}
		block = append(block,
			jen.Id("m").
			Dot(setterId).
			Call(
				jen.Id("req").
				Dot(getterId).
				Call(),
			),
		)
		
	}

	// Block
	statement.Block(
		block...,
	)
	
	return statement, nil
}

func (entityGenerator *entityGenerator) scaffoldEntityInterfaceMarshalJSONFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement

	// Struct ID
	structId , err := entityGenerator.formatter.OutputScaffoldDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}

	// Func
	statement.Func()

	// Params
	statement.Params(
		jen.Id("m").
		Op("*").
		Qual("", structId),
	)

	// ID
	statement.Id("MarshalJSON")

	// Params
	statement.Params()

	// Parens
	statement.Parens(
		jen.List(
			jen.Index().
			Byte(),
			jen.Error(),
		),
	)	

	var jsonStruct []jen.Code
	jsonStructDict := make(jen.Dict)
	for _, field := range entity.Fields {

		// Getter ID
		getterId , err := entityGenerator.formatter.OutputScaffoldDomainEntityGetterId(field)
		if err != nil {
			return nil, err
		}
		
		code, err := entityGenerator.scaffoldEntityJSONStructField(field, entity)
		if err != nil {
			return nil, err
		}
		
		jsonStruct = append(jsonStruct, code)
		jsonStructDict[jen.Id(getterId)] = jen.Id("m").Dot(getterId).Call()
		
	}

	// Block
	statement.Block(

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
	
	
	return statement, nil
}

func (entityGenerator *entityGenerator) scaffoldEntityInterfaceUnmarshalJSONFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement

	// Struct ID
	structId , err := entityGenerator.formatter.OutputScaffoldDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}

	// Func
	statement.Func()

	// Params
	statement.Params(
		jen.Id("m").
		Op("*").
		Qual("", structId),
	)

	// ID
	statement.Id("UnmarshalJSON")

	// Params
	statement.Params(
		jen.Id("data").
		Index().
		Byte(),
	)

	// Parens
	statement.Parens(
		jen.List(
			jen.Error(),
		),
	)	

	var jsonStruct []jen.Code
	var jsonSetterFunctions []jen.Code
	for _, field := range entity.Fields {

		// Getter ID
		getterId , err := entityGenerator.formatter.OutputScaffoldDomainEntityGetterId(field)
		if err != nil {
			return nil, err
		}

		// Setter ID
		setterId , err := entityGenerator.formatter.OutputScaffoldDomainEntitySetterId(field)
		if err != nil {
			return nil, err
		}
		
		code, err := entityGenerator.scaffoldEntityJSONStructField(field, entity)
		if err != nil {
			return nil, err
		}

		jsonStruct = append(jsonStruct, code)

		jsonSetterFunctions = append(jsonSetterFunctions, jen.Id("m").
			Dot(setterId).
			Call(
				jen.Id("jsonStruct").
				Dot(getterId),
			),
		)
		
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

	statement.Block(
		block...,
	)
	
	
	return statement, nil
}

func (entityGenerator *entityGenerator) scaffoldEntityInterfaceGetterFunction(field model.Field, entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntityGetterId(field)
	if err != nil {
		return nil, err
	}

	// Struct ID
	structId , err := entityGenerator.formatter.OutputScaffoldDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}

	// Field ID
	fieldId , err := entityGenerator.formatter.OutputScaffoldDomainEntityFieldId(field)
	if err != nil {
		return nil, err
	}

	// Func
	statement.Func()

	// Params
	statement.Params(
		jen.Id("m").
		Op("*").
		Qual("", structId),
	)

	// ID
	statement.Id(id)

	// Params
	statement.Params()

	// Field
	err = entityGenerator.helperGenerator.Field("", field, entity, &statement)
	if err != nil {
		return nil, err
	}

	// Return
	statement.Block(
		jen.Return(
			jen.Id("m").
			Dot(fieldId),
		),
	)
	
	return statement, nil
}



func (entityGenerator *entityGenerator) scaffoldEntityJSONStructField(field model.Field, entity model.Entity) (jen.Code, error){
	
	// ID
	var statement jen.Statement
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntityGetterId(field)
	if err != nil {
		return nil, err
	}
	tagId , err := entityGenerator.formatter.OutputScaffoldDomainEntityJSONTagId(field)
	if err != nil {
		return nil, err
	}

	// Set
	statement.Id(id)

	// Field
	err = entityGenerator.helperGenerator.Field("", field, entity, &statement)
	if err != nil {
		return nil, err
	}

	// Tag
	statement.Tag(map[string]string{"json": tagId})

	return &statement, nil
}