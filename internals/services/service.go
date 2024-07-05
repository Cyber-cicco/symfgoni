package services

import (
	"fmt"
	"fr/diginamic/go-cap/internals/config"
	"fr/diginamic/go-cap/internals/types"
	"fr/diginamic/go-cap/internals/utils"
	"os"
)

func WriteService(serviceName string) {

    checkDirectoryExists[SERVICE]()

	service := types.File{
		NameSpace: utils.ConvertPathToNameSpace(config.CONFIG.Service),
		ClassName: serviceName,
	}

	tplBytes := getTmpleBytes(serviceName, "Service.php", service)
	path := SRC + CONF.Service + serviceName + "Service.php"
	createFile(tplBytes, path)

}

func WriteServiceLinkedToEntity(entityName string) {

    checkDirectoryExists[SERVICE]()
    checkDirectoryExists[ENTITY]()
    checkDirectoryExists[MAPPER]()
    checkDirectoryExists[REPOSITORY]()
    checkDependencies(entityName)

	service := types.ServiceFile{
		NameSpace: utils.ConvertPathToNameSpace(CONF.Service),
		ClassName: entityName,
        RepositoryNS: utils.ConvertPathToNameSpace(CONF.Repository),
        MapperNS: utils.ConvertPathToNameSpace(CONF.Mapper),
	}

	tplBytes := getTmpleBytes(entityName, "ServiceRepo.php", service)
	path := SRC + CONF.Service + entityName + "Service.php"
	createFile(tplBytes, path)

}

func checkDependencies(entityName string) {
    suffixes := map[string]string{
        CONF.Mapper : "Mapper.php",
        CONF.Entity : ".php",
        CONF.Repository : "Repository.php",
    }

    for k, v := range suffixes {
        _, err := os.Stat(SRC + k + entityName + v)
        
        if err != nil {
            fmt.Println(SRC + k + entityName + v)
            fmt.Printf("required dependency %s wasn't found. Make sure it exists before creating the service\n", entityName + v)
            os.Exit(1)
        }
    }
}
