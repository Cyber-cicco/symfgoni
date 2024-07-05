package services

import (
	"fmt"
	"fr/diginamic/go-cap/internals/config"
	"fr/diginamic/go-cap/internals/types"
	"fr/diginamic/go-cap/internals/utils"
	"os"
	"strings"
)

func WriteDTO(arg string, fieldFlag string) {

    checkDirectoryExists[DTO]()

	fields := createFields(fieldFlag)


    mainClass := arg + "Dto"

    dto := types.DtoFile{
        NameSpace: utils.ConvertPathToNameSpace(config.CONFIG.Dto),
        ClassName: mainClass,
        Fields:    fields,
    }

    tplBytes := getTmpleBytes(mainClass, "Dto.php", dto)
    path := config.CONFIG.SourceDirectory + config.CONFIG.Dto + mainClass + ".php"
    createFile(tplBytes, path)
}

func createFields(fieldFlag string) []types.Field {

    if fieldFlag == "" {
        return []types.Field{}
    }

    fieldsStr := strings.Split(fieldFlag, " ")
    fields := make([]types.Field, len(fieldsStr))

    for i, field := range fieldsStr {
        strippedField := strings.Split(field, ":")

        if len(strippedField) != 2 {
            fmt.Println("Error was found in one of the fields. Use this syntax to declare a field : 'nom:type'")
            os.Exit(1)
        }

        newField := types.Field{
            Name: strippedField[0],
            Type: strippedField[1],
        }
        fields[i] = newField
    }
    return fields
}
