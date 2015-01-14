package sunriseset

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"
)

const (
	UTo = 180.0
	PI  = 3.1415126
	H   = -0.833
)

type SunRiseSet struct {
	UTo  float64 `float64:180.0`
	Date string
	Long float64
	Lat  float64
}

//GetSunRise 获得日出时间
func (srs *SunRiseSet) GetSunRise() string {
	t := t_century(srs.Date, srs.UTo)
	l := l_sun(t)
	g := g_sun(t)
	ecliptic_Longitude := ecliptic_Longitude(l, g)
	earth_tilt := earth_tilt(t)
	GHA := gha(UTo, g, ecliptic_Longitude)
	sun_deviation := sun_deviation(earth_tilt, ecliptic_Longitude)
	e := e(srs.Lat, sun_deviation)
	UT := ut_rise(UTo, GHA, srs.Long, e)
	return toTime(srs, result_rise(UT, UTo, srs.Long, srs.Lat, srs.Date))
}

//GetSunSet 获得日落时间
func (srs *SunRiseSet) GetSunSet() string {
	t := t_century(srs.Date, srs.UTo)
	l := l_sun(t)
	g := g_sun(t)
	ecliptic_Longitude := ecliptic_Longitude(l, g)
	earth_tilt := earth_tilt(t)
	GHA := gha(UTo, g, ecliptic_Longitude)
	sun_deviation := sun_deviation(earth_tilt, ecliptic_Longitude)
	e := e(srs.Lat, sun_deviation)
	UT := ut_set(UTo, GHA, srs.Long, e)
	return toTime(srs, result_set(UT, UTo, srs.Long, srs.Lat, srs.Date))
}

//toTime 时间格式化
func toTime(srs *SunRiseSet, result float64) string {
	value := result/15.0 + float64(zone(srs.Long))
	decimal := value - float64(int(value))
	minutes := strconv.Itoa(int(decimal * 60))
	hour := ""
	if int(value) > 9 {
		hour = strconv.Itoa(int(value))
	} else {
		hour = "0" + strconv.Itoa(int(value))
	}
	return srs.Date + " " + hour + ":" + minutes + ":00"
}

//days 求从格林威治时间公元2000年1月1日到计算日天数days
func days(Date string) int {
	start, err := time.Parse("2006-01-02", "2000-01-01")
	if err != nil {
		log.Println("Fatal error:%s", err.Error())
	}
	end, err := time.Parse("2006-01-02", Date)
	if err != nil {
		log.Println("Fatal error:%s", err.Error())
	}
	days := int(end.Sub(start).Seconds() / 3600 / 24)
	return days
}

//t_century 求格林威治时间公元2000年1月1日到计算日的世纪数t
func t_century(Date string, UTo float64) float64 {
	days := days(Date)
	return (float64(days) + UTo/360.0) / 36525.0
}

//l_sun 求太阳的平黄径
func l_sun(t float64) float64 {
	return 280.460 + 36000.770*t
}

//g_sun 求太阳的平近点角
func g_sun(t float64) float64 {
	return 357.528 + 35999.050*t
}

//ecliptic_Longitude 求黄道经度
func ecliptic_Longitude(l float64, g float64) float64 {
	return l + 1.915*math.Sin(g*PI/180.0) + 0.020*math.Sin(2*g*PI/180.0)
}

//earth_tilt 求地球倾角
func earth_tilt(t float64) float64 {
	return 23.4393 - 0.0130*t
}

//sun_deviation 计算太阳的偏差
func sun_deviation(earth_tilt, ecliptic_Longitude float64) float64 {
	return 180.0 / PI * math.Asin(math.Sin(earth_tilt*PI/180.0)*math.Sin(ecliptic_Longitude*PI/180.0))
}

//gha 求格林威治时间的太阳时间角GHA
func gha(UTo, g, ecliptic_Longitude float64) float64 {
	return UTo - 180.0 - 1.915*math.Sin(g*PI/180.0) - 0.020*math.Sin(2*g*PI/180.0) + 2.466*math.Sin(2*ecliptic_Longitude*PI/180.0) - 0.053*math.Sin(4*ecliptic_Longitude*PI/180.0)
}

//e 求修正值e
func e(gLat, sun_deviation float64) float64 {
	return 180.0 / PI * math.Acos((math.Sin(H*PI/180.0)-math.Sin(gLat*PI/180.0)*math.Sin(sun_deviation*PI/180.0))/(math.Cos(gLat*PI/180.0)*math.Cos(sun_deviation*PI/180.0)))
}

//ut_rise 求日出时间
func ut_rise(UTo, GHA, gLong, e float64) float64 {
	return (UTo - (GHA + gLong + e))
}

//ut_set 求日落时间
func ut_set(UTo, GHA, gLong, e float64) float64 {
	return (UTo - (GHA + gLong - e))
}

func result_rise(UT, UTo, gLong, gLat float64, Date string) float64 {
	var d float64
	if UT >= UTo {
		d = UT - UTo
	} else {
		d = UTo - UT
	}
	if d >= 0.1 {
		UTo = UT
		t := t_century(Date, UTo)
		g := g_sun(t)
		l := l_sun(t)
		ecliptic_Longitude := ecliptic_Longitude(l, g)
		earth_tilt := earth_tilt(t)
		GHA := gha(UTo, g, ecliptic_Longitude)
		sun_deviation := sun_deviation(earth_tilt, ecliptic_Longitude)
		e := e(gLat, sun_deviation)
		UT = ut_rise(UTo, GHA, gLong, e)
		result_rise(UT, UTo, gLong, gLat, Date)
	}
	return UT
}

func result_set(UT, UTo, gLong, gLat float64, Date string) float64 {
	var d float64
	if UT >= UTo {
		d = UT - UTo
	} else {
		d = UTo - UT
	}
	if d >= 0.1 {
		UTo = UT
		t := t_century(Date, UTo)
		g := g_sun(t)
		l := l_sun(t)
		ecliptic_Longitude := ecliptic_Longitude(l, g)
		earth_tilt := earth_tilt(t)
		sun_deviation := sun_deviation(earth_tilt, ecliptic_Longitude)
		GHA := gha(UTo, g, ecliptic_Longitude)
		e := e(gLat, sun_deviation)
		UT = ut_set(UTo, GHA, gLong, e)
		result_set(UT, UTo, gLong, gLat, Date)
	}
	return UT
}

//zone 计算时区
func zone(gLong float64) int {
	if gLong >= 0 {
		return int(gLong/15.0) + 1
	} else {
		return int(gLong/15.0) - 1
	}
}
