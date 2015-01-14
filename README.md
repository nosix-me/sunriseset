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
		"github.com/nosix-me/sunriseset"
		"fmt"
	)

	func main() {
		srs := &sunriseset.SunRiseSet{Lat: 39.92, Long: 116.46, Date: "2015-01-14"}
		fmt.Println(srs.GetSunRise())
		fmt.Println(srs.GetSunSet())
	}



	
