package settings

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

var (
	AppSettings Settings
)

func ReadSettings() Settings {
	var option Settings
	viper.AddConfigPath(".")
	viper.SetConfigName("settings")
	viper.SetConfigType("json")

	if err:=viper.ReadInConfig(); err!=nil{
		log.Fatalln(err)
	}

	if err:=viper.Unmarshal(&option); err!=nil{
		log.Fatalln(err)
	}

	return option
}

func FullNameToInitials(fullName string) string {
	var initials string

	parts := strings.Fields(fullName)

	// when full name length is 2 (Example: "Aли Ализода"  -->  "Али А.")
	if len(parts) == 2 {
		initials = fmt.Sprintf("%s %s.", parts[0], parts[1][:2])

		// when full name length is 3 (Example: "Алиев Вали Ганиевич"  -->  "Вали А. Г.")
	} else {
		initials = fmt.Sprintf("%s %s. %s.", parts[1], parts[0][:2], parts[2][:2])
	}

	return initials
}
