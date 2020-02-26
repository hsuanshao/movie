# movie
a tcp example to query movie information and a sample monitor on port 8083


how to execute it.

- go get https://github.com/hsuanshao/movie

if you want to modify code to do more application
please  

git clone project

- if you are going to execute it,
after go build
execute in terminal movie

than you can use

telnet 127.0.0.1:8081

to test, enter movie name, than you will reseive movie info,
if that movie is exists.

and you can use your browser to open

http://127.0.0.1:8083

you can get the simple monitor dashboard, descripbe following numbers:

Current connection: 
Current Request rate: 
Processed request:  
Remaing jobs: 