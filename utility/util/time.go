package util

import (
	"github.com/gogf/gf/v2/os/gtime"
	"time"
)

const (
	DATE_TIME_FORMAT       = "2006-01-02 15:04:05"
	DATE_TIME_MONTH_FORMAT = "01-02 15:04:05"
	DATE_MINUTE_FORMAT     = "2006-01-02 15:04"
	DATE_FORMAT            = "2006-01-02"
)

type DateTime struct {
	StartTimestamp int64
	EndTimestamp   int64
	StartDate      string
	StartDateTime  string
	EndDate        string
	EndDateTime    string
}

var Location *time.Location

func init() {
	Location, _ = time.LoadLocation("Asia/Shanghai")
}

func FormatDateTimeMonth(timestamp int64) string {
	if timestamp == 0 {
		return ""
	}
	return gtime.NewFromTimeStamp(timestamp).Layout(DATE_TIME_MONTH_FORMAT)
}

func FormatDateTime(timestamp int64) string {
	if timestamp == 0 {
		return ""
	}
	return gtime.NewFromTimeStamp(timestamp).String()
}

// 过期时间转换
func ConvExpiresAt(expiresAt string) int64 {
	if expiresAt == "" {
		return 0
	}
	return gtime.NewFromStrLayout(expiresAt, time.DateTime).TimestampMilli() + 999
}

func Day(startTime, endTime string) (dateTimeList []*DateTime) {

	dateTimeList = make([]*DateTime, 0)

	defer func() {
		if len(dateTimeList) > 0 {
			dateTime := dateTimeList[len(dateTimeList)-1]
			dateTime.EndDateTime = endTime
			dateTime.EndDate = dateTime.EndDateTime[:10]
			t, _ := time.ParseInLocation(DATE_TIME_FORMAT, dateTime.EndDateTime, Location)
			dateTime.EndTimestamp = t.Unix()
		}
	}()

	sDateTime, err := time.ParseInLocation(DATE_TIME_FORMAT, startTime, Location)
	if err != nil {
		return dateTimeList
	}

	eDateTime, err := time.ParseInLocation(DATE_TIME_FORMAT, endTime, Location)
	if err != nil {
		return dateTimeList
	}

	if eDateTime.Before(sDateTime) {
		return dateTimeList
	}

	eDateStr := eDateTime.Format(DATE_FORMAT)

	startDateTime := time.Date(sDateTime.Year(), sDateTime.Month(), sDateTime.Day(), sDateTime.Hour(), sDateTime.Minute(), sDateTime.Second(), 0, Location)
	endDateTime := time.Date(sDateTime.Year(), sDateTime.Month(), sDateTime.Day(), 23, 59, 59, 0, Location)

	dateTime := &DateTime{
		StartTimestamp: startDateTime.Unix(),
		EndTimestamp:   endDateTime.Unix(),
		StartDate:      startDateTime.Format(DATE_FORMAT),
		StartDateTime:  startDateTime.Format(DATE_TIME_FORMAT),
		EndDate:        endDateTime.Format(DATE_FORMAT),
		EndDateTime:    endDateTime.Format(DATE_TIME_FORMAT),
	}

	dateTimeList = append(dateTimeList, dateTime)

	if dateTime.StartDate == eDateStr {
		return dateTimeList
	}

	for {

		sDateTime = sDateTime.AddDate(0, 0, 1)
		sDateStr := sDateTime.Format(DATE_FORMAT)
		startDateTime := time.Date(sDateTime.Year(), sDateTime.Month(), sDateTime.Day(), 0, 0, 0, 0, Location)
		endDateTime := time.Date(sDateTime.Year(), sDateTime.Month(), sDateTime.Day(), 23, 59, 59, 0, Location)

		dateTime := &DateTime{
			StartTimestamp: startDateTime.Unix(),
			EndTimestamp:   endDateTime.Unix(),
			StartDate:      startDateTime.Format(DATE_FORMAT),
			StartDateTime:  startDateTime.Format(DATE_TIME_FORMAT),
			EndDate:        endDateTime.Format(DATE_FORMAT),
			EndDateTime:    endDateTime.Format(DATE_TIME_FORMAT),
		}

		dateTimeList = append(dateTimeList, dateTime)
		if sDateStr == eDateStr {
			break
		}
	}

	return dateTimeList
}

func Week(startTime, endTime string) (dateTimeList []*DateTime) {

	dateTimeList = make([]*DateTime, 0)
	dateTimeList = weekRange(startTime, endTime, dateTimeList)

	return dateTimeList
}

func Month(startTime, endTime string) (dateTimeList []*DateTime) {

	dateTimeList = make([]*DateTime, 0)

	dateTimeList = monthRange(startTime, endTime, dateTimeList)

	return dateTimeList
}

func weekRange(startTime, endTime string, dateTimeList []*DateTime) []*DateTime {

	sDateTime, _ := time.ParseInLocation(DATE_TIME_FORMAT, startTime, Location)
	year, week := sDateTime.ISOWeek()
	startWeekTime := weekStart(year, week)

	endWeekTime := startWeekTime.AddDate(0, 0, 6)

	endDate := time.Date(endWeekTime.Year(), endWeekTime.Month(), endWeekTime.Day(), 23, 59, 59, 0, Location)

	sTime := startWeekTime.Format(DATE_TIME_FORMAT)
	eTime := endDate.Format(DATE_TIME_FORMAT)

	if sTime < startTime {
		sTime = startTime
	}

	if eTime > endTime {
		eTime = endTime
	}

	startDateTime, _ := time.ParseInLocation(DATE_TIME_FORMAT, sTime, Location)
	endDateTime, _ := time.ParseInLocation(DATE_TIME_FORMAT, eTime, Location)

	dateTime := DateTime{
		StartTimestamp: startDateTime.Unix(),
		EndTimestamp:   endDateTime.Unix(),
		StartDate:      startDateTime.Format(DATE_FORMAT),
		StartDateTime:  sTime,
		EndDate:        endDateTime.Format(DATE_FORMAT),
		EndDateTime:    eTime,
	}

	dateTimeList = append(dateTimeList, &dateTime)

	if dateTime.EndDateTime < endTime {
		sDateTime := time.Date(endDateTime.Year(), endDateTime.Month(), endDateTime.Day(), 0, 0, 0, 0, Location)
		startTime = sDateTime.AddDate(0, 0, 1).Format(DATE_TIME_FORMAT)
		dateTimeList = weekRange(startTime, endTime, dateTimeList)
	}

	return dateTimeList
}

func weekStart(year, week int) time.Time {

	t := time.Date(year, 7, 1, 0, 0, 0, 0, Location)

	if wd := t.Weekday(); wd == time.Sunday {
		t = t.AddDate(0, 0, -6)
	} else {
		t = t.AddDate(0, 0, -int(wd)+1)
	}

	_, w := t.ISOWeek()
	t = t.AddDate(0, 0, (week-w)*7)

	return t
}

func monthRange(startTime, endTime string, dateTimeList []*DateTime) []*DateTime {

	sDateTime, eDateTime := monthStartAndEnd(startTime, endTime)

	startDateTime, _ := time.ParseInLocation(DATE_TIME_FORMAT, sDateTime, Location)
	endDateTime, _ := time.ParseInLocation(DATE_TIME_FORMAT, eDateTime, Location)

	dateTime := DateTime{
		StartTimestamp: startDateTime.Unix(),
		EndTimestamp:   endDateTime.Unix(),
		StartDate:      startDateTime.Format(DATE_FORMAT),
		StartDateTime:  sDateTime,
		EndDate:        endDateTime.Format(DATE_FORMAT),
		EndDateTime:    eDateTime,
	}

	dateTimeList = append(dateTimeList, &dateTime)

	if eDateTime < endTime {
		eTime, _ := time.ParseInLocation(DATE_TIME_FORMAT, dateTime.EndDateTime, Location)
		sTime := time.Date(eTime.Year(), eTime.Month()+1, 1, 0, 0, 0, 0, Location).Format(DATE_TIME_FORMAT)
		dateTimeList = monthRange(sTime, endTime, dateTimeList)
	}

	return dateTimeList
}

func monthStartAndEnd(startTime, endTime string) (sDateTime, eDateTime string) {

	sTime, _ := time.ParseInLocation(DATE_TIME_FORMAT, startTime, Location)

	year := sTime.Year()
	month := sTime.Month()

	sDateTime = time.Date(year, month, 1, 0, 0, 0, 0, Location).Format(DATE_TIME_FORMAT)
	eDateTime = time.Date(year, month+1, 0, 23, 59, 59, 0, Location).Format(DATE_TIME_FORMAT)

	if sDateTime < startTime {
		sDateTime = startTime
	}

	if eDateTime > endTime {
		eDateTime = endTime
	}

	return sDateTime, eDateTime
}

func MinuteRange(startTime, endTime int64) []*DateTime {

	startTimestamp, _ := time.ParseInLocation(DATE_MINUTE_FORMAT, time.Unix(startTime, 0).Format(DATE_MINUTE_FORMAT), Location)
	endTimestamp, _ := time.ParseInLocation(DATE_MINUTE_FORMAT, time.Unix(endTime, 0).Format(DATE_MINUTE_FORMAT), Location)

	dateTimes := make([]*DateTime, 0)

	for timestamp := startTimestamp.Unix(); timestamp < endTimestamp.Unix(); timestamp += 60 {
		dateTime := new(DateTime)
		dateTime.StartTimestamp = timestamp
		dateTimes = append(dateTimes, dateTime)
	}

	return dateTimes
}
