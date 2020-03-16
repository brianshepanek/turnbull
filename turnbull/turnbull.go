package turnbull

import(
	"os"
	"bytes"
	"github.com/brianshepanek/turnbull/domain/model"
	"github.com/brianshepanek/turnbull/formatter"
	"github.com/brianshepanek/turnbull/structure"
	"github.com/brianshepanek/turnbull/generator"
)

type turnbull struct{
	formatter formatter.Formatter
	structure structure.Structure
	generator generator.Generator
}

func New(formatter formatter.Formatter, structure structure.Structure, generator generator.Generator) *turnbull {
	return &turnbull{
		formatter : formatter,
		structure : structure,
		generator : generator,
	}
}

// Build
func (turnbull *turnbull) buildStructure() (error){

	var err error
	err = turnbull.structure.Build()
	if err != nil {
		return err
	}

	return nil
}

// Build Domain Entity
func (turnbull *turnbull) buildDomainEntity(entity model.Entity) (error){

	// Build
	buf := &bytes.Buffer{}
	err := turnbull.generator.ScaffoldEntity(entity, buf)
	if err != nil {
		return err
	}

	// File Name
	fileName, err := turnbull.formatter.OutputScaffoldDomainEntityFile(entity)
	if err != nil {
		return err
	}

	// File
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write
	_, err = file.WriteString(buf.String())
	if err != nil {
		return err
	}

	return nil
}

