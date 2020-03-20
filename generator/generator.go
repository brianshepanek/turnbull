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
	ScaffoldUsecaseRepository(entity model.Entity, writer io.Writer) (error)
	ScaffoldUsecasePresenter(entity model.Entity, writer io.Writer) (error)
	ScaffoldUsecaseInteractor(entity model.Entity, writer io.Writer) (error)
}

type InterfaceGenerator interface {
	ScaffoldInterfaceRepository(entity model.Entity, writer io.Writer) (error)
	ScaffoldInterfacePresenter(entity model.Entity, writer io.Writer) (error)
	ScaffoldInterfaceController(entity model.Entity, writer io.Writer) (error)
}

