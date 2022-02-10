package internal

import (
	"regexp"
	"strconv"

	"github.com/andersfylling/snowflake/v5"
	log "github.com/sirupsen/logrus"
)

func ConvertStringtoSnowflake(userIDStr string) snowflake.Snowflake {
	rx := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	match := rx.FindAllString(userIDStr, -1)
	var userID snowflake.Snowflake
	for _, element := range match {
		log.Info(element)
		number, _ := strconv.ParseUint(element, 10, 64)
		userID = snowflake.NewSnowflake(number)
	}

	return userID
}
