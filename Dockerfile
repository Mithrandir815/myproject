FROM node:latest
WORKDIR /frontend
COPY ./frontend /frontend/

FROM golang:latest
WORKDIR /backend
COPY ./backend /backend/ 