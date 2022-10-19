# Orientation Simulcast Sender -> InlfuxDB Cloud

## Components
- Sender: RPi + senseHAT 
- Cloud DB: influxDB Cloud

## Features
Two Goroutines getting data and sending at different frequencies.
Data is represented as a High Resolution (with real float64 type) and Low Resolution (with casted float type from int), e.g: 356.12315 and 356.
InfluxDB currently does not support int type yet. This is a work-around solution to show-case the simulcast function
