package generator

import(
	"io"
	"github.com/brianshepanek/turnbull/domain/model"
)

type Generator interface {
	DomainGenerator
	UsecaseGenerator
	InterfaceGenerator

}

type DomainGenerator interface {
	Entity(entity model.Entity, writer io.Writer) (error)
	ScaffoldEntity(entity model.Entity, writer io.Writer) (error)
}

type UsecaseGenerator interface {
	UsecaseRepository(entity model.Entity, writer io.Writer) (error)
	ScaffoldUsecaseRepository(entity model.Entity, writer io.Writer) (error)
	UsecasePresenter(entity model.Entity, writer io.Writer) (error)
	ScaffoldUsecasePresenter(entity model.Entity, writer io.Writer) (error)
	UsecaseInteractor(entity model.Entity, writer io.Writer) (error)
	ScaffoldUsecaseInteractor(entity model.Entity, writer io.Writer) (error)
}

type InterfaceGenerator interface {
	InterfaceRepository(driver string, entity model.Entity, writer io.Writer) (error)
	ScaffoldInterfaceRepository(driver string, entity model.Entity, writer io.Writer) (error)
	InterfaceRepositoryEntity(driver string, entity model.Entity, writer io.Writer) (error)
	InterfacePresenter(driver string, entity model.Entity, writer io.Writer) (error)
	ScaffoldInterfacePresenter(driver string, entity model.Entity, writer io.Writer) (error)
	InterfaceAppController(driver string, entities []model.Entity, writer io.Writer) (error)
	InterfaceController(driver string, entity model.Entity, writer io.Writer) (error)
	ScaffoldInterfaceController(driver string, entity model.Entity, writer io.Writer) (error)
	InterfaceControllerEntity(driver string, entity model.Entity, writer io.Writer) (error)
}

