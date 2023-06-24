/*
 * File: main.go
 * File Created: Tuesday, 13th June 2023 3:19:45 pm
 * Last Modified: Friday, 23rd June 2023 4:45:12 pm
 * Author: Akhil Datla
 * Copyright © Akhil Datla 2023
 */

package main

import (
	"flag"
	"main/components/courses"
	"main/server"

	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"github.com/pterm/pterm"
)

func main() {
	// Parse command-line flags
	portPtr := flag.Int("port", 3000, "Port to listen on")
	flag.Parse()

	// Load environment variables from the "variables.env" file
	err := godotenv.Load("variables.env")
	if err != nil {
		pterm.Error.Println(err)
	}

	// Download courses
	err = courses.DownloadCourses()
	if err != nil {
		pterm.Error.Println(err)
	}

	// Load courses
	err = courses.LoadCourses()
	if err != nil {
		pterm.Error.Println(err)
	}

	// Schedule cron jobs
	go schedule()

	// Display the banner
	banner()

	// Start the server
	server.Start(*portPtr)
}

// schedule schedules the cron jobs.
func schedule() {
	gocron.Every(1).Hour().Do(courses.DownloadCourses)
	gocron.Every(1).Hour().Do(courses.LoadCourses)
	<-gocron.Start()
}

// banner displays a custom banner message.
func banner() {
	pterm.DefaultCenter.Print(pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Sprint("Learnado: Igniting Minds, Inspiring Learning"))
	pterm.Info.Println("Student Edition")
	pterm.Info.Println("(c)2023 by Akhil Datla")
}
