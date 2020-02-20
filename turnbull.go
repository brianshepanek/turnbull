package turnbull

import(
	"os"
	"strings"
	"io/ioutil"
	"github.com/BurntSushi/toml"
	// "fmt"
)

const entityName = "entity"
const entitiesDirName = "entities"
const pathSeparator = "/"
const stringSeparator = "."
const setterVerb = "set"

type config struct {
	entities []entity
}

type entity struct {
	Name string `toml:"name"`
	Fields []field `toml:"fields"`
	JSON bool `toml:"json"`
}

type field struct{
	Op string `toml:"op"`
	Name string `toml:"name"`
	Package string `toml:"package"`
	Type string `toml:"type"`
	Slice bool `toml:"slice"`
}

func Build(configPath, outputPath string) (error) {

	// Load Config
	configPointer, err := loadConfig(configPath)
	if err != nil {
		return err
	}

	// Create Output Dir
	if _, err := os.Stat(strings.Join([]string{outputPath}, pathSeparator)); os.IsNotExist(err) {
		err = os.Mkdir(strings.Join([]string{outputPath}, pathSeparator), os.ModePerm)
	}

	// Create Output Entities Dir
	if _, err := os.Stat(strings.Join([]string{outputPath, entitiesDirName}, pathSeparator)); os.IsNotExist(err) {
		err = os.Mkdir(strings.Join([]string{outputPath, entitiesDirName}, pathSeparator), os.ModePerm)
	}

	// Entity Files
	if configPointer != nil {

		// Deref
		config := *configPointer

		// Build Entity File
		for _, entity := range config.entities {
			
			// Parse Entity Config
			entityFilePointer, err := buildEntityFile(entity)
			if err != nil {
				return err
			}

			// Save
			err = entityFilePointer.Save(strings.Join([]string{outputPath, entitiesDirName, strings.Join([]string{entity.Name, entityName, "go"}, stringSeparator)}, pathSeparator))
			if err != nil {
				return err
			}
			
		}

	}
	

	// Return
	return nil

}

func loadConfig(path string) (*config, error) {

	// Vars
	var config config

	// Load Entity Files
	entitiesPath := strings.Join([]string{path, entitiesDirName}, pathSeparator)
	entityFilesInfo, err := ioutil.ReadDir(entitiesPath)
	if err != nil {
		return nil, err
	}

	// Parse Entity Files
	for _, entityFileInfo := range entityFilesInfo {

		// Open
		entityFile, err := os.Open(strings.Join([]string{entitiesPath, entityFileInfo.Name()}, pathSeparator))
		if err != nil {
			return nil, err
		}
		defer entityFile.Close()

		// Read
		entityFileBytes, err := ioutil.ReadAll(entityFile)
		if err != nil {
			return nil, err
		}

		// Parse
		entityPointer, err := parseEntityConfig(entityFileBytes)
		if err != nil {
			return nil, err
		}

		// Append
		if entityPointer != nil {
			config.entities = append(config.entities, *entityPointer)
		}

	}

	return &config, nil
}

func parseEntityConfig(entityConfigBytes []byte) (*entity, error){

	// Vars
	var entity entity

	// Unmarshal
	err := toml.Unmarshal(entityConfigBytes, &entity)
	if err != nil {
		return nil, err
	}

	// Return
	return &entity, nil
}

