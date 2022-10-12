import time
import sys
from sense_hat import SenseHat

sense = SenseHat()
pressure = sense.get_pressure()
print(pressure)