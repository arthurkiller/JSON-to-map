package jtm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type JsonToMap interface {
	Generate(url string) error
	Get() map[string]string
}

type jsontomap struct {
	mapS map[string]interface{}
	mapR map[string]string
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

	fmt.Println(c.mapS)

	getMap(c.mapS, c.mapR, "#NULL")
	return nil
}

func getMap(mapS map[string]interface{}, mapR map[string]string, lastkey string) {
	for key, val := range mapS {
		switch v := val.(type) {
		case float64:
			if lastkey == "#NULL" {
				mapR[string(key)] = fmt.Sprint(v)
			} else {
				mapR[lastkey+"-"+string(key)] = fmt.Sprint(v)
			}
		case string:
			if lastkey == "#NULL" {
				mapR[string(key)] = fmt.Sprint(v)
			} else {
				mapR[lastkey+"-"+string(key)] = fmt.Sprint(v)
			}
		case bool:
			if lastkey == "#NULL" {
				mapR[string(key)] = fmt.Sprint(v)
			} else {
				mapR[lastkey+"-"+string(key)] = fmt.Sprint(v)
			}
		case []interface{}:
			for i, v1 := range v {
				if v11, ok := v1.(map[string]interface{}); ok {
					if lastkey == "#NULL" {
						getMap(v11, mapR, string(key))
					} else {
						getMap(v11, mapR, lastkey+"-"+strconv.Itoa(i)+"-"+string(key))
					}
				} else {
					if lastkey == "#NULL" {
						mapR[string(key)] = fmt.Sprint(v1)
					} else {
						mapR[lastkey+"-"+string(key)] = fmt.Sprint(v1)
					}
				}
			}
		case map[string]interface{}:
			if lastkey == "#NULL" {
				getMap(v, mapR, string(key))
			} else {
				getMap(v, mapR, lastkey+"-"+string(key))
			}
		case interface{}:
			if v1, ok := v.(map[string]interface{}); ok {
				if lastkey == "#NULL" {
					getMap(v1, mapR, string(key))
				} else {
					getMap(v1, mapR, lastkey+"-"+string(key))
				}

			} else {
				if lastkey == "#NULL" {
					mapR[string(key)] = fmt.Sprint(v)
				} else {
					mapR[lastkey+"-"+string(key)] = fmt.Sprint(v)
				}
			}
		default:
			if lastkey == "#NULL" {
				mapR[string(key)] = fmt.Sprint(v)
			} else {
				mapR[lastkey+"-"+string(key)] = fmt.Sprint(v)
			}
		}
	}
}
