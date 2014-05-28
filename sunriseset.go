package sunriseset

import (
	"log"
	"math"
	"strconv"
	"time"
)

const h = -0.833
const pi = 3.1415126

type SunRiseSet struct {
	uto  float64 `float64:180.0`
	lat  float64
	long float64
	date string
	rise string
	set  string
}

//GetRise 获得日出时间
func (s *SunRiseSet) GetRise() string {
	return s.rise
}

//GetSet 获得日落时间
func (s *SunRiseSet) GetSet() string {
	return s.set
}

//SunRise 计算日出时间
func (s *SunRiseSet) SunRise() *SunRiseSet {
	s.rise = toTime(result_rise(ut_rise(s.uto, gha(s.uto, g_sun(t_century(days(s.date), s.uto)), ecliptic_longitude(l_sun(t_century(days(s.date), s.uto)), g_sun(t_century(days(s.date), s.uto)))), s.long, e(h, s.lat, sun_deviation(earth_tilt(t_century(days(s.date), s.uto)), ecliptic_longitude(l_sun(t_century(days(s.date), s.uto)), g_sun(t_century(days(s.date), s.uto)))))), s.uto, s.long, s.lat, s.date))
	return s
}

//SunSet 计算日落时间
func (s *SunRiseSet) SunSet() *SunRiseSet {
	s.set = toTime(result_set(ut_rise(s.uto, gha(s.uto, g_sun(t_century(days(s.date), s.uto)), ecliptic_longitude(l_sun(t_century(days(s.date), s.uto)), g_sun(t_century(days(s.date), s.uto)))), s.long, e(h, s.lat, sun_deviation(earth_tilt(t_century(days(s.date), s.uto)), ecliptic_longitude(l_sun(t_century(days(s.date), s.uto)), g_sun(t_century(days(s.date), s.uto)))))), s.uto, s.long, s.lat, s.date))
	return s
}

//toTime 装换成时间
func toTime(temp float64) string {
	return strconv.Itoa(int(temp/15.0+8.0)) + ":" + strconv.Itoa(int(((temp/15+8)-float64(int(temp/15+8)))*60.0))
}

//days 求从格林威治时间公元2000年1月1日到计算日天数days
func days(date string) int {
	end, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Println("Fatal error:%s", err.Error())
	}
	start, err := time.Parse("2006-01-02", "2000-01-01")
	if err != nil {
		log.Println("Fatal error:%s", err.Error())
	}
	days := int(end.Sub(start).Seconds() / 3600 / 24)
	return days
}

//t_century 求格林威治时间公元2000年1月1日到计算日的世纪数t
func t_century(days int, UTo float64) float64 {
	return (float64(days) + UTo/360.0) / 36525.0
}

//l_sun 求太阳的平黄径
func l_sun(t_century float64) float64 {
	return (280.460 + 36000.770*t_century)
}

//g_sun 求太阳的平近点角
func g_sun(t_century float64) float64 {
	return (357.528 + 35999.050*t_century)
}

//earth_tilt 求地球倾角
func earth_tilt(t_century float64) float64 {
	return (23.4393 - 0.0130*t_century)
}

//ecliptic_longitude 求黄道经度
func ecliptic_longitude(l_sun, g_sun float64) float64 {
	return (l_sun + 1.915*math.Sin(g_sun*pi/180) + 0.02*math.Sin(2*g_sun*pi/180))
}

//gha 求格林威治时间的太阳时间角GHA
func gha(UTo, g_sun, ecliptic_longitude float64) float64 {
	return (UTo - 180 - 1.915*math.Sin(g_sun*pi/180) - 0.02*math.Sin(2*g_sun*pi/180) + 2.466*math.Sin(2*ecliptic_longitude*pi/180) - 0.053*math.Sin(4*ecliptic_longitude*pi/180))
}

//sun_deviation
func sun_deviation(earth_tilt, ecliptic_longitude float64) float64 {
	return (180 / pi * math.Asin(math.Sin(pi/180*earth_tilt)*math.Sin(pi/180*ecliptic_longitude)))
}

//e 求修正值e
func e(h, glat, sun_deviation float64) float64 {
	return 180 / pi * math.Acos((math.Sin(h*pi/180)-math.Sin(glat*pi/180)*math.Sin(sun_deviation*pi/180))/(math.Cos(glat*pi/180)*math.Cos(sun_deviation*pi/180)))
}

//ut_rise 求日出时间
func ut_rise(UTo, GHA, glong, e float64) float64 {
	return (UTo - (GHA + glong + e))
}

//ut_set 求日落时间
func ut_set(UTo, GHA, glong, e float64) float64 {
	return (UTo - (GHA + glong - e))
}

//result_rise 判断并返回结果（日出）
func result_rise(UT, UTo, glong, glat float64, date string) float64 {
	var d float64
	if UT >= UTo {
		d = UT - UTo
	} else {
		d = UTo - UT
	}
	if d >= 0.1 {
		UTo = UT
		UT = ut_rise(UTo,
			gha(UTo, g_sun(t_century(days(date), UTo)),
				ecliptic_longitude(l_sun(t_century(days(date), UTo)),
					g_sun(t_century(days(date), UTo)))),
			glong,
			e(h, glat, sun_deviation(earth_tilt(t_century(days(date), UTo)),
				ecliptic_longitude(l_sun(t_century(days(date), UTo)),
					g_sun(t_century(days(date), UTo))))))
		result_rise(UT, UTo, glong, glat, date)
	}
	return UT
}

//result_set 判断并返回结果（日落）
func result_set(UT, UTo, glong, glat float64, date string) float64 {
	var d float64
	if UT >= UTo {
		d = UT - UTo
	} else {
		d = UTo - UT
	}
	if d >= 0.1 {
		UTo = UT
		UT = ut_set(UTo,
			gha(UTo, g_sun(t_century(days(date), UTo)),
				ecliptic_longitude(l_sun(t_century(days(date), UTo)),
					g_sun(t_century(days(date), UTo)))),
			glong,
			e(h, glat, sun_deviation(earth_tilt(t_century(days(date), UTo)),
				ecliptic_longitude(l_sun(t_century(days(date), UTo)),
					g_sun(t_century(days(date),
						UTo))))))
		result_set(UT, UTo, glong, glat, date)
	}
	return UT
}

//zone 时区
func zone(glong float64) int {
	if glong >= 0 {
		return int(glong/15.0) + 1
	} else {
		return int(glong/15.0) - 1
	}
}
