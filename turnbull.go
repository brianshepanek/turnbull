package turnbull

import(
	"os"
	"strings"
	"io/ioutil"
	"github.com/BurntSushi/toml"
	"path/filepath"
)

func Build(configPath, outputPath string) (error) {

	// Abs Output Path
	absOutputPath, err := filepath.Abs(outputPath)
	if err != nil {
		return err
	}
	config.absOutputPath = absOutputPath

	// Load Config
	configPointer, err := loadConfig(configPath)
	if err != nil {
		return err
	}

	// Load Config Methods
	err = loadConfigMethods()
	if err != nil {
		return err
	}

	// Build Directory Structure
	err = buildDirectoryStructure()
	if err != nil {
		return err
	}

	// Build Scaffold
	err = buildScaffold(configPointer)
	if err != nil {
		return err
	}

	// Return
	return nil

}

func loadConfig(path string) (*configEntity, error) {

	// Load Config
	_, err := toml.DecodeFile(path + "/config.toml", config)
	if err != nil {
		return nil, err
	}

	// Load Entity Files
	entitiesPath := strings.Join([]string{path, config.entitiesDirName}, config.pathSeparator)
	entityFilesInfo, err := ioutil.ReadDir(entitiesPath)
	if err != nil {
		return nil, err
	}

	// Parse Entity Files
	for _, entityFileInfo := range entityFilesInfo {

		// Open
		entityFile, err := os.Open(strings.Join([]string{entitiesPath, entityFileInfo.Name()}, config.pathSeparator))
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

	return config, nil
}

func loadConfigMethods() (error) {

	// Config Method Path
	configMethodsPath := "/go/src/github.com/brianshepanek/turnbull/config/methods"

	// Load Method Files
	methodFilesInfo, err := ioutil.ReadDir(configMethodsPath)
	if err != nil {
		return err
	}

	// Parse Method Files
	for _, methodFileInfo := range methodFilesInfo {

		// Open
		methodFile, err := os.Open(strings.Join([]string{configMethodsPath, methodFileInfo.Name()}, config.pathSeparator))
		if err != nil {
			return err
		}
		defer methodFile.Close()

		// Read
		methodFileBytes, err := ioutil.ReadAll(methodFile)
		if err != nil {
			return err
		}

		// Parse
		methodPointer, err := parseMethodConfig(methodFileBytes)
		if err != nil {
			return err
		}

		// Append
		if methodPointer != nil {
			config.methods = append(config.methods, *methodPointer)
		}

	}

	return nil
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

func parseMethodConfig(methodConfigBytes []byte) (*method, error){

	// Vars
	var method method

	// Unmarshal
	err := toml.Unmarshal(methodConfigBytes, &method)
	if err != nil {
		return nil, err
	}

	// Return
	return &method, nil
}


func buildDirectoryStructure()(error){

	// Create Output Dir
	if _, err := os.Stat(strings.Join([]string{config.absOutputPath}, config.pathSeparator)); os.IsNotExist(err) {
		err = os.Mkdir(strings.Join([]string{config.absOutputPath}, config.pathSeparator), os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Create Output Scaffold Dir
	if _, err := os.Stat(strings.Join([]string{config.absOutputPath, config.scaffoldDirName}, config.pathSeparator)); os.IsNotExist(err) {
		err = os.Mkdir(strings.Join([]string{config.absOutputPath, config.scaffoldDirName}, config.pathSeparator), os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Create Output Scaffold Domain Dir
	if _, err := os.Stat(strings.Join([]string{config.absOutputPath, config.scaffoldDirName, config.domainLayerName}, config.pathSeparator)); os.IsNotExist(err) {
		err = os.Mkdir(strings.Join([]string{config.absOutputPath, config.scaffoldDirName, config.domainLayerName}, config.pathSeparator), os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Create Output Scaffold Entities Dir
	if _, err := os.Stat(strings.Join([]string{config.absOutputPath, config.scaffoldDirName, config.domainLayerName, config.entitiesDirName}, config.pathSeparator)); os.IsNotExist(err) {
		err = os.Mkdir(strings.Join([]string{config.absOutputPath, config.scaffoldDirName, config.domainLayerName, config.entitiesDirName}, config.pathSeparator), os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Create Output Scaffold Useacse Dir
	if _, err := os.Stat(strings.Join([]string{config.absOutputPath, config.scaffoldDirName, config.usecaseLayerName}, config.pathSeparator)); os.IsNotExist(err) {
		err = os.Mkdir(strings.Join([]string{config.absOutputPath, config.scaffoldDirName, config.usecaseLayerName}, config.pathSeparator), os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Create Output Scaffold Useacse Repository Dir
	if _, err := os.Stat(strings.Join([]string{config.absOutputPath, config.scaffoldDirName, config.usecaseLayerName, config.repositoryName}, config.pathSeparator)); os.IsNotExist(err) {
		err = os.Mkdir(strings.Join([]string{config.absOutputPath, config.scaffoldDirName, config.usecaseLayerName, config.repositoryName}, config.pathSeparator), os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Create Output Scaffold Useacse Presenter Dir
	if _, err := os.Stat(strings.Join([]string{config.absOutputPath, config.scaffoldDirName, config.usecaseLayerName, config.presenterName}, config.pathSeparator)); os.IsNotExist(err) {
		err = os.Mkdir(strings.Join([]string{config.absOutputPath, config.scaffoldDirName, config.usecaseLayerName, config.presenterName}, config.pathSeparator), os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Create Output Scaffold Useacse Interactor Dir
	if _, err := os.Stat(strings.Join([]string{config.absOutputPath, config.scaffoldDirName, config.usecaseLayerName, config.interactorName}, config.pathSeparator)); os.IsNotExist(err) {
		err = os.Mkdir(strings.Join([]string{config.absOutputPath, config.scaffoldDirName, config.usecaseLayerName, config.interactorName}, config.pathSeparator), os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Create Output Scaffold Interface Dir
	if _, err := os.Stat(strings.Join([]string{config.absOutputPath, config.scaffoldDirName, config.interfaceLayerName}, config.pathSeparator)); os.IsNotExist(err) {
		err = os.Mkdir(strings.Join([]string{config.absOutputPath, config.scaffoldDirName, config.interfaceLayerName}, config.pathSeparator), os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Create Output Scaffold Useacse Repository Dir
	if _, err := os.Stat(strings.Join([]string{config.absOutputPath, config.scaffoldDirName, config.interfaceLayerName, config.repositoryName}, config.pathSeparator)); os.IsNotExist(err) {
		err = os.Mkdir(strings.Join([]string{config.absOutputPath, config.scaffoldDirName, config.interfaceLayerName, config.repositoryName}, config.pathSeparator), os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Create Output Scaffold Useacse Presenter Dir
	if _, err := os.Stat(strings.Join([]string{config.absOutputPath, config.scaffoldDirName, config.interfaceLayerName, config.presenterName}, config.pathSeparator)); os.IsNotExist(err) {
		err = os.Mkdir(strings.Join([]string{config.absOutputPath, config.scaffoldDirName, config.interfaceLayerName, config.presenterName}, config.pathSeparator), os.ModePerm)
		if err != nil {
			return err
		}
	}

	// Create Output Scaffold Useacse Controller Dir
	if _, err := os.Stat(strings.Join([]string{config.absOutputPath, config.scaffoldDirName, config.interfaceLayerName, config.controllerName}, config.pathSeparator)); os.IsNotExist(err) {
		err = os.Mkdir(strings.Join([]string{config.absOutputPath, config.scaffoldDirName, config.interfaceLayerName, config.controllerName}, config.pathSeparator), os.ModePerm)
		if err != nil {
			return err
		}
	}

	

	// Create Entities Dir
	// if _, err := os.Stat(strings.Join([]string{config.absOutputPath, config.entitiesDirName}, config.pathSeparator)); os.IsNotExist(err) {
	// 	err = os.Mkdir(strings.Join([]string{config.absOutputPath, config.entitiesDirName}, config.pathSeparator), os.ModePerm)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	

	return nil
}

func buildScaffold(configPointer *configEntity)(error){

	// Entity Files
	if configPointer != nil {

		// Deref
		config := *configPointer

		// Build Entity File
		for _, entity := range config.entities {

			/******
			Entity
			******/

			// Build Scaffold Entity File
			scaffoldEntityFilePointer, err := buildScaffoldEntityFile(entity)
			if err != nil {
				return err
			}

			// Save
			err = scaffoldEntityFilePointer.Save(scaffoldEntityFileName(config.absOutputPath, entity))
			if err != nil {
				return err
			}

			/******
			Usecase
			******/

			// Build Scaffold Usecase Interactor File
			scaffoldUsecaseInteractorFilePointer, err := buildScaffoldUsecaseInteractorFile(entity)
			if err != nil {
				return err
			}

			// Save
			err = scaffoldUsecaseInteractorFilePointer.Save(scaffoldUsecaseInteractorFileName(config.absOutputPath, entity))
			if err != nil {
				return err
			}

			// Build Scaffold Usecase Presenter File
			scaffoldUsecasePresenterFilePointer, err := buildScaffoldUsecasePresenterFile(entity)
			if err != nil {
				return err
			}

			// Save
			err = scaffoldUsecasePresenterFilePointer.Save(scaffoldUsecasePresenterFileName(config.absOutputPath, entity))
			if err != nil {
				return err
			}

			// Build Scaffold Usecase Repository File
			scaffoldUsecaseRepositoryFilePointer, err := buildScaffoldUsecaseRepositoryFile(entity)
			if err != nil {
				return err
			}

			// Save
			err = scaffoldUsecaseRepositoryFilePointer.Save(scaffoldUsecaseRepositoryFileName(config.absOutputPath, entity))
			if err != nil {
				return err
			}

			/******
			Interface
			******/

			// Build Scaffold Interface Respository File
			scaffoldInterfaceRespositoryFilePointer, err := buildScaffoldInterfaceRepositoryFile(entity)
			if err != nil {
				return err
			}

			// Save
			err = scaffoldInterfaceRespositoryFilePointer.Save(scaffoldInterfaceRepositoryFileName(config.absOutputPath, entity))
			if err != nil {
				return err
			}
			
			// // Build Scaffold Interface Presenter File
			// scaffoldInterfacePresenterFilePointer, err := buildScaffoldInterfacePresenterFile(entity)
			// if err != nil {
			// 	return err
			// }

			// // Save
			// err = scaffoldInterfacePresenterFilePointer.Save(scaffoldInterfacePresenterFileName(config.absOutputPath, entity))
			// if err != nil {
			// 	return err
			// }

		}

	}

	return nil
}