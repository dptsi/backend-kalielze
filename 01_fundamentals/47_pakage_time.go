package main

import (
	"fmt"
	"time"
)

func main() {

	waktu := time.Now()
	fmt.Println("time now", waktu)
	fmt.Println("year", waktu.Year())
	fmt.Println("year day", waktu.YearDay())
	fmt.Println("month", waktu.Month())
	fmt.Println("day", waktu.Day())
	fmt.Println(waktu.Date())
	fmt.Println("hour", waktu.Hour())
	fmt.Println("minute", waktu.Minute())
	fmt.Println("second", waktu.Second())
	fmt.Println("nanosecond", waktu.Nanosecond())
	fmt.Println("unix", waktu.Unix())
	fmt.Println("unixmicro", waktu.UnixMicro())
	fmt.Println("unixmili", waktu.UnixMilli())
	fmt.Println("unixnano", waktu.UnixNano())

	utc := time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)

	// layout := time.RFC3339
	parse, _ := time.Parse("2006-01-02", "2022-04-01") // membuat tanggal
	fmt.Println(parse)

	// periode := parse.Format("2006") // untuk mengambil periode tahun dalam bentuk yyyy
	// periode := parse.Format("01") // untuk mengambil periode bulan dalam bentuk mm
	// periode := parse.Format("02") // untuk mengambil periode tanggal dalam bentuk dd
	fmt.Println("periode tahun:", parse.Format("2006"))
	fmt.Println("periode bulan:", parse.Format("01"))
	fmt.Println("periode tanggal:", parse.Format("02"))

	fmt.Println("tahun:", parse.Year())
	fmt.Println("bulan:", parse.Month())
	fmt.Println("tanggal:", parse.Day())

	/*
		masih banyak documentation mengenai package time

		https://golang.org/pkg/time/
	*/
}
