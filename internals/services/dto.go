package services

import (
	"fr/diginamic/go-cap/internals/config"
	"fr/diginamic/go-cap/internals/types"
	"fr/diginamic/go-cap/internals/utils"
)

func WriteDTO(args []string, fieldFlag string) {
	err := config.InitDto()

	if err != nil {
        askCustomDir("Dto", func(dir string) {
            config.CONFIG.Dto = dir
        })
	}

    for _, arg := range args {

        mainClass := arg + "Dto"

        dto := types.DtoFile {
            NameSpace: utils.ConvertPathToNameSpace(config.CONFIG.Dto),
            Imports: []string{},
            ClassName: mainClass,
        }

        tplBytes := getTmpleBytes(mainClass, "Dto.php", dto)
        path := config.CONFIG.SourceDirectory + config.CONFIG.Dto + mainClass + ".php"
        createFile(tplBytes, path)
    }
}
