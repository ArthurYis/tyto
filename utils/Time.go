/*
 * @Description:
 * @Author: Arthur
 * @Date: 2019-09-20 14:18:28
 * @LastEditTime: 2019-09-20 14:18:50
 * @LastEditors: Arthur
 */
package utils

import (
	"time"
)

const (
	FIVE_MINUTE_STEP = 5 * 60       //5分钟
	ONE_HOUR_STEP    = 60 * 60      // 一小时
	ONE_DAY_STEP     = 24 * 60 * 60 //一天
)

func GetUnix13NowTime() int64 {
	return time.Now().UTC().UnixNano() / 1000000
}

func GetUnix13Time(t time.Time) int64 {
	return t.UnixNano() / 1000000
}

//获取传入日期零点时间
func GetZeroTime(t time.Time) time.Time {
	zt, _ := time.ParseInLocation("2006-01-02", t.Format("2006-01-02"), time.Local)
	return zt
}

//获取当天日期零点时间
func GetNowZeroTime() time.Time {
	zt, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local)
	return zt
}

//获取第n个时间间隔的时间
func GetTimeByIndex(zeroTime time.Time, timeStep, index int64) time.Time {

	return zeroTime.Add(time.Duration(timeStep*index) * time.Second)
}

//获取距传入时间最近时间间隔点是一天中的第几个
func GetIndexOfTime(t time.Time, timeStep int64) int64 {
	zeroTime := GetZeroTime(t)
	index := (t.Unix() - zeroTime.Unix()) / timeStep
	return index
}
