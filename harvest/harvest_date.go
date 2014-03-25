package harvest

import (
	"encoding/json"
	"regexp"
	"strings"
	"time"
)

type HarvestDate struct {
	time.Time
}

func (harvestDate *HarvestDate) UnmarshalJSON(input []byte) (err error) {
	dateString := string(input)
	dateString = strings.Trim(dateString, "\"")

	matchesShortDate, _ := regexp.MatchString("^[0-9][0-9][0-9][0-9]-[0-9][0-9]-[0-9][0-9]$", dateString)

	if matchesShortDate {
		harvestDate.Time, err = time.Parse("2006-01-02", dateString)
	} else if dateString != "null" {
		err = json.Unmarshal(input, &harvestDate.Time)
	}

	return
}
