/**
 * Created by GoLand.
 * User: adrienli@tencent.com
 * Date: 2019/1/24
 * Time: 4:19 PM
 * Desc:
 */
package data

type ApiResponse struct {
	Status   string   `json:"status"`
	Count    string   `json:"count"`
	Info     string   `json:"info"`
	InfoCode string   `json:"infocode"`
	Lives    []LivesArr `json:"lives"`
}

type LivesArr struct {
	Province      string `json:"province"`
	City          string `json:"city"`
	Adcode        string `json:"adcode"`
	Weather       string `json:"weather"`
	Temperature   string `json:"temperature"`
	WindDirection string `json:"winddirection"`
	WindPower     string `json:"windpower"`
	Humidity      string `json:"humidity"`
	ReportTime    string `json:"reporttime"`
}
