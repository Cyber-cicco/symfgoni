package services

import (
	"fr/diginamic/go-cap/internals/config"
	"fr/diginamic/go-cap/internals/types"
	"fr/diginamic/go-cap/internals/utils"
)

func WriteService(serviceName string) {

    checkDirectoryExists[SERVICE]()

	service := types.File{
		NameSpace: utils.ConvertPathToNameSpace(config.CONFIG.Service),
		ClassName: serviceName,
	}

	tplBytes := getTmplBytes(serviceName, "Service.php", service)
	path := SRC + CONF.Service + serviceName + "Service.php"
	createFile(tplBytes, path)

}

func WriteServiceLinkedToEntity(entityName string) {

    checkDirectoryExists[SERVICE]()
    checkDirectoryExists[ENTITY]()
    checkDirectoryExists[MAPPER]()
    checkDirectoryExists[REPOSITORY]()
    checkDependencies([]string{
        CONF.Mapper,
        CONF.Repository,
    },entityName)

	service := types.ServiceFile{
		NameSpace: utils.ConvertPathToNameSpace(CONF.Service),
		ClassName: entityName,
        RepositoryNS: utils.ConvertPathToNameSpace(CONF.Repository),
        MapperNS: utils.ConvertPathToNameSpace(CONF.Mapper),
	}

	tplBytes := getTmplBytes(entityName, "ServiceRepo.php", service)
	path := SRC + CONF.Service + entityName + "Service.php"
	createFile(tplBytes, path)

}

