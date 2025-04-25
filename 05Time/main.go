package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Say Hi to time packgae")

	pstTime := time.Now()
	fmt.Println("Current time :", pstTime)
	//Current time : 2025-04-25 18:19:23.5005374 +0530 IST m=+0.000506901

	// Formating Time
	fmt.Println(pstTime.Format("01-02-2006 Monday 15:04:05"))
	// 04-25-2025 Friday 18:21:40

	// Create Custom Date / Time
	createdDate := time.Date(2025, time.April, 1, 12, 12, 0, 0, time.Local)
	// time.Date(year(int),month(time.month),day(int),hour(int),min(int),sec(int),nanosec(int),timezone(time.zone))
	fmt.Println(createdDate)
	// 2025-04-01 12:12:00 +0530 IST
	fmt.Println(createdDate.Format("Monday 01-02-2006"))
	// Tuesday 04-01-2025

	// Wait for user input to keep the terminal open
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()

	// GOOS=darwin GOARCH=arm64 go build  or just GOOS=darwin go build
	// GOOS=linux GOARCH=amd64 go build || GOOS=linux go build
	// GOOS=windows GOARCH=amd64 go build || GOOS=windows go build

	// if using windows powershell
	//$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o time_linux
	//$env:GOOS="darwin"; $env:GOARCH="amd64"; go build -o time_mac

	// cmd
	// set GOOS=linux
	// set GOARCH=amd64
	// go build -o time_linux

}
