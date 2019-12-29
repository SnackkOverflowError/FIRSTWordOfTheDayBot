package Utility

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

func ReadJson() (string, int) {

	botData := make(map[string]interface{})
	data, errRead := ioutil.ReadFile("data/botData.json")
	if errRead != nil {
		panic(errRead)
	}


	errUnMarshall := json.Unmarshal(data, &botData)

	if errUnMarshall != nil {
		panic(errUnMarshall)
	}


	if botData["token"] == nil || botData["token"].(string) == "" {
		panic(errors.New("need to add token"))
	}
	if botData["current_index"] == nil {
		botData["current_index"] = 0
		toSave, errMarshall := json.Marshal(botData)

		if errMarshall != nil {
			panic(errMarshall)
		}
		errWrite := ioutil.WriteFile("data/botData.json",toSave, 0644)
		if errWrite != nil {
			panic(errWrite)
		}

	}

	return botData["token"].(string),int(botData["current_index"].(float64))

}

func UpdateIndex(index int,token string) {
	botData := make(map[string]interface{})
	botData["current_index"] = index
	botData["token"] = token
	toSave, errMarshall := json.Marshal(botData)
	if errMarshall != nil {
		panic(errMarshall)
	}
	errWrite := ioutil.WriteFile("data/botData.json",toSave, 0644)
	if errWrite != nil {
		panic(errWrite)
	}

}
