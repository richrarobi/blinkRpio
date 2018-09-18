# blinkRpio
A port of Pimoroni's Blinkt Python library in go
Created this today, each Python def converted to go. I have tried it on a pi3 and a pi2.
If you get any issues, please do let me know. I am starting full testing tomorrow. It isn't a package yet.
I haven't read how to do that yet! It has functions and a main for testing....

It uses go-rpio https://github.com/stianeikeland/go-rpio. This uses the Broadcom pin numbering.

I did also get some very useful information from https://github.com/alexellis/rpi, which is based on Gordon's 
wiringpi library (I personally like the broadcom numbering and couldn't find a Blinkt library that did it)
Hence this port.

sudo apt-get install rpi.gpio

sudo apt-get update

sudo apt-get upgrade

go get github.com/stianeikeland/go-rpio

go get github.com/richrarobi/blinkRpio

see examples

Rich Robinson
