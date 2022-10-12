#!/usr/bin/python3
import time
import sys
from sense_hat import SenseHat
import json

def main():
    sense = SenseHat()

    temperature = str(sense.get_temperature())
    humidity = sense.get_humidity()
    pressure = sense.get_pressure()

    measureObj = {
        "temperature": str(temperature),
        "humidity": str(humidity),
        "pressure": str(pressure)
    }
    
    measureStr = json.dumps(measureObj)

    print(measureStr)
    

if __name__ == "__main__":
    main()
