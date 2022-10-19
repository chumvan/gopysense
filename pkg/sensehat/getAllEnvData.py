#!/usr/bin/python3
from sense_hat import SenseHat
import json

def main():
    sense = SenseHat()

    temperature = sense.get_temperature()
    humidity = sense.get_humidity()
    pressure = sense.get_pressure()

    measureObj = {
        "temperature": temperature,
        "humidity": humidity,
        "pressure": pressure
    }

    measureStr = json.dumps(measureObj)

    print(measureStr)
    

if __name__ == "__main__":
    main()
