package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ptrv/go-gpx"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		log.Fatalln("Please provide a GPX file path!")
	}

	gpxFileArg := args[0]
	gpxFile, err := gpx.ParseFile(gpxFileArg)

	if err != nil {
		log.Fatalln("Error opening gpx file: ", err)
	}

	fmt.Println("# distance, sufferscore (hr/speed)")

	distance := 0.0

	for _, trk := range gpxFile.Tracks {
		for _, segment := range trk.Segments {
			for i, wpt := range segment.Waypoints {
				spd := segment.Waypoints.Speed(i) * 3.6
				hr := wpt.Hr

				if hr < 110 || spd > 70 || spd < 12 {
					continue
				}

				// fmt.Printf("dist: %.2f\tspd: %.1f\thr: %v\n", distance, spd, hr)

				suffer := float64(hr) / spd

				fmt.Printf("%.2f, %.2f\n", distance, suffer)

				if i > 0 {
					prevWpt := segment.Waypoints[i-1]
					distance += wpt.Distance2D(&prevWpt) / 1000
				}
			}
		}
	}
}
