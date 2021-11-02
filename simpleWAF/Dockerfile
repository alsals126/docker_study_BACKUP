# docker에 베이스가 될 이미지(FROM)가 없으면 자동으로 다운받는다.
FROM ubuntu:latest 
FROM golang:1.17.1

ARG DEBIAN_FRONTEND=noninteractive
# apt update와 upgrade의 차이
# - update : 설치 가능한 패키지 리스트를 최신화
# - upgrade : 실제 업데이트
RUN apt-get update

RUN mkdir /go/src/app
ADD . /go/src/app
# 작업 공간 이동 (=cd)
WORKDIR /go/src/app

# ENV POSTGRES_USER = jeongmin
# ENV POSTGRES_PASSWORD = 1234asdf#$
# ENV POSTGRES_DB = jeongmindb

RUN go build -o main . 

ENTRYPOINT  [ "./main" ]