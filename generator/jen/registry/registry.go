package registry

import(
	"github.com/dave/jennifer/jen"
	"github.com/brianshepanek/turnbull/domain/model"
	"github.com/brianshepanek/turnbull/formatter"
	"github.com/brianshepanek/turnbull/generator/jen/helper"
)

type registryGenerator struct{
	formatter formatter.Formatter
	helperGenerator helper.Generator
}

type RegistryGenerator interface{

	// File
	File(entities []model.Entity) (*jen.File, error)
	ScaffoldFile(entities []model.Entity) (*jen.File, error)

}

func New(formatter formatter.Formatter, helperGenerator helper.Generator) RegistryGenerator {
	return &registryGenerator{
		formatter : formatter,
		helperGenerator : helperGenerator,
	}
}

func (registryGenerator *registryGenerator) File(entities []model.Entity) (*jen.File, error){
	
	
	// File
	packageName , err := registryGenerator.formatter.OutputRegistryPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)

	return f, nil
}

func (registryGenerator *registryGenerator) ScaffoldFile(entities []model.Entity) (*jen.File, error){
	
	// Vars
	var resp jen.Statement
	var fields []jen.Code
	
	// File
	packageName , err := registryGenerator.formatter.OutputRegistryPackageName()
	if err != nil {
		return nil, err
	}
	f := jen.NewFile(packageName)

	for _, entity := range entities {
		for _, repository := range entity.Repositories {
			
			// ID
			registryStructId , err := registryGenerator.formatter.OutputScaffoldInterfaceRepositoryRegistryStructId(repository.Type, entity)
			if err != nil {
				return nil, err
			}

			fields = append(fields, jen.Id(registryStructId))

		}

		for _, presenter := range entity.Presenters {
			
			// ID
			registryStructId , err := registryGenerator.formatter.OutputScaffoldInterfacePresenterRegistryStructId(presenter.Type, entity)
			if err != nil {
				return nil, err
			}

			fields = append(fields, jen.Id(registryStructId))

		}
		
	}

	// Type
	resp.Type()

	// ID
	resp.Id(packageName)

	// Struct
	resp.Struct(fields...)

	f.Add(&resp)

	return f, nil
}