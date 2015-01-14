package sunriseset

import (
	"log"
	"testing"
)

func TestSunriseset(t *testing.T) {
	srs := &SunRiseSet{Lat: 34.1234, Long: 123.43434, Date: "2014-05-28"}
	log.Println(srs.GetSunRise())
	log.Println(srs.GetSunSet())
}
