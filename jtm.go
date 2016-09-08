package jtm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//JSONToMap object : provide the convertion of json to map
type JSONToMap interface {
	Generate(url string) error
	Get() map[string]string
}

type jsontomap struct {
	mapS map[string]interface{}
	mapR map[string]string
}

//Newjtm init a jtm stauct
func Newjtm() JSONToMap {
	return &jsontomap{}
}

func getMap(mapS map[string]interface{}, mapR map[string]string, lastkey string) {
	for key, val := range mapS {
		switch v := val.(type) {
		case float64:
			if lastkey == "#NULL" {
				mapR[fmt.Sprint(key)] = fmt.Sprint(v)
			} else {
				mapR[lastkey+"-"+fmt.Sprint(key)] = fmt.Sprint(v)
			}
		case string:
			if lastkey == "#NULL" {
				mapR[fmt.Sprint(key)] = fmt.Sprint(v)
			} else {
				mapR[lastkey+"-"+fmt.Sprint(key)] = fmt.Sprint(v)
			}
		case bool:
			if lastkey == "#NULL" {
				mapR[fmt.Sprint(key)] = fmt.Sprint(v)
			} else {
				mapR[lastkey+"-"+fmt.Sprint(key)] = fmt.Sprint(v)
			}
		case map[string]interface{}:
			if lastkey == "#NULL" {
				getMap(v, mapR, fmt.Sprint(key))
			} else {
				getMap(v, mapR, lastkey+"-"+fmt.Sprint(key))
			}
		case []interface{}:
			for i, v1 := range v {
				if v11, ok := v1.(map[string]interface{}); ok {
					if lastkey == "#NULL" {
						getMap(v11, mapR, fmt.Sprint(key))
					} else {
						getMap(v11, mapR, lastkey+"-"+strconv.Itoa(i)+"-"+fmt.Sprint(key))
					}
					//			} else if v11, ok := v1.([]interface{}); ok {
					//				if lastkey == "#NULL" {
					//					getMap(v11, mapR, fmt.Sprint(key))
					//				} else {
					//					getMap(v11, mapR, lastkey+"-"+strconv.Itoa(i)+"-"+fmt.Sprint(key))
					//				}
				} else {
					if lastkey == "#NULL" {
						mapR[fmt.Sprint(key)+"-"+fmt.Sprint(i)] = fmt.Sprint(v1)
					} else {
						mapR[lastkey+"-"+fmt.Sprint(key)+"-"+fmt.Sprint(i)] = fmt.Sprint(v1)
					}
				}
			}
		default:
			if lastkey == "#NULL" {
				mapR[fmt.Sprint(key)] = fmt.Sprint(v)
			} else {
				mapR[lastkey+"-"+fmt.Sprint(key)] = fmt.Sprint(v)
			}
		}
	}
}

func (c *jsontomap) Generate(url string) error {
	c.mapS = make(map[string]interface{}, 1024)
	c.mapR = make(map[string]string, 1024)

	content := make([]byte, 102400)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error in http get: ", err)
		return err
	}
	content, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("io readall error: ", err)
		return err
	}
	err = json.Unmarshal(content, &c.mapS)
	if err != nil {
		fmt.Println("format json error : ", err)
	}

	getMap(c.mapS, c.mapR, "#NULL")
	return nil
}

func (c *jsontomap) Get() map[string]string {
	return c.mapR
}
