.PHONY: install build run

APP_NAME := sorting-algo-visualizer.exe

install:
	@echo Installing dependencies...
	@go get


build:
	@echo Building binary...
	@go build -o bin/$(APP_NAME)

run: build
	@echo Running the binary...
	@./bin/$(APP_NAME)
