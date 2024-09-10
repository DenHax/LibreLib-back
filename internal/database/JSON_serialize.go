package database

import (
	"encoding/json"
)

func ToJSON(data_struct any) string {
	data, err := json.Marshal(data_struct)
	if err != nil {
		panic(err)
	}
	return string(data)
}
