package generator

import(
	"io"
	"github.com/brianshepanek/turnbull/domain/model"
)

type UsecaseGenerator interface {
	ScaffoldEntity(entity model.Entity, writer io.Writer) (error)
	ScaffoldUsecaseRepository(entity model.Entity, writer io.Writer) (error)
	ScaffoldUsecasePresenter(entity model.Entity, writer io.Writer) (error)
	ScaffoldUsecaseInteractor(entity model.Entity, writer io.Writer) (error)
}

type InterfaceGenerator interface {
	ScaffoldInterfaceRepository(entity model.Entity, writer io.Writer) (error)
}

