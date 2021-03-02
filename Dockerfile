FROM golang:latest
RUN apt-get -y update
RUN apt-get -y install git

RUN git clone https://github.com/akshitIsFuture/Parking-lot-Service.git