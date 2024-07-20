package statsmodel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type PropertyMap map[string]interface{}

// Структура
type Stats struct {
	PlayerID     int         `json:"playerid"`
	PlayerStats  PropertyMap `json:"PlayerStats"`
	HeroStats    PropertyMap `json:"HeroStats"`
	MatchesStats PropertyMap `json:"MatchesStats"`
	WordCloud    PropertyMap `json:"WordCloud"`
}

func (p PropertyMap) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

func (p *PropertyMap) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .[]byte failed")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*p, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("type assertion (map) failed")
	}

	return nil
}
