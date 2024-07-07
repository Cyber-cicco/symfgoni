package services

import (
	"fr/diginamic/go-cap/internals/types"
	"fr/diginamic/go-cap/internals/utils"
	"os"
)

func WriteHtmlController(name string) {

    checkDirectoryExists[HTML_CONTROLLER]()

    snakeName := utils.SnakeCase(name)

	controller := types.HtmlController{
		NameSpace: utils.ConvertPathToNameSpace(CONF.JsonController),
		ClassName: name,
        Route: utils.ConvertNameToRoute(name),
        RouteName: utils.ConvertNameToRouteName(name),
        TemplateName: snakeName,
	}

	tmplBytes := getTmplBytes(name, "HtmlController.php", controller)
	path := SRC + CONF.HtmlController + name + "Controller.php"
    templatePath := CONF.TemplateDir + CONF.Page + snakeName
	createFile(tmplBytes, path)
    os.Mkdir(templatePath, 0777)
    createFile(nil, templatePath + "/index.html.twig")
}
