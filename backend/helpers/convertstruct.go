package helpers

import "encoding/json"

func ConvertStruct[T any](input interface{}) (*T, error) {

	jsonOutput, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	var output T
	err = json.Unmarshal(jsonOutput, &output)

	return &output, err

}
