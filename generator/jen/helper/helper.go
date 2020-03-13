package helper

import(
	"github.com/dave/jennifer/jen"
	"github.com/brianshepanek/turnbull/domain/model"
	"github.com/brianshepanek/turnbull/formatter"
)

type generator struct{
	formatter formatter.Formatter
}

type Generator interface{

	// General
	Field(id string, field model.Field, entity model.Entity, statement *jen.Statement) (error)

}	

func New(formatter formatter.Formatter) Generator {
	return &generator{
		formatter : formatter,
	}
}

func (generator *generator) Field(id string, field model.Field, entity model.Entity, statement *jen.Statement) (error){

	// ID
	if id != "" {
		statement.Id(id)
	}
	
	// Op
	if field.Op != "" {
		statement.Op(field.Op)
	}

	// Slice
	if field.Type != "self" {
		if field.Slice {
			statement.Index()
		}
	}

	// Qual
	if field.Type != "" {
		if field.Type == "self" {

			// Vars
			var id string
			// Import Path
			importPath , err := generator.formatter.OutputScaffoldDomainEntityDirectoryImportPath()
			if err != nil {
				return nil
			}

			if field.Slice {

				// Slice Interface ID
				id, err = generator.formatter.OutputScaffoldDomainEntitySliceInterfaceId(entity)
				if err != nil {
					return nil
				}
		
			} else {

				// Interface ID
				id , err = generator.formatter.OutputScaffoldDomainEntityInterfaceId(entity)
				if err != nil {
					return nil
				}

			}

			// Set
			statement.Qual(importPath, id)
		} else {
			statement.Qual(field.Package, field.Type)
		}
	} 

	
	return nil
}