package sunriseset

import (
	"log"
	"testing"
)

func TestSunriseset(t *testing.T) {
	srs := &SunRiseSet{uto: 180.0, lat: 34.1234, long: 123.43434, date: "2014-05-28"}
	log.Println(srs.SunRise().GetRise())
	log.Println(srs.SunSet().GetSet())
}
