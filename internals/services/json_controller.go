package services

import (
	"fr/diginamic/go-cap/internals/types"
	"fr/diginamic/go-cap/internals/utils"
)

func WriteJsonController(ctrlName string) {

    checkDirectoryExists[JSON_CONTROLLER]()

	controller := types.BasicJsonController{
		NameSpace: utils.ConvertPathToNameSpace(CONF.JsonController),
		ClassName: ctrlName,
        Route: utils.ConvertNameToRoute(ctrlName),
        RouteName: utils.ConvertNameToRouteName(ctrlName),
	}

	tmplBytes := getTmplBytes(ctrlName, "JsonController.php", controller)
	path := SRC + CONF.JsonController + ctrlName + "Controller.php"
	createFile(tmplBytes, path)

}

func WriteJsonControllerWithServ(ctrlName string) {

    checkDirectoryExists[JSON_CONTROLLER]()
    checkDirectoryExists[SERVICE]()
    checkDirectoryExists[ENTITY]()
    checkDirectoryExists[DTO]()
    checkDependencies([]string{
        CONF.Service,
    },ctrlName)

    controller := types.AdvancedJsonController{
		NameSpace: utils.ConvertPathToNameSpace(CONF.JsonController),
		ClassName: ctrlName,
        Route: utils.ConvertNameToRoute(ctrlName),
        RouteName: utils.ConvertNameToRouteName(ctrlName),
        ServiceNS: utils.ConvertPathToNameSpace(CONF.Service),
        RepositoryNS: utils.ConvertPathToNameSpace(CONF.Repository),
    }
	tmplBytes := getTmplBytes(ctrlName, "JsonControllerService.php", controller)
	path := SRC + CONF.JsonController + ctrlName + "Controller.php"
	createFile(tmplBytes, path)
}
