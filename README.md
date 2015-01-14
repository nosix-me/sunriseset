sunriseset
========
[![Build Status](https://drone.io/github.com/nosix-me/sunriseset/status.png)](https://drone.io/github.com/nosix-me/sunriseset/latest)
##描述

计算日出日落

##安装方法
	
	go get github.com/nosix-me/sunriseset

##使用方法
    //北京的经纬度 116.46,39.92
    package main

	import (
		"fmt"
		"github.com/nosix-me/sunriseset"
	)

	func main() {
		srs := &SunRiseSet{Lat: 34.1234, Long: 123.43434, Date: "2014-05-28"}
		log.Println(srs.GetSunRise())
		log.Println(srs.GetSunSet())
	}



	
