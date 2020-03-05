package turnbull

type configEntity struct {
	EntityName string `toml:"entity_name"`
	entitiesDirName string `toml:"entities_dir_name"`
	domainLayerName string
	interfaceLayerName string
	usecaseLayerName string
	interactorName string
	presenterName string
	repositoryName string
	controllerName string
	scaffoldName string
	scaffoldDirName string
	pathSeparator string
	stringSeparator string
	setterVerb string
	entities []entity
	methods []method
	absOutputPath string
	workspaceSourceDirName string
}

type entity struct {
	Name string `toml:"name"`
	Fields []field `toml:"fields"`
	JSON bool `toml:"json"`
	Methods []entityMethod `toml:"methods"`
}

type field struct{
	Primary bool `toml:"primary"`
	Op string `toml:"op"`
	Name string `toml:"name"`
	Package string `toml:"package"`
	Type string `toml:"type"`
	Slice bool `toml:"slice"`
}

type entityMethod struct {
	Name string `toml:"name"`
	Type string `toml:"type"`
}

type method struct {
	Type string `toml:"type"`
	Repository repositoryMethod `toml:"repository"`
	Presenter presenterMethod `toml:"presenter"`
}

type repository struct {
	Methods []repositoryMethod `toml:"methods"`
}

type repositoryMethod struct {
	Name string `toml:"name"`
	Arguments []field `toml:"arguments"`
	ReturnValues []field `toml:"return_values"`
}

type presenter struct {
	Methods []presenterMethod `toml:"methods"`
}

type presenterMethod struct {
	Name string `toml:"name"`
	Arguments []field `toml:"arguments"`
	ReturnValues []field `toml:"return_values"`
}

var config = &configEntity{
	scaffoldName : "scaffold",
	scaffoldDirName : "scaffold",
	EntityName : "model",
	entitiesDirName : "model",
	domainLayerName : "domain",
	interfaceLayerName : "interface",
	usecaseLayerName : "usecase",
	interactorName : "interactor",
	presenterName : "presenter",
	repositoryName : "repository",
	controllerName : "controller",
	pathSeparator : "/",
	stringSeparator : ".",
	setterVerb : "set",
	workspaceSourceDirName : "src",
}