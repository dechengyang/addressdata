package data

import (
	"io/ioutil"
	"fmt"
	"os"
	"encoding/json"
	"bufio"
)

//town level address
//乡镇、街道等级的地址
type Town struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	DistrictCode string `json:"areaCode"`
	CityCode     string `json:"cityCode"`
	ProvinceCode string `json:"provinceCode"`
}

//district level address
//区县级别的地址
type District struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	CityCode     string `json:"cityCode"`
	ProvinceCode string `json:"provinceCode"`
}

//city level address
//城市级别的地址
type City struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	ProvinceCode string `json:"provinceCode"`
}

//province level address
//省份级别的地址
type Province struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Parser struct {
	Provinces []Province
	Cities    []City
	Districts []District
	Towns     []Town
}

func (p *Parser) Generate() {
	p.parseProvince()
	p.parseCity()
	p.parseDistrict()
	p.parseTown()

	var results []string

	for _, t := range p.Towns {
		d := p.district(t.DistrictCode)
		if d != nil {
			c := p.city(t.CityCode)
			if c != nil {
				p := p.province(t.ProvinceCode)
				if p != nil {
					if "北京市" == p.Name || "上海市" == p.Name ||
						"天津市" == p.Name || "重庆市" == p.Name {
						results = append(results, "\"\",\""+p.Name+"\",\""+d.Name+"\",\""+t.Name+"\"")
					} else {
						results = append(results, "\""+p.Name+"\",\""+c.Name+"\",\""+d.Name+"\",\""+t.Name+"\"")
					}
				}
			}
		}
	}

	file, err := os.Create("./data/pcdt.csv")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	fmt.Fprintln(w, "province,city,district,town")
	for _, line := range results {
		fmt.Fprintln(w, line)
	}
	w.Flush()

}

func (p *Parser) district(code string) *District {
	for _, district := range p.Districts {
		if code == district.Code {
			return &district
		}
	}
	return nil
}

func (p *Parser) city(code string) *City {
	for _, city := range p.Cities {
		if code == city.Code {
			return &city
		}
	}
	return nil
}

func (p *Parser) province(code string) *Province {
	for _, province := range p.Provinces {
		if code == province.Code {
			return &province
		}
	}
	return nil
}

func (p *Parser) parse(file string, code int) {
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if code == 1 {
		err = json.Unmarshal(raw, &p.Provinces)
	} else if code == 2 {
		err = json.Unmarshal(raw, &p.Cities)
	} else if code == 3 {
		err = json.Unmarshal(raw, &p.Districts)
	} else if code == 4 {
		err = json.Unmarshal(raw, &p.Towns)
	}
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func (p *Parser) parseProvince() {
	p.parse("./data/province.json", 1)
}

func (p *Parser) parseCity() {
	p.parse("./data/city.json", 2)
}

func (p *Parser) parseDistrict() {
	p.parse("./data/district.json", 3)
}

func (p *Parser) parseTown() {
	p.parse("./data/town.json", 4)
}
