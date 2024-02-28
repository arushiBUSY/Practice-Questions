package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
type CityWeatherInfo struct{
	Name string `json:"name"`
	Weather string `json:"weather"`
	//The data type for the "status" field in the given data structure appears to be an array. Each element within the array contains a string representing a particular status, such as "Wind: 6Kmph" and "Humidity: 5%". Therefore, the data type for the "status" field is an array of strings.

	Status []string `json:"status"`
}
type WeatherInfo struct{
	Page       int      `json:"page"`
    PerPage    int      `json:"per_page"`
    Total      int      `json:"total"`
    TotalPages int      `json:"total_pages"`
    Data       []CityWeatherInfo  `json:"data"`
}
func fetchData(apiUrl string,name string)([]CityWeatherInfo,error){
	var cities []CityWeatherInfo
	apiUrl=apiUrl+name
	//)
	//step 1->to perform get request

    //fetch data from each page
	page:=1
	for {
		//fmt.Sprintf is a function in Go's fmt package that formats strings according to a format specifier and returns the resulting string.

		url:=fmt.Sprintf("%s&page=%d",apiUrl,page)
		//url:="https://jsonmock.hackerrank.com/api/weather/search?name=B&page=2"
		response,err:=http.Get(url)
		//fmt.Println(response)
		if err!=nil{
			fmt.Println("Error in fetching data:",err)
		   return nil,err
	    }
		body,err:=io.ReadAll(response.Body)
		//fmt.Println(string(body))
		if err!=nil{
			return nil,err
	    }
		//The json.Unmarshal function is used in Go to decode JSON data into Go data structures
		var wi WeatherInfo
		err=json.Unmarshal(body,&wi)
		if err!=nil{
			return nil,err
	    }
		//In the context of append, the ... unpacks the elements of weatherData.Data, which is a slice of City structs, and appends each element individually to the allCities slice.

		cities=append(cities,wi.Data...)
		if page>=wi.TotalPages{
			break;
		}
		page++;


	}
	return cities,nil

}
func main(){
	apiUrL:="https://jsonmock.hackerrank.com/api/weather/search?name="
	fmt.Println("Enter the name whose pages you need to access:")
	var name string
	fmt.Scan(&name)
	allCities,err:=fetchData(apiUrL,name)
	if err!=nil{
		fmt.Println("Error fetching data:",err)
		return
	}
	for _,city:=range allCities{
		
		
		fmt.Println("Name:",city.Name)
		fmt.Println("Weather:", city.Weather)
        fmt.Println("Status:")
		for _,status:=range city.Status{
			//fmt.Printf("%d: ",i)
			fmt.Println(status)
		}
		fmt.Println(" ")

	}
}
