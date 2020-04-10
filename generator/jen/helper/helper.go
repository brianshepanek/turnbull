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
	PrimaryField(entity model.Entity) (*model.Field, error)

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
	if !entity.Interface && field.Type == "self" {
		statement.Op("*")
	}

	// Slice
	if entity.Interface {
		if field.Type != "self" {
			if field.Slice {
				statement.Index()
			}
		}
	} else {
		if field.Slice {
			statement.Index()
		}
	}
	

	// Qual
	if field.Embedded {
		
		// Vars
		var entityId string

		// Interface ID
		if field.Entity.Interface {
			structId , err := generator.formatter.OutputDomainEntityStructId(field.Entity)
			if err != nil {
				return err
			}
			entityId = structId
		} else {
			interfaceId , err := generator.formatter.OutputDomainEntityInterfaceId(field.Entity)
			if err != nil {
				return err
			}
			entityId = interfaceId
		}
		statement.Qual("", entityId)


	} else {
		if field.Type != "" {
			if field.Type == "self" {
	
				// Vars
				var id string
	
				// Import Path
				importPath , err := generator.formatter.OutputScaffoldDomainEntityDirectoryImportPath()
				if err != nil {
					return nil
				}
	
				if field.Slice && entity.Interface {
	
					// Slice Interface ID
					id, err = generator.formatter.OutputDomainEntitySliceInterfaceId(entity)
					if err != nil {
						return nil
					}
			
				} else {
	
					// Interface ID
					id , err = generator.formatter.OutputDomainEntityInterfaceId(entity)
					if err != nil {
						return nil
					}
	
				}
	
				// Set
				statement.Qual(importPath, id)
	
			} else if (field.Type == "primary"){ 
				
				p, t, err := generator.primary(entity)
				if err != nil {
					return nil
				}
				statement.Qual(p, t)
	
			} else {
				statement.Qual(field.Package, field.Type)
			}
		} 
	}
	

	
	return nil
}

func (generator *generator) PrimaryField(entity model.Entity) (*model.Field, error){
	var primaryField model.Field

	for _, field := range entity.Fields {
		if field.Primary {
			return &primaryField, nil
		}
	}
	for _, field := range entity.Fields {
		if field.Embedded {
			for _, embeddedField := range field.Entity.Fields {
				if embeddedField.Primary {
					return &embeddedField, nil
				}
			}
		}
	}

	return nil, nil
}

func (generator *generator) primary(entity model.Entity) (string, string, error){
	var p, t string

	for _, field := range entity.Fields {
		if field.Primary {
			return field.Package, field.Type, nil
		}
	}
	for _, field := range entity.Fields {
		if field.Embedded {
			for _, embeddedField := range field.Entity.Fields {
				if embeddedField.Primary {
					return embeddedField.Package, embeddedField.Type, nil
				}
			}
		}
	}

	return p, t , nil
}