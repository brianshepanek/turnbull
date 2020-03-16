package generator


import(
	"github.com/dave/jennifer/jen"
	"github.com/brianshepanek/turnbull/domain/model"
)

type RepositoryGenerator interface{
	ScaffoldFile(entity model.Entity) (*jen.File, error)
}

type PresenterGenerator interface{
	ScaffoldFile(entity model.Entity) (*jen.File, error)
}

type ControllerGenerator interface{
	ScaffoldFile(entity model.Entity) (*jen.File, error)
}