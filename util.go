package gocountries

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/liuzl/go.geojson"
	"io/ioutil"
	"sync"
)

type CountryInfo struct {
	Name         string                     `json:"name"`
	AltSpellings []string                   `json:"altSpellings"`
	Area         float64                    `json:"area"`
	Borders      []string                   `json:"borders"`
	CallingCodes []string                   `json:"callingCodes"`
	Capital      string                     `json:"capital"`
	Currencies   []string                   `json:"currencies"`
	Demonym      string                     `json:"demonym"`
	Flag         string                     `json:"flag"`
	GeoJSON      *geojson.FeatureCollection `json:"geoJSON"`
	ISO          map[string]string          `json:"ISO"`
	Languages    []string                   `json:"languages"`
	Latlng       []float32                  `json:"latlng"`
	NativeName   string                     `json:"nativeName"`
	Population   int64                      `json:"population"`
	Provinces    []string                   `json:"provinces"`
	Region       string                     `json:"region"`
	Subregion    string                     `json:"subregion"`
	Timezones    []string                   `json:"timezones"`
	Tld          []string                   `json:"tld"`
	Translations map[string]string          `json:"translations"`
	Wiki         string                     `json:"wiki"`
}

var countries map[string]*CountryInfo
var once sync.Once

func loadData() {
	once.Do(func() {
		countries = make(map[string]*CountryInfo)
		for _, body := range countriesFile {
			data, err := base64.StdEncoding.DecodeString(body)
			if err != nil {
				fmt.Println(err)
				continue
			}
			reader, err := gzip.NewReader(bytes.NewReader(data))
			if err != nil {
				fmt.Println(err)
				continue
			}
			rawBytes, err := ioutil.ReadAll(reader)
			if err != nil {
				continue
			}
			var item CountryInfo
			err = json.Unmarshal(rawBytes, &item)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if item.Name == "" || item.ISO == nil || item.ISO["alpha2"] == "" {
				//fmt.Printf("%s: %+v\n\n", file, item)
				continue
			}
			countries[item.ISO["alpha2"]] = &item
		}
		fmt.Printf("total %d countries infomation loaded\n", len(countries))
	})
}

func FindCountryByAlpha(code string) *CountryInfo {
	loadData()
	return countries[code]
}
