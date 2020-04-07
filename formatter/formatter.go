package formatter

import(
	"os"
	"strings"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"github.com/brianshepanek/turnbull/config"
	"github.com/brianshepanek/turnbull/domain/model"
)

type formatter struct{
	config *config.Config
}

type Formatter interface{

	OutputDirectory() (string, error)
	OutputScaffoldDirectory() (string, error)

	OutputRegistryPackageName() (string, error)
	OutputRegistryDirectory() (string, error)
	OutputRegistryFile() (string, error)
	OutputRegistryEntityFile(entity model.Entity) (string, error)
	OutputRegistryEntityPresenterConstructorFunctionId(entity model.Entity) (string, error)
	OutputRegistryEntityRepositoryConstructorFunctionId(entity model.Entity) (string, error)
	OutputRegistryEntityControllerConstructorFunctionId(driver string, entity model.Entity) (string, error)

	OutputScaffoldDomainDirectory() (string, error)
	OutputScaffoldDomainEntityDirectory() (string, error)
	OutputScaffoldDomainEntityDirectoryImportPath() (string, error)
	OutputDomainEntityFile(entity model.Entity) (string, error)
	OutputScaffoldDomainEntityFile(entity model.Entity) (string, error)

	OutputScaffoldUsecaseDirectory() (string, error)
	OutputScaffoldUsecaseInteractorDirectory() (string, error)
	OutputScaffoldUsecaseInteractorDirectoryImportPath() (string, error)
	OutputUsecaseInteractorFile(entity model.Entity) (string, error)
	OutputScaffoldUsecaseInteractorFile(entity model.Entity) (string, error)
	OutputScaffoldUsecaseRepositoryDirectory() (string, error)
	OutputScaffoldUsecaseRepositoryDirectoryImportPath() (string, error)
	OutputUsecaseRepositoryFile(entity model.Entity) (string, error)
	OutputScaffoldUsecaseRepositoryFile(entity model.Entity) (string, error)
	OutputScaffoldUsecasePresenterDirectory() (string, error)
	OutputScaffoldUsecasePresenterDirectoryImportPath() (string, error)
	OutputUsecasePresenterFile(entity model.Entity) (string, error)
	OutputScaffoldUsecasePresenterFile(entity model.Entity) (string, error)

	OutputScaffoldInterfaceDirectory() (string, error)
	
	OutputScaffoldInterfaceControllerDirectory() (string, error)
	OutputInterfaceControllerDirectory(driver string, entity model.Entity) (string, error)
	OutputInterfaceControllerDirectoryImportPath(driver string, entity model.Entity) (string, error)
	OutputInterfaceControllerFile(driver string, entity model.Entity) (string, error)
	OutputInterfaceControllerEntityFile(driver string, entity model.Entity) (string, error)
	OutputInterfaceControllerRegistryFile(driver string, entity model.Entity) (string, error)
	OutputScaffoldInterfaceControllerFile(driver string, entity model.Entity) (string, error)

	OutputScaffoldInterfaceRepositoryDirectory() (string, error)
	OutputInterfaceRepositoryFile(driver string, entity model.Entity) (string, error)
	OutputScaffoldInterfaceRepositoryFile(driver string, entity model.Entity) (string, error)
	OutputInterfaceRepositoryEntityFile(driver string, entity model.Entity) (string, error)
	OutputInterfaceRepositoryRegistryFile(driver string, entity model.Entity) (string, error)
	OutputInterfaceRepositoryDirectory(driver string, entity model.Entity) (string, error)
	OutputInterfaceRepositoryDirectoryImportPath(driver string, entity model.Entity) (string, error)

	OutputScaffoldInterfacePresenterDirectory() (string, error)
	OutputInterfacePresenterFile(driver string, entity model.Entity) (string, error)
	OutputScaffoldInterfacePresenterFile(driver string, entity model.Entity) (string, error)
	OutputInterfacePresenterEntityFile(driver string, entity model.Entity) (string, error)
	OutputInterfacePresenterRegistryFile(driver string, entity model.Entity) (string, error)
	OutputInterfacePresenterDirectory(driver string, entity model.Entity) (string, error)
	OutputInterfacePresenterDirectoryImportPath(driver string, entity model.Entity) (string, error)

	OutputScaffoldDomainEntityPackageName() (string, error)
	OutputDomainEntityStructId(entity model.Entity) (string, error)
	OutputDomainEntityLocalStructId(entity model.Entity) (string, error)
	OutputDomainEntitySliceStructId(entity model.Entity) (string, error)
	OutputDomainEntityInterfaceId(entity model.Entity) (string, error)
	OutputDomainEntitySliceInterfaceId(entity model.Entity) (string, error)
	OutputScaffoldDomainEntityStructId(entity model.Entity) (string, error)
	OutputScaffoldDomainEntitySliceStructId(entity model.Entity) (string, error)
	OutputScaffoldDomainEntityInterfaceId(entity model.Entity) (string, error)
	OutputScaffoldDomainEntitySliceInterfaceId(entity model.Entity) (string, error)
	OutputScaffoldDomainEntityFieldId(field model.Field) (string,  error)
	OutputScaffoldDomainEntityGetterId(field model.Field) (string,  error)
	OutputScaffoldDomainEntitySetterId(field model.Field) (string,  error)
	OutputScaffoldDomainEntityCallbackId(callback model.Callback, method model.Method) (string,  error)
	OutputScaffoldDomainEntitySetAllSetterId() (string,  error)
	OutputScaffoldDomainEntityToPrimaryId() (string,  error)
	OutputScaffoldDomainEntityLenId() (string,  error)
	OutputScaffoldDomainEntityAppendId() (string,  error)
	OutputScaffoldDomainEntityElementsId() (string,  error)
	OutputScaffoldDomainEntityJSONTagId(field model.Field) (string,  error)
	
	OutputDomainEntityInterfaceConstructorFunctionId(entity model.Entity) (string, error)
	OutputDomainEntitySliceInterfaceConstructorFunctionId(entity model.Entity) (string, error)
	OutputScaffoldDomainEntityInterfaceConstructorFunctionId(entity model.Entity) (string, error)
	OutputScaffoldDomainEntitySliceInterfaceConstructorFunctionId(entity model.Entity) (string, error)

	OutputScaffoldUsecaseRepositoryPackageName() (string, error)
	OutputUsecaseRepositoryInterfaceId(entity model.Entity) (string, error)
	OutputScaffoldUsecaseRepositoryInterfaceId(entity model.Entity) (string, error)
	OutputScaffoldUsecaseRepositoryInterfaceMethodId(method model.Method) (string, error)

	OutputScaffoldUsecasePresenterPackageName() (string, error)
	OutputUsecasePresenterInterfaceId(entity model.Entity) (string, error)
	OutputScaffoldUsecasePresenterInterfaceId(entity model.Entity) (string, error)
	OutputScaffoldUsecasePresenterInterfaceMethodId(method model.Method) (string, error)

	OutputScaffoldUsecaseInteractorPackageName() (string, error)
	OutputUsecaseInteractorStructId(entity model.Entity) (string, error)
	OutputScaffoldUsecaseInteractorStructId(entity model.Entity) (string, error)
	OutputUsecaseInteractorInterfaceId(entity model.Entity) (string, error)
	OutputScaffoldUsecaseInteractorInterfaceId(entity model.Entity) (string, error)
	OutputScaffoldUsecaseInteractorInterfaceMethodId(method model.Method) (string, error)
	OutputUsecaseInteractorInterfaceConstructorFunctionId(entity model.Entity) (string, error)
	OutputScaffoldUsecaseInteractorRegistryLocalConstructorFunctionId(entity model.Entity) (string, error)

	OutputScaffoldInterfaceRepositoryPackageName() (string, error)
	OutputInterfaceRepositoryStructId(driver string, entity model.Entity) (string, error)
	OutputScaffoldInterfaceRepositoryStructId(driver string, entity model.Entity) (string, error)
	OutputScaffoldInterfaceRepositoryConstructorFunctionId(driver string, entity model.Entity) (string, error)

	OutputScaffoldInterfaceRepositoryRegistryStructId(driver string, entity model.Entity) (string, error)
	OutputScaffoldInterfaceRepositoryRegistryConstructorFunctionId(driver string, entity model.Entity) (string, error)
	OutputScaffoldInterfaceRepositoryRegistryLocalConstructorFunctionId(driver string, entity model.Entity) (string, error)
	
	OutputScaffoldInterfacePresenterPackageName() (string, error)
	OutputInterfacePresenterStructId(driver string, entity model.Entity) (string, error)
	OutputScaffoldInterfacePresenterStructId(driver string, entity model.Entity) (string, error)
	OutputScaffoldInterfacePresenterInterfaceId(driver string, entity model.Entity) (string, error)
	OutputScaffoldInterfacePresenterConstructorFunctionId(driver string, entity model.Entity) (string, error)

	OutputScaffoldInterfacePresenterRegistryStructId(driver string, entity model.Entity) (string, error)
	OutputScaffoldInterfacePresenterRegistryConstructorFunctionId(driver string, entity model.Entity) (string, error)
	OutputScaffoldInterfacePresenterRegistryLocalConstructorFunctionId(driver string, entity model.Entity) (string, error)

	OutputScaffoldInterfaceControllerPackageName() (string, error)
	OutputInterfaceControllerStructId(driver string, entity model.Entity) (string, error)
	OutputScaffoldInterfaceControllerStructId(driver string, entity model.Entity) (string, error)
	OutputInterfaceControllerInterfaceId(driver string, entity model.Entity) (string, error)
	OutputScaffoldInterfaceControllerInterfaceId(driver string, entity model.Entity) (string, error)
	OutputScaffoldInterfaceControllerConstructorFunctionId(driver string, entity model.Entity) (string, error)

	OutputScaffoldInterfaceControllerRegistryStructId(driver string, entity model.Entity) (string, error)
	OutputScaffoldInterfaceControllerRegistryConstructorFunctionId(driver string, entity model.Entity) (string, error)
	OutputScaffoldInterfaceControllerRegistryLocalConstructorFunctionId(driver string, entity model.Entity) (string, error)

}

func New(config *config.Config) Formatter {
	return &formatter{
		config : config,
	}
}

// Output

func (formatter *formatter) OutputDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath}, formatter.config.PathSeparator), nil
}

// Registry
func (formatter *formatter) OutputRegistryPackageName() (string, error) {
	return strcase.ToSnake(formatter.config.Registry.Name), nil
}

func (formatter *formatter) OutputRegistryDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath, formatter.config.Registry.Name}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputRegistryFile() (string, error) {
	
	path, err  := formatter.OutputRegistryDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{formatter.config.Registry.Name}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, file}, formatter.config.PathSeparator), nil

}
func (formatter *formatter) OutputRegistryEntityFile(entity model.Entity) (string, error) {
	
	path, err  := formatter.OutputRegistryDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{entity.Name, formatter.config.Registry.Name}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, file}, formatter.config.PathSeparator), nil

}

func (formatter *formatter) OutputRegistryEntityPresenterConstructorFunctionId(entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{"new", entity.Name, formatter.config.Layers.Usecase.Presenter.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputRegistryEntityRepositoryConstructorFunctionId(entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{"new", entity.Name, formatter.config.Layers.Usecase.Repository.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputRegistryEntityControllerConstructorFunctionId(driver string, entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{"new", driver, entity.Name, formatter.config.Layers.Interface.Controller.Name}, formatter.config.StringSeparator)), nil
}

// Domain

func (formatter *formatter) OutputScaffoldDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldDomainDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath, formatter.config.Layers.Domain.Name}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath, formatter.config.Layers.Domain.Name, formatter.config.Layers.Domain.Entity.Name}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityDirectoryImportPath() (string, error) {
	path, err  := formatter.OutputScaffoldDomainEntityDirectory()
	if err != nil {
		return "", nil
	}
	return strings.TrimLeft(strings.Replace(path, strings.Join([]string{os.Getenv("GOPATH"), formatter.config.WorkspaceSourceDirName}, formatter.config.PathSeparator), "", 1), formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputDomainEntityFile(entity model.Entity) (string, error) {
	path, err  := formatter.OutputScaffoldDomainEntityDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{entity.Name}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, file}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityFile(entity model.Entity) (string, error) {
	path, err  := formatter.OutputScaffoldDomainEntityDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{entity.Name, formatter.config.Scaffold.Name}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, file}, formatter.config.PathSeparator), nil
}

// Usecase

func (formatter *formatter) OutputScaffoldUsecaseDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath, formatter.config.Layers.Usecase.Name}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldUsecaseInteractorDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath, formatter.config.Layers.Usecase.Name, formatter.config.Layers.Usecase.Interactor.Name}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldUsecaseInteractorDirectoryImportPath() (string, error) {
	path, err  := formatter.OutputScaffoldUsecaseInteractorDirectory()
	if err != nil {
		return "", nil
	}
	return strings.TrimLeft(strings.Replace(path, strings.Join([]string{os.Getenv("GOPATH"), formatter.config.WorkspaceSourceDirName}, formatter.config.PathSeparator), "", 1), formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputUsecaseInteractorFile(entity model.Entity) (string, error) {
	path, err  := formatter.OutputScaffoldUsecaseInteractorDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{entity.Name, formatter.config.Layers.Usecase.Interactor.Name}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, file}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldUsecaseInteractorFile(entity model.Entity) (string, error) {
	path, err  := formatter.OutputScaffoldUsecaseInteractorDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{entity.Name, formatter.config.Layers.Usecase.Interactor.Name, formatter.config.Scaffold.Name}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, file}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldUsecaseRepositoryDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath, formatter.config.Layers.Usecase.Name, formatter.config.Layers.Usecase.Repository.Name}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldUsecaseRepositoryDirectoryImportPath() (string, error) {
	path, err  := formatter.OutputScaffoldUsecaseRepositoryDirectory()
	if err != nil {
		return "", nil
	}
	return strings.TrimLeft(strings.Replace(path, strings.Join([]string{os.Getenv("GOPATH"), formatter.config.WorkspaceSourceDirName}, formatter.config.PathSeparator), "", 1), formatter.config.PathSeparator), nil
}


func (formatter *formatter) OutputUsecaseRepositoryFile(entity model.Entity) (string, error) {
	path, err  := formatter.OutputScaffoldUsecaseRepositoryDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{entity.Name, formatter.config.Layers.Usecase.Repository.Name}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, file}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldUsecaseRepositoryFile(entity model.Entity) (string, error) {
	path, err  := formatter.OutputScaffoldUsecaseRepositoryDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{entity.Name, formatter.config.Layers.Usecase.Repository.Name, formatter.config.Scaffold.Name}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, file}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldUsecasePresenterDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath, formatter.config.Layers.Usecase.Name, formatter.config.Layers.Usecase.Presenter.Name}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldUsecasePresenterDirectoryImportPath() (string, error) {
	path, err  := formatter.OutputScaffoldUsecasePresenterDirectory()
	if err != nil {
		return "", nil
	}
	return strings.TrimLeft(strings.Replace(path, strings.Join([]string{os.Getenv("GOPATH"), formatter.config.WorkspaceSourceDirName}, formatter.config.PathSeparator), "", 1), formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputUsecasePresenterFile(entity model.Entity) (string, error) {
	path, err  := formatter.OutputScaffoldUsecasePresenterDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{entity.Name, formatter.config.Layers.Usecase.Presenter.Name}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, file}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldUsecasePresenterFile(entity model.Entity) (string, error) {
	path, err  := formatter.OutputScaffoldUsecasePresenterDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{entity.Name, formatter.config.Layers.Usecase.Presenter.Name, formatter.config.Scaffold.Name}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, file}, formatter.config.PathSeparator), nil
}

// Interface

func (formatter *formatter) OutputScaffoldInterfaceDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath, formatter.config.Layers.Interface.Name}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldInterfaceControllerDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath, formatter.config.Layers.Interface.Name, formatter.config.Layers.Interface.Controller.Name}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputInterfaceControllerDirectory(driver string, entity model.Entity) (string, error) {
	path, err  := formatter.OutputScaffoldInterfaceControllerDirectory()
	if err != nil {
		return "", nil
	}
	return strings.Join([]string{path, entity.Name, driver}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputInterfaceControllerDirectoryImportPath(driver string, entity model.Entity) (string, error) {
	path, err  := formatter.OutputInterfaceControllerDirectory(driver, entity)
	if err != nil {
		return "", nil
	}
	return strings.TrimLeft(strings.Replace(path, strings.Join([]string{os.Getenv("GOPATH"), formatter.config.WorkspaceSourceDirName}, formatter.config.PathSeparator), "", 1), formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputInterfaceControllerFile(driver string, entity model.Entity) (string, error) {
	path, err  := formatter.OutputScaffoldInterfaceControllerDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{formatter.config.Layers.Interface.Controller.Name}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, entity.Name, driver, file}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldInterfaceControllerFile(driver string, entity model.Entity) (string, error) {
	path, err  := formatter.OutputScaffoldInterfaceControllerDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{formatter.config.Scaffold.Name}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, entity.Name, driver, file}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputInterfaceControllerEntityFile(driver string, entity model.Entity) (string, error) {
	
	path, err  := formatter.OutputScaffoldDomainEntityDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{entity.Name, driver}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, file}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldInterfaceRepositoryDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath, formatter.config.Layers.Interface.Name, formatter.config.Layers.Interface.Repository.Name}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputInterfaceRepositoryDirectory(driver string, entity model.Entity) (string, error) {
	path, err  := formatter.OutputScaffoldInterfaceRepositoryDirectory()
	if err != nil {
		return "", nil
	}
	return strings.Join([]string{path, entity.Name, driver}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputInterfaceRepositoryDirectoryImportPath(driver string, entity model.Entity) (string, error) {
	path, err  := formatter.OutputInterfaceRepositoryDirectory(driver, entity)
	if err != nil {
		return "", nil
	}
	return strings.TrimLeft(strings.Replace(path, strings.Join([]string{os.Getenv("GOPATH"), formatter.config.WorkspaceSourceDirName}, formatter.config.PathSeparator), "", 1), formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputInterfacePresenterDirectory(driver string, entity model.Entity) (string, error) {
	path, err  := formatter.OutputScaffoldInterfacePresenterDirectory()
	if err != nil {
		return "", nil
	}
	return strings.Join([]string{path, entity.Name, driver}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputInterfacePresenterDirectoryImportPath(driver string, entity model.Entity) (string, error) {
	path, err  := formatter.OutputInterfacePresenterDirectory(driver, entity)
	if err != nil {
		return "", nil
	}
	return strings.TrimLeft(strings.Replace(path, strings.Join([]string{os.Getenv("GOPATH"), formatter.config.WorkspaceSourceDirName}, formatter.config.PathSeparator), "", 1), formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputInterfaceRepositoryFile(driver string, entity model.Entity) (string, error) {
	path, err  := formatter.OutputScaffoldInterfaceRepositoryDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{formatter.config.Layers.Interface.Repository.Name}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, entity.Name, driver, file}, formatter.config.PathSeparator), nil
}
func (formatter *formatter) OutputScaffoldInterfaceRepositoryFile(driver string, entity model.Entity) (string, error) {
	path, err  := formatter.OutputScaffoldInterfaceRepositoryDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{formatter.config.Scaffold.Name}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, entity.Name, driver, file}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputInterfaceRepositoryEntityFile(driver string, entity model.Entity) (string, error) {
	
	path, err  := formatter.OutputScaffoldDomainEntityDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{entity.Name, driver}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, file}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputInterfaceRepositoryRegistryFile(driver string, entity model.Entity) (string, error) {
	
	path, err  := formatter.OutputRegistryDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{entity.Name, driver, formatter.config.Layers.Usecase.Repository.Name, formatter.config.Registry.Name}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, file}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputScaffoldInterfacePresenterDirectory() (string, error) {
	return strings.Join([]string{formatter.config.AbsOutputPath, formatter.config.Layers.Interface.Name, formatter.config.Layers.Interface.Presenter.Name}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputInterfacePresenterFile(driver string, entity model.Entity) (string, error) {
	path, err  := formatter.OutputScaffoldInterfacePresenterDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{formatter.config.Layers.Interface.Presenter.Name}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strcase.ToSnake(strings.Join([]string{path, entity.Name, driver, file}, formatter.config.PathSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfacePresenterFile(driver string, entity model.Entity) (string, error) {
	path, err  := formatter.OutputScaffoldInterfacePresenterDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{formatter.config.Scaffold.Name}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strcase.ToSnake(strings.Join([]string{path, entity.Name, driver,  file}, formatter.config.PathSeparator)), nil
}

func (formatter *formatter) OutputInterfacePresenterRegistryFile(driver string, entity model.Entity) (string, error) {
	
	path, err  := formatter.OutputRegistryDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{entity.Name, driver, formatter.config.Layers.Usecase.Presenter.Name, formatter.config.Registry.Name}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, file}, formatter.config.PathSeparator), nil
}

func (formatter *formatter) OutputInterfacePresenterEntityFile(driver string, entity model.Entity) (string, error) {
	
	path, err  := formatter.OutputScaffoldDomainEntityDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{entity.Name, driver}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, file}, formatter.config.PathSeparator), nil
}



func (formatter *formatter) OutputInterfaceControllerRegistryFile(driver string, entity model.Entity) (string, error) {
	
	path, err  := formatter.OutputRegistryDirectory()
	if err != nil {
		return "", nil
	}
	file := strings.Join([]string{strcase.ToSnake(strings.Join([]string{entity.Name, driver, formatter.config.Layers.Interface.Controller.Name, formatter.config.Registry.Name}, " ")), "go"}, formatter.config.StringSeparator)
	
	return strings.Join([]string{path, file}, formatter.config.PathSeparator), nil
}

// Domain Entity

func (formatter *formatter) OutputScaffoldDomainEntityPackageName() (string, error) {
	return strcase.ToSnake(formatter.config.Layers.Domain.Entity.Name), nil
}

func (formatter *formatter) OutputDomainEntityStructId(entity model.Entity) (string, error) {

	var resp string
	if entity.Interface {
		resp = strcase.ToLowerCamel(strings.Join([]string{entity.Name}, formatter.config.StringSeparator))
	} else {
		resp =strcase.ToCamel(strings.Join([]string{entity.Name}, formatter.config.StringSeparator))
	}
	return resp, nil

}

func (formatter *formatter) OutputDomainEntityLocalStructId(entity model.Entity) (string, error) {

	return strcase.ToLowerCamel(strings.Join([]string{entity.Name, "local"}, formatter.config.StringSeparator)), nil

}

func (formatter *formatter) OutputDomainEntitySliceStructId(entity model.Entity) (string, error) {
	pluralize := pluralize.NewClient()
	return strcase.ToLowerCamel(strings.Join([]string{pluralize.Plural(entity.Name)}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputDomainEntityInterfaceId(entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{entity.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputDomainEntitySliceInterfaceId(entity model.Entity) (string, error) {
	pluralize := pluralize.NewClient()
	return strcase.ToCamel(strings.Join([]string{pluralize.Plural(entity.Name)}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityStructId(entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{entity.Name, "struct"}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntitySliceStructId(entity model.Entity) (string, error) {
	pluralize := pluralize.NewClient()
	return strcase.ToLowerCamel(strings.Join([]string{pluralize.Plural(entity.Name), "struct"}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityInterfaceId(entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{entity.Name, "interface"}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntitySliceInterfaceId(entity model.Entity) (string, error) {
	pluralize := pluralize.NewClient()
	return strcase.ToLowerCamel(strings.Join([]string{pluralize.Plural(entity.Name), "interface"}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityGetterId(field model.Field) (string, error) {
	return strcase.ToCamel(field.Name), nil
}

func (formatter *formatter) OutputScaffoldDomainEntitySetterId(field model.Field) (string, error) {
	return strcase.ToCamel(strings.Join([]string{formatter.config.Layers.Domain.SetterVerb, field.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityCallbackId(callback model.Callback, method model.Method) (string, error) {
	return strcase.ToCamel(strings.Join([]string{callback.Type, method.Type}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntitySetAllSetterId() (string, error) {
	return strcase.ToCamel(strings.Join([]string{formatter.config.Layers.Domain.SetterVerb, "all"}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityToPrimaryId() (string, error) {
	return strcase.ToCamel(strings.Join([]string{formatter.config.Layers.Domain.ToPrimaryName}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityLenId() (string, error) {
	return strcase.ToCamel(strings.Join([]string{formatter.config.Layers.Domain.LenName}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityAppendId() (string, error) {
	return strcase.ToCamel(strings.Join([]string{formatter.config.Layers.Domain.AppendName}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityElementsId() (string, error) {
	return strcase.ToCamel(strings.Join([]string{formatter.config.Layers.Domain.ElementsName}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityFieldId(field model.Field) (string, error) {
	var resp string
	if field.Private {
		resp = strcase.ToLowerCamel(field.Name)
	} else {
		resp = strcase.ToCamel(field.Name)
	}
	return resp, nil
}

func (formatter *formatter) OutputDomainEntityInterfaceConstructorFunctionId(entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{"new", entity.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityInterfaceConstructorFunctionId(entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{"new", entity.Name, formatter.config.Scaffold.Name, "struct"}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputDomainEntitySliceInterfaceConstructorFunctionId(entity model.Entity) (string, error) {
	pluralize := pluralize.NewClient()
	return strcase.ToCamel(strings.Join([]string{"new", pluralize.Plural(entity.Name)}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntitySliceInterfaceConstructorFunctionId(entity model.Entity) (string, error) {
	pluralize := pluralize.NewClient()
	return strcase.ToCamel(strings.Join([]string{"new", pluralize.Plural(entity.Name), formatter.config.Scaffold.Name, "struct"}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldDomainEntityJSONTagId(field model.Field) (string, error) {
	return strcase.ToSnake(field.Name), nil
}


// Usecase

func (formatter *formatter) OutputScaffoldUsecaseRepositoryPackageName() (string, error) {
	return strcase.ToSnake(formatter.config.Layers.Usecase.Repository.Name), nil
}

func (formatter *formatter) OutputUsecaseRepositoryInterfaceId(entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{entity.Name, formatter.config.Layers.Usecase.Repository.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldUsecaseRepositoryInterfaceId(entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{entity.Name, formatter.config.Layers.Usecase.Repository.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldUsecaseRepositoryInterfaceMethodId(method model.Method) (string, error) {
	return strcase.ToCamel(method.Name), nil
}

func (formatter *formatter) OutputScaffoldUsecasePresenterPackageName() (string, error) {
	return strcase.ToSnake(formatter.config.Layers.Usecase.Presenter.Name), nil
}

func (formatter *formatter) OutputUsecasePresenterInterfaceId(entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{entity.Name, formatter.config.Layers.Usecase.Presenter.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldUsecasePresenterInterfaceId(entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{entity.Name, formatter.config.Layers.Usecase.Presenter.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldUsecasePresenterInterfaceMethodId(method model.Method) (string, error) {
	return strcase.ToCamel(method.Name), nil
}

func (formatter *formatter) OutputScaffoldUsecaseInteractorPackageName() (string, error) {
	return strcase.ToSnake(formatter.config.Layers.Usecase.Interactor.Name), nil
}

func (formatter *formatter) OutputUsecaseInteractorStructId(entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{entity.Name, formatter.config.Layers.Usecase.Interactor.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldUsecaseInteractorStructId(entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{entity.Name, formatter.config.Layers.Usecase.Interactor.Name, "struct"}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputUsecaseInteractorInterfaceId(entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{entity.Name, formatter.config.Layers.Usecase.Interactor.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldUsecaseInteractorInterfaceId(entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{entity.Name, formatter.config.Layers.Usecase.Interactor.Name, "interface"}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldUsecaseInteractorInterfaceMethodId(method model.Method) (string, error) {
	return strcase.ToCamel(method.Name), nil
}

func (formatter *formatter) OutputUsecaseInteractorInterfaceConstructorFunctionId(entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{"new", entity.Name, formatter.config.Layers.Usecase.Interactor.Name}, formatter.config.StringSeparator)), nil
}

// Interface

func (formatter *formatter) OutputScaffoldInterfaceRepositoryPackageName() (string, error) {
	return strcase.ToSnake(formatter.config.Layers.Interface.Repository.Name), nil
}

func (formatter *formatter) OutputInterfaceRepositoryStructId(driver string, entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{driver, entity.Name, formatter.config.Layers.Interface.Repository.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfaceRepositoryStructId(driver string, entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{driver, entity.Name, formatter.config.Layers.Interface.Repository.Name, "struct"}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfaceRepositoryRegistryStructId(driver string, entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{entity.Name, driver, formatter.config.Layers.Interface.Repository.Name, formatter.config.Registry.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfaceRepositoryRegistryConstructorFunctionId(driver string, entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{"register", entity.Name, driver, formatter.config.Layers.Interface.Repository.Name, formatter.config.Registry.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfaceRepositoryRegistryLocalConstructorFunctionId(driver string, entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{"new", entity.Name, driver, formatter.config.Layers.Interface.Repository.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfaceRepositoryConstructorFunctionId(driver string, entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{"new"}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfacePresenterPackageName() (string, error) {
	return strcase.ToSnake(formatter.config.Layers.Interface.Presenter.Name), nil
}

func (formatter *formatter) OutputInterfacePresenterStructId(driver string, entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{driver, entity.Name, formatter.config.Layers.Interface.Presenter.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfacePresenterStructId(driver string, entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{driver, entity.Name,formatter.config.Layers.Interface.Presenter.Name, "struct"}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfacePresenterInterfaceId(driver string, entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{driver, entity.Name, formatter.config.Layers.Interface.Presenter.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfacePresenterConstructorFunctionId(driver string, entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{"new"}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfacePresenterRegistryStructId(driver string, entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{entity.Name, driver, formatter.config.Layers.Interface.Presenter.Name, formatter.config.Registry.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfacePresenterRegistryConstructorFunctionId(driver string, entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{"register", entity.Name, driver, formatter.config.Layers.Interface.Presenter.Name, formatter.config.Registry.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfacePresenterRegistryLocalConstructorFunctionId(driver string, entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{"new", entity.Name, driver, formatter.config.Layers.Interface.Presenter.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfaceControllerPackageName() (string, error) {
	return strcase.ToSnake(formatter.config.Layers.Interface.Controller.Name), nil
}

func (formatter *formatter) OutputInterfaceControllerStructId(driver string, entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{driver, entity.Name, formatter.config.Layers.Interface.Controller.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfaceControllerStructId(driver string, entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{driver, entity.Name, formatter.config.Layers.Interface.Controller.Name, "struct"}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputInterfaceControllerInterfaceId(driver string, entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{driver, entity.Name, formatter.config.Layers.Interface.Controller.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfaceControllerInterfaceId(driver string, entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{driver, entity.Name, formatter.config.Layers.Interface.Controller.Name, "interface"}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfaceControllerConstructorFunctionId(driver string, entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{"new"}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfaceControllerRegistryStructId(driver string, entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{entity.Name, driver, formatter.config.Layers.Interface.Controller.Name, formatter.config.Registry.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfaceControllerRegistryConstructorFunctionId(driver string, entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{"register", entity.Name, driver, formatter.config.Layers.Interface.Controller.Name, formatter.config.Registry.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldInterfaceControllerRegistryLocalConstructorFunctionId(driver string, entity model.Entity) (string, error) {
	return strcase.ToCamel(strings.Join([]string{"new", entity.Name, driver, formatter.config.Layers.Interface.Controller.Name}, formatter.config.StringSeparator)), nil
}

func (formatter *formatter) OutputScaffoldUsecaseInteractorRegistryLocalConstructorFunctionId(entity model.Entity) (string, error) {
	return strcase.ToLowerCamel(strings.Join([]string{"new", entity.Name, formatter.config.Layers.Usecase.Interactor.Name}, formatter.config.StringSeparator)), nil
}