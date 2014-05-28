sunriseset
========
[![Build Status](https://drone.io/github.com/widuu/goini/status.png)](https://drone.io/github.com/nosix-me/sunriseset/4)

##描述

计算日出日落

##安装方法
	
	go get github.com/nosix-me/sunriseset

##使用方法
    //北京的经纬度 116.46,39.92
    rsr := &SunRiseSet{uto: 180.0, lat: 39.92, long: 116.46, date: "2014-05-28"}
    rise := srs.SunRise().GetRise()
    set := srs.SunSet().GetSet()



	
