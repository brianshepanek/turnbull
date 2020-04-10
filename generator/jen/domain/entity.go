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
	File(entity model.Entity) (*jen.File, error)
	ScaffoldFile(entity model.Entity) (*jen.File, error)

	// Entity
	entityStruct(entity model.Entity) (jen.Statement, error)
	entitySliceStruct(entity model.Entity) (jen.Statement, error)
	entityInterface(entity model.Entity) (jen.Statement, error)
	entitySliceInterface(entity model.Entity) (jen.Statement, error)
	entityInterfaceConstructorFunction(entity model.Entity) (jen.Statement, error)
	entitySliceInterfaceConstructorFunction(entity model.Entity) (jen.Statement, error)

	// Scaffold Entity Struct
	scaffoldEntityStruct(entity model.Entity) (jen.Statement, error)
	scaffoldEntitySliceStruct(entity model.Entity) (jen.Statement, error)

	scaffoldEntityStructField(field model.Field, entity model.Entity) (jen.Code, error)

	// Scaffold Entity Interface
	scaffoldEntityInterface(entity model.Entity) (jen.Statement, error)

	// Scaffold Entity Interface Methods
	scaffoldEntityInterfaceGetter(field model.Field, entity model.Entity) (jen.Code, error)
	scaffoldEntityInterfaceSetter(field model.Field, entity model.Entity) (jen.Code, error)
	scaffoldEntityInterfaceSetAllSetter(entity model.Entity) (jen.Code, error)

	// Scaffold Entity Slice Interface
	scaffoldEntitySliceInterface(entity model.Entity) (jen.Statement, error)

	// Scaffold Entity Slice Interface Methods
	scaffoldEntitySliceInterfaceLen() (jen.Code, error)
	scaffoldEntitySliceInterfaceAppend(entity model.Entity) (jen.Code, error)
	scaffoldEntitySliceInterfaceElements(entity model.Entity) (jen.Code, error)

	// Scaffold Entity Interface Functions
	scaffoldEntityStructConstructorFunction(entity model.Entity) (jen.Statement, error)
	scaffoldEntityScaffoldStructConstructorFunction(entity model.Entity) (jen.Statement, error)
	scaffoldEntityInterfaceConstructorFunction(entity model.Entity) (jen.Statement, error)
	scaffoldEntitySliceInterfaceConstructorFunction(entity model.Entity) (jen.Statement, error)
	scaffoldEntitySliceInterfaceLenFunction(entity model.Entity) (jen.Statement, error)
	scaffoldEntitySliceInterfaceAppendFunction(entity model.Entity) (jen.Statement, error)
	scaffoldEntitySliceInterfaceElementsFunction(entity model.Entity) (jen.Statement, error)
	scaffoldEntityInterfaceSetterFunction(field model.Field, entity model.Entity) (jen.Statement, error)
	scaffoldEntityInterfaceGetterFunction(field model.Field, entity model.Entity) (jen.Statement, error)
	scaffoldEntityInterfaceSetAllSetterFunction(entity model.Entity) (jen.Statement, error)

}

func NewEntityGenerator(formatter formatter.Formatter, helperGenerator helper.Generator) EntityGenerator {
	return &entityGenerator{
		formatter : formatter,
		helperGenerator : helperGenerator,
	}
}

func (entityGenerator *entityGenerator) File(entity model.Entity) (*jen.File, error){
	
	
	// File
	packageName , err := entityGenerator.formatter.OutputScaffoldDomainEntityPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)

	// Struct
	entityStruct, err := entityGenerator.entityStruct(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&entityStruct)
	f.Line()

	// Interface
	if entity.Interface {
		
		// Slice Struct
		entitySliceStruct, err := entityGenerator.entitySliceStruct(entity)
		if err != nil {
			return nil, err
		}
		f.Add(&entitySliceStruct)
		f.Line()

		// Interface
		entityInterface, err := entityGenerator.entityInterface(entity)
		if err != nil {
			return nil, err
		}
		f.Add(&entityInterface)

		// Slice Interface
		entitySliceInterface, err := entityGenerator.entitySliceInterface(entity)
		if err != nil {
			return nil, err
		}
		f.Add(&entitySliceInterface)
	}	

	// Interface Constructor Function
	interfaceConstructorFunction, err := entityGenerator.entityInterfaceConstructorFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&interfaceConstructorFunction)
	f.Line()

	// Slice Interface Constructor Function
	sliceInterfaceConstructorFunction, err :=  entityGenerator.entitySliceInterfaceConstructorFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&sliceInterfaceConstructorFunction)
	f.Line()
	
	
	return f, nil
}

func (entityGenerator *entityGenerator) ScaffoldFile(entity model.Entity) (*jen.File, error){
	
	
	// File
	packageName , err := entityGenerator.formatter.OutputScaffoldDomainEntityPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)

	// Struct
	scaffoldEntityStruct, err := entityGenerator.scaffoldEntityStruct(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&scaffoldEntityStruct)
	f.Line()

	// Struct Constructor Struct
	entityStructConstructor, err := entityGenerator.scaffoldEntityStructConstructorFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&entityStructConstructor)
	f.Line()

	// Scaffold Struct Constructor Struct
	scaffoldEntityStructConstructor, err := entityGenerator.scaffoldEntityScaffoldStructConstructorFunction(entity)
	if err != nil {
		return nil, err
	}
	f.Add(&scaffoldEntityStructConstructor)
	f.Line()


	// Interface
	if entity.Interface {

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
			if !field.Embedded {
				entityGetter, err := entityGenerator.scaffoldEntityInterfaceGetterFunction(field, entity)
				if err != nil {
					return nil, err
				}
				f.Add(&entityGetter)
				f.Line()
			}
			
		}
		
		// Setters
		for _, field := range entity.Fields {
			if !field.Embedded {
				entitySetter, err := entityGenerator.scaffoldEntityInterfaceSetterFunction(field, entity)
				if err != nil {
					return nil, err
				}
				f.Add(&entitySetter)
				f.Line()
			}	
		}

		// // Set All Setter
		// setAllSetter, err := entityGenerator.scaffoldEntityInterfaceSetAllSetterFunction(entity)
		// if err != nil {
		// 	return nil, err
		// }
		// f.Add(&setAllSetter)
		// f.Line()

	}

	// Callbacks
	for _, method := range entity.Methods {
		for _, callback := range method.Callbacks {
			entityCallback, err := entityGenerator.scaffoldEntityInterfaceCallbackFunction(callback, method, entity)
			if err != nil {
				return nil, err
			}
			f.Add(&entityCallback)
			f.Line()
		}
	}

	// To Primary
	if entityGenerator.hasPrimary(entity) {
		toPrimary, err := entityGenerator.scaffoldEntityInterfaceToPrimaryFunction(entity)
		if err != nil {
			return nil, err
		}
		f.Add(&toPrimary)
		f.Line()
	}
	

	
	return f, nil
}

func (entityGenerator *entityGenerator) hasPrimary(entity model.Entity) (bool){
	var hasPrimary bool
	for _, field := range entity.Fields {
		if field.Primary {
			hasPrimary = true
		}
	}
	return hasPrimary
}

func (entityGenerator *entityGenerator) entityStruct(entity model.Entity) (jen.Statement, error){
	
	// Vars
	var resp jen.Statement
	var fields []jen.Code

	// Struct
	resp.Type()

	// Type
	id , err := entityGenerator.formatter.OutputDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Scaffold
	scaffoldId , err := entityGenerator.formatter.OutputScaffoldDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}
	
	// Fields
	fields = append(fields, jen.Op("*").Id(scaffoldId))

	// Struct
	resp.Struct(fields...)

	
	return resp, nil
}

func (entityGenerator *entityGenerator) entitySliceStruct(entity model.Entity) (jen.Statement, error){
	
	// Vars
	var resp jen.Statement
	var fields []jen.Code

	// Struct
	resp.Type()

	// Type
	id , err := entityGenerator.formatter.OutputDomainEntitySliceStructId(entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Scaffold
	scaffoldId , err := entityGenerator.formatter.OutputScaffoldDomainEntitySliceStructId(entity)
	if err != nil {
		return nil, err
	}
	
	// Fields
	fields = append(fields, jen.Id(scaffoldId))

	// Struct
	resp.Struct(fields...)

	
	return resp, nil
}

func (entityGenerator *entityGenerator) entityInterface(entity model.Entity) (jen.Statement, error){
	
	// Vars
	var resp jen.Statement
	var fields []jen.Code

	// Struct
	resp.Type()

	// Type
	id , err := entityGenerator.formatter.OutputDomainEntityInterfaceId(entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Scaffold
	scaffoldId , err := entityGenerator.formatter.OutputScaffoldDomainEntityInterfaceId(entity)
	if err != nil {
		return nil, err
	}

	// Fields
	fields = append(fields, jen.Id(scaffoldId))

	// Interface
	resp.Interface(fields...)

	
	return resp, nil
}

func (entityGenerator *entityGenerator) entitySliceInterface(entity model.Entity) (jen.Statement, error){
	
	// Vars
	var resp jen.Statement
	var fields []jen.Code

	// Struct
	resp.Type()

	// Type
	id , err := entityGenerator.formatter.OutputDomainEntitySliceInterfaceId(entity)
	if err != nil {
		return nil, err
	}
	resp.Id(id)

	// Scaffold
	scaffoldId , err := entityGenerator.formatter.OutputScaffoldDomainEntitySliceInterfaceId(entity)
	if err != nil {
		return nil, err
	}

	// Fields
	fields = append(fields, jen.Id(scaffoldId))

	// Interface
	resp.Interface(fields...)

	
	return resp, nil
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
		// if !field.Embedded {
			code, err := entityGenerator.scaffoldEntityStructField(field, entity)
			if err != nil {
				return nil, err
			}
			fields = append(fields, code)
		// }	
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

		if !field.Embedded {

			// Getter
			code, err := entityGenerator.scaffoldEntityInterfaceGetter(field, entity)
			if err != nil {
				return nil, err
			}
			fields = append(fields, code)

		} else {

			if field.Entity.Interface {

				interfaceId, err := entityGenerator.formatter.OutputDomainEntityInterfaceId(field.Entity)
				if err != nil {
					return nil, err
				}

				fields = append(fields, jen.Id(interfaceId))
				
			}
			

		}

		
	}

	// Setter
	for _, field := range entity.Fields {

		if !field.Embedded {

			// Setter
			code, err := entityGenerator.scaffoldEntityInterfaceSetter(field, entity)
			if err != nil {
				return nil, err
			}
			fields = append(fields, code)

		}	
	}

	// Callbacks
	for _, method := range entity.Methods {
		for _, callback := range method.Callbacks {

			// Callback
			code, err := entityGenerator.scaffoldEntityInterfaceCallback(callback, method, entity)
			if err != nil {
				return nil, err
			}
			fields = append(fields, code)

		}
	}

	// // Set All
	// setAllCode, err := entityGenerator.scaffoldEntityInterfaceSetAllSetter(entity)
	// if err != nil {
	// 	return nil, err
	// }
	// fields = append(fields, setAllCode)

	// To Primary
	if entityGenerator.hasPrimary(entity) {
		toPrimaryCode, err := entityGenerator.scaffoldEntityInterfaceToPrimary(entity)
		if err != nil {
			return nil, err
		}
		fields = append(fields, toPrimaryCode)
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
	field.Op = "*"
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
	field.Op = "*"
	err = entityGenerator.helperGenerator.Field("", field, entity, &args)
	if err != nil {
		return nil, err
	}

	// Params
	statement.Params(&args)
	
	return &statement, nil
}

func (entityGenerator *entityGenerator) scaffoldEntityInterfaceCallback(callback model.Callback, method model.Method, entity model.Entity) (jen.Code, error){
	
	// ID
	var statement, args, vals jen.Statement
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntityCallbackId(callback, method)
	if err != nil {
		return nil, err
	}

	// ID
	statement.Id(id)

	// Args
	args.Id("ctx")
	err = entityGenerator.helperGenerator.Field("", model.Field{Package : "context", Type : "Context"}, entity, &args)
	if err != nil {
		return nil, err
	}

	// Params
	statement.Params(&args)

	// Vals
	err = entityGenerator.helperGenerator.Field("", model.Field{Type : "error"}, entity, &vals)
	if err != nil {
		return nil, err
	}

	// List
	statement.List(&vals)


	return &statement, nil

}

func (entityGenerator *entityGenerator) scaffoldEntityInterfaceToPrimary(entity model.Entity) (jen.Code, error){
	
	// ID
	var statement, vals jen.Statement
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntityToPrimaryId()
	if err != nil {
		return nil, err
	}
	
	// Set
	statement.Id(id)

	// Params
	statement.Params(
		jen.List(
			jen.Id("ctx").
			Qual("context", "Context"),

			jen.Id("req").
			Interface(),
		),
	)

	// Vals
	for _, field := range entity.Fields {
		if field.Primary {
			err = entityGenerator.helperGenerator.Field("", field, entity, &vals)
			if err != nil {
				return nil, err
			}
		}
	}
	err = entityGenerator.helperGenerator.Field("", model.Field{Type : "error"}, entity, &vals)

	// Set
	statement.Parens(
		jen.List(vals...),
	)
	
	return &statement, nil
}

func (entityGenerator *entityGenerator) scaffoldEntityInterfaceSetAllSetter(entity model.Entity) (jen.Code, error){
	
	// ID
	var statement jen.Statement
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntitySetAllSetterId()
	if err != nil {
		return nil, err
	}

	// Interface ID
	interfaceId , err := entityGenerator.formatter.OutputScaffoldDomainEntityInterfaceId(entity)
	if err != nil {
		return nil, err
	}

	// Set
	statement.Id(id)

	// Params
	statement.Params(
		jen.Id("req").
		Qual("", interfaceId),
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
	field.Op = "*"
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

func (entityGenerator *entityGenerator) scaffoldEntityStructConstructorFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntityStructConstructorFunctionId(entity)
	if err != nil {
		return nil, err
	}

	scaffoldStructConstructorId , err := entityGenerator.formatter.OutputScaffoldDomainEntityScaffoldStructConstructorFunctionId(entity)
	if err != nil {
		return nil, err
	}
	
	// Struct ID
	structId , err := entityGenerator.formatter.OutputDomainEntityStructId(entity)
	if err != nil {
		return nil, err
	}

	// Scaffold Struct ID
	scaffoldStructId , err := entityGenerator.formatter.OutputScaffoldDomainEntityStructId(entity)
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
	var list []jen.Code
	list = append(list, jen.Op("*").Id(structId))
	statement.List(
		list...,
	)

	// Block
	statement.Block(
		jen.Return(
			jen.Op("&").
			Id(structId).
			Values(
				jen.Dict{
					jen.Id(scaffoldStructId) : jen.Id(scaffoldStructConstructorId).Call(),
				},
			),
		),
	)
	
	
	return statement, nil
}

func (entityGenerator *entityGenerator) scaffoldEntityScaffoldStructConstructorFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntityScaffoldStructConstructorFunctionId(entity)
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
	var list []jen.Code
	list = append(list, jen.Op("*").Id(structId))
	statement.List(
		list...,
	)

	
	dict := make(jen.Dict)
	for _, field := range entity.Fields {
		if field.Embedded {

			// Vars
			var entityId string

			// Interface ID
			if field.Entity.Interface {
				structId , err := entityGenerator.formatter.OutputDomainEntityStructId(field.Entity)
				if err != nil {
					return nil, err
				}
				entityId = structId
			} else {
				interfaceId , err := entityGenerator.formatter.OutputDomainEntityInterfaceId(field.Entity)
				if err != nil {
					return nil, err
				}
				entityId = interfaceId
			}
			

			constructorId , err := entityGenerator.formatter.OutputScaffoldDomainEntityStructConstructorFunctionId(field.Entity)
			if err != nil {
				return nil, err
			}

			dict[jen.Id(entityId)] = jen.Id(constructorId).Call()

		}
	}

	// Block
	statement.Block(
		jen.Return(
			jen.Op("&").
			Id(structId).
			Values(
				dict,
			),
		),
	)
	
	
	return statement, nil
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
	var list []jen.Code
	if entity.Interface {
		list = append(list, jen.Id(interfaceId))
	} else {
		list = append(list, jen.Op("*").Id(interfaceId))
	}
	statement.List(
		list...,
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

func (entityGenerator *entityGenerator) entityInterfaceConstructorFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement
	id , err := entityGenerator.formatter.OutputDomainEntityInterfaceConstructorFunctionId(entity)
	if err != nil {
		return nil, err
	}
	
	// Interface ID
	interfaceId , err := entityGenerator.formatter.OutputDomainEntityInterfaceId(entity)
	if err != nil {
		return nil, err
	}

	// Struct ID
	// structId , err := entityGenerator.formatter.OutputDomainEntityStructId(entity)
	// if err != nil {
	// 	return nil, err
	// }

	// Struct Constructor ID
	structConstructorId , err := entityGenerator.formatter.OutputScaffoldDomainEntityStructConstructorFunctionId(entity)
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
	var list []jen.Code
	if entity.Interface {
		list = append(list, jen.Id(interfaceId))
	} else {
		list = append(list, jen.Op("*").Id(interfaceId))
	}
	statement.List(
		list...,
	)

	// Block
	statement.Block(
		jen.Return(
			jen.Id(structConstructorId).Call(),
		),
	)
	
	
	return statement, nil
}

func (entityGenerator *entityGenerator) entitySliceInterfaceConstructorFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement jen.Statement
	id , err := entityGenerator.formatter.OutputDomainEntitySliceInterfaceConstructorFunctionId(entity)
	if err != nil {
		return nil, err
	}
	
	// Interface ID
	var interfaceId string
	if entity.Interface {
		interfaceId , err = entityGenerator.formatter.OutputDomainEntitySliceInterfaceId(entity)
		if err != nil {
			return nil, err
		}
	} else {
		interfaceId , err = entityGenerator.formatter.OutputDomainEntityInterfaceId(entity)
		if err != nil {
			return nil, err
		}
	}	

	// Slice Struct ID
	sliceStructId , err := entityGenerator.formatter.OutputDomainEntitySliceStructId(entity)
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
	var list []jen.Code
	if entity.Interface {
		list = append(list, jen.Id(interfaceId))
	} else {
		list = append(list, jen.Op("*").Index().Id(interfaceId))
	}
	statement.List(
		list...,
	)

	// Return
	var returnValues []jen.Code
	if entity.Interface {
		returnValues = append(returnValues, 
			jen.Op("&").
			Id(sliceStructId).
			Values(),
		)
	} else {
		returnValues = append(returnValues, 
			jen.Op("&").
			Index().
			Id(interfaceId).
			Values(),
		)
	}
	// Block
	statement.Block(
		jen.Return(
			returnValues...
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
	interfaceId , err := entityGenerator.formatter.OutputScaffoldDomainEntityInterfaceId(entity)
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
	interfaceId , err := entityGenerator.formatter.OutputScaffoldDomainEntityInterfaceId(entity)
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
	field.Op = "*"
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

func (entityGenerator *entityGenerator) scaffoldEntityInterfaceCallbackFunction(callback model.Callback, method model.Method, entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement, args, vals jen.Statement
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntityCallbackId(callback, method)
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

	// Args
	args.Id("ctx")
	err = entityGenerator.helperGenerator.Field("", model.Field{Package : "context", Type : "Context"}, entity, &args)
	if err != nil {
		return nil, err
	}

	// Params
	statement.Params(&args)

	// Vals
	err = entityGenerator.helperGenerator.Field("", model.Field{Type : "error"}, entity, &vals)
	if err != nil {
		return nil, err
	}

	// List
	statement.List(&vals)

	// Return
	statement.Block(
		jen.Return(
			jen.Nil(),
		),	
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

		if !field.Embedded {

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
		
		
	}

	// Block
	statement.Block(
		block...,
	)
	
	return statement, nil
}

func (entityGenerator *entityGenerator) scaffoldEntityInterfaceToPrimaryFunction(entity model.Entity) (jen.Statement, error){
	
	// ID
	var statement, vals jen.Statement
	id , err := entityGenerator.formatter.OutputScaffoldDomainEntityToPrimaryId()
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
	
	// Set
	statement.Id(id)

	// Params
	statement.Params(
		jen.List(
			jen.Id("ctx").
			Qual("context", "Context"),

			jen.Id("req").
			Interface(),
		),
	)

	// Vals
	var primaryField model.Field
	for _, field := range entity.Fields {
		if field.Primary {
			primaryField = field
		}
	}
	err = entityGenerator.helperGenerator.Field("", primaryField, entity, &vals)
	if err != nil {
		return nil, err
	}
	err = entityGenerator.helperGenerator.Field("", model.Field{Type : "error"}, entity, &vals)
	if err != nil {
		return nil, err
	}

	// Set
	statement.Parens(
		jen.List(vals...),
	)

	// Block
	var block []jen.Code
	block = append(block, 
		jen.Var().
		Id("resp").
		Qual(primaryField.Package, primaryField.Type),
	)
	block = append(block, 
		jen.Return(
			jen.List(
				jen.Id("resp"),
				jen.Nil(),
			),
		),
	)
	statement.Block(block...)
	
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
	field.Op = "*"
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



