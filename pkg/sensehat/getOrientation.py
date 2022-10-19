#!/usr/bin/python3
from sense_hat import SenseHat

def main():
    sense = SenseHat()
    orientation = sense.orientation
    orientation = str(orientation).replace("'", '"')
    print(orientation)

if __name__ == "__main__":
    main()