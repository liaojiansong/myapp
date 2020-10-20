package weather

import (
	"encoding/json"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

// 天气
type Weather struct {
	Code int `json:"code"`
	Data struct {
		City struct {
			CityID        int    `json:"cityId"`
			Counname      string `json:"counname"`
			Ianatimezone  string `json:"ianatimezone"`
			Name          string `json:"name"`
			Pname         string `json:"pname"`
			Secondaryname string `json:"secondaryname"`
			Timezone      string `json:"timezone"`
		} `json:"city"`
		Condition struct {
			Condition   string `json:"condition"`
			Humidity    string `json:"humidity"`
			Icon        string `json:"icon"`
			Temp        string `json:"temp"`
			Updatetime  string `json:"updatetime"`
			Vis         string `json:"vis"`
			WindDegrees string `json:"windDegrees"`
			WindDir     string `json:"windDir"`
			WindLevel   string `json:"windLevel"`
		} `json:"condition"`
	} `json:"data"`
	Msg string `json:"msg"`
	Rc  struct {
		C int    `json:"c"`
		P string `json:"p"`
	} `json:"rc"`
}

// 获取天气
func GetWeather(cityId string) *Weather {
	w := &Weather{}
	//天气
	cfg := g.Cfg()
	url := cfg.GetString("weather.url")
	authCode := cfg.GetString("weather.appcode")
	client := g.Client()
	client.SetHeader("Authorization", authCode)

	clientResponse, err := client.Post(url, map[string]string{"cityId": cityId})
	if err != nil {
		glog.Error(err)
	}
	defer clientResponse.Close()
	if clientResponse.StatusCode != 200 {
		glog.Errorf("get weather failed; response is :%s", clientResponse.Raw())
		return w
	}
	jsonRes := clientResponse.ReadAllString()
	err = json.Unmarshal([]byte(jsonRes), w)
	if err != nil {
		glog.Errorf("unmarshal weather json failed;%s", err.Error())
	}
	return w
}
