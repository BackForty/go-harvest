package harvest

import (
	"encoding/json"
	"time"
)

type HarvestDate struct {
	time.Time
}

func (harvestDate *HarvestDate) UnmarshalJSON(input []byte) (err error) {
	dateString := string(input)

	if len(dateString) == 12 {
		harvestDate.Time, err = time.Parse("\"2006-01-02\"", dateString)
	} else if dateString != "null" {
		err = json.Unmarshal(input, &harvestDate.Time)
	}

	return
}
