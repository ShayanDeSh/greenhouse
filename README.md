# greenhouse

This program is consisted of two parts: server and board

##Server

- This is the part which is responsible for listening to boards request and storing
the sent data to database. This part is developed using go programming language and uses
postgresql as database so you must have both installed on you computer

- To execute server you should run the following command:
cd server; go run ./cmd/main.go

##Board

This program is built upon arduino and  platformio so you must have both installed

- This is the program which will be run on the ESP board and mainly is consisted of two parts:

- first part is an api which is responsible for listening to requests from user to read and return the
humidiy of plant you can simply make this request by typing the following url into your browser:
"http://<board-ip>/humidity"
NOTICE: the board ip is returned to serial out put when you first run your program

- second part is responsible for reading data from sensor periodically and sending it to
server at specific times through interface designed on server

- Almost no memory is used for saving data on arduino as we use an external database for saving data
and for using less energy for sending data we only read and send data due to specifc times

- For running this program after you plugged you board to computer run the following command:
cd board; platformio run
