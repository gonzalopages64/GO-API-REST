FROM golang

WORKDIR /go

RUN apt update 
RUN apt install vim -y 
RUN apt install nano -y

EXPOSE 3000 

# Start process
RUN /API/exe/api.exe