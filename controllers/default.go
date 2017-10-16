package controllers

import (
	"github.com/astaxie/beego"
	"github.com/lflxp/ips/utils"
	"flag"
)

var Data *[]utils.Origin
var Locations *map[string]utils.CityLocations
var Asn *[]utils.AsnBlocks
var path = flag.String("path", "./data", "GeoIP2 文件目录")
var port = flag.String("port", "8080", "端口")

func init() {
	flag.Parse()
	Data, Locations, Asn = utils.NewOrigin(*path)
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Check() {
	ip := c.Ctx.Input.Param(":ip")
	json := utils.ParseIp(Data, Locations, Asn, ip)

	c.Data["json"] = map[string]interface{}{
		"time": json.Time,
		"ip":json.Ip,
		"GeoIP":map[string]interface{}{
			"Locations":map[string]interface{}{
				"Geoname_id":json.Locations.GeonameId,
				"LocaleCode":json.Locations.LocaleCode,
				"ContinentCode":json.Locations.ContinentCode,
				"ContinentName":json.Locations.ContinentName,
				"CountryIsoCode":json.Locations.CountryIsoCode,
				"CountryName":json.Locations.CountryName,
				"S1IsoCode":json.Locations.S1IsoCode,
				"S1Name":json.Locations.S1Name,
				"S2IsoCode":json.Locations.S2IsoCode,
				"S2Name":json.Locations.S2Name,
				"CityName":json.Locations.CityName,
				"MetroCode":json.Locations.MetroCode,
				"TimeZone":json.Locations.TimeZone,
			},
			"Blocks":map[string]interface{}{
				"Start":json.Blocks.Start,
				"End":json.Blocks.End,
				"FirstIp":json.Blocks.FirstIp,
				"EndIp":json.Blocks.EndIp,
				"Network":json.Blocks.Network,
				"Geoname_id":json.Blocks.Geoname_id,
				"Registered_country_geoname_id":json.Blocks.Registered_country_geoname_id,
				"Represented_country_geoname_id":json.Blocks.Represented_country_geoname_id,
				"Is_anonymous_proxy":json.Blocks.Is_anonymous_proxy,
				"Is_satellite_provider":json.Blocks.Is_satellite_provider,
				"Postal_code":json.Blocks.Postal_code,
				"Latitude":json.Blocks.Latitude,
				"Longitude":json.Blocks.Longitude,
				"Accuracy_radius":json.Blocks.Accuracy_radius,
			},
			"Asn":map[string]interface{}{
				"Network":json.Asn.Network,
				"Autonomous_system_number":json.Asn.Autonomous_system_number,
				"Autonomous_system_organization":json.Asn.Autonomous_system_organization,
			},
		},
		"status":json.Status,
	}
	c.ServeJSON()
}
