package main

import (
    "fmt"
    forecast "github.com/mlbright/forecast/v2"
    gcfg "code.google.com/p/gcfg"
    "log"
    "flag"
    "errors"
    "strconv"
)


type Config struct {
    Profile map[string]*struct {
        Name string
        Lat  string
        Long string
    }  
}

var cityFlag = flag.String("p", "06", "Hava durumu istenen şehrin plaka kodu")


func main() {

    flag.Parse()    

    key := "e096247fcf57315b7af6fe75efefa5c5"

    var cfg Config
    err := gcfg.ReadFileInto(&cfg, "citiesConfig.cfg")
    if err != nil {
        log.Fatal(errors.New("Şehir bilgi dosyası bulunamadı ya da okunamadı!"))
    } 

    if nCityFlag, err := strconv.Atoi(*cityFlag); err != nil || nCityFlag < 1 || nCityFlag > 81 {
        log.Fatal(errors.New("Geçersiz il plaka kodu!") )
    }


    lat := cfg.Profile[*cityFlag].Lat
    long := cfg.Profile[*cityFlag].Long
    name := cfg.Profile[*cityFlag].Name

    f, err := forecast.Get(key, lat, long, "now", forecast.SI)
    if err != nil {
        log.Fatal(errors.New("Hava durumu ölçüm merkezine bağlantıda hata oluştu!"))
    }

    var wc string

    switch f.Currently.Icon {
    case "clear-day":
        wc = "Açık bir gün"
    case "clear-night":
        wc = "Açık bir gece"
    case "rain":
        wc = "Yağmurlu"
    case "snow":
        wc = "Karlı"
    case "sleet":
        wc = "Buzlu"
    case "wind":
        wc = "Rüzgarlı"
    case "fog":
        wc = "Sisli"
    case "cloudy":
        wc = "Bulutlu"
    case "partly-cloudy-day":
        wc = "Parçalı bulutlu bir gün"
    case "partly-cloudy-night":
        wc = "Parçalı bulutlu bir gece"
    default:
        wc = "Hava durumu bilinmiyor ?"
    }

    sep := "-----------------------------------------------------"

    fmt.Printf("\n%s - %s \n", name, wc)
    fmt.Println(sep)
    fmt.Println("Sıcaklık(°C) | Orantısal Nem(Y) | Rüzgar Hızı (m/sn)")
    fmt.Println(sep)
    fmt.Printf("%12.2f | %16.2f | %18.2f\n\n",  
                            f.Currently.Temperature,
                            f.Currently.Humidity * 100.0,
                            f.Currently.WindSpeed)

}
