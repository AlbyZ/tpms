package main

import (
	"flag"
	"fmt"
	"github.com/AlbyZ/tpms"
	"log"
	"time"
)

var (
	duration = flag.Duration("duration", 0, "monitoring duration Xs, 0 for indefinitely")
)

func main() {
	flag.Parse()
	tires, err := tpms.NewTpms()
	if err != nil {
		log.Fatal(err)
	}
	tires.Log("./log")
	tires.StartMonitoring()
	start := time.Now()
	defer tires.StopMonitoring()
	for *duration == 0 || time.Now().Sub(start) < *duration {
		for _, sensor := range tires.Read() {
			if sensor != nil {
				// fmt.Printf("Sensor: %d, kPa: %d, °C: %d\n", sensor.Id, sensor.Kilopascal, sensor.Celsius)
				// fmt.Printf("Sensor: %d, Psi: %d, °C: %d\n", sensor.Id, sensor.Psi, sensor.Celsius)
				fmt.Printf("Sensor: %d, Psi: %.1f, °C: %.1f, %d%%\n", sensor.Id, sensor.Psi, sensor.Celsius, sensor.Percentage)
			}
		}
		time.Sleep(5 * time.Second)
	}
}
