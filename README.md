# 3D Printer Go

This Go program allows you to control your 3D printer using the Marlin firmware. It provides real-time data streaming, logging, and the ability to send commands to your printer.

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)

## Introduction

3D Printer Control is a Go program that interfaces with 3D printers running the Marlin firmware. It allows users to read real-time data from the printer, log the data to a file, and send custom commands.


https://github.com/tarikcaliskan/3d-printer-go/assets/29506704/acc43da0-ad63-4596-885b-7ac5ada2a4d3



## Features

- Real-time data streaming from the 3D printer
- Logging of printer data to a specified file
- Sending custom commands to the 3D printer

## Getting Started

### Prerequisites

Make sure you have the following installed:

- Go (at least Go 1.x)
- Marlin firmware on your 3D printer

<br>

### Installation

Clone the repository and build the program:

```bash
git clone https://github.com/tarikcaliskan/3d-printer-go.git
cd 3d-printer-go
go build
```

<br>

### Usage

Run the compiled binary with the desired port, baud rate, and log file:

```bash
./3d-printer-go --port /dev/tty.usbserial-1234 --baud-rate 115200 --log-file printer_log.txt
```

This will start the program, allowing you to interact with your 3D printer.

<br>

### Configuration

You can customize the behavior of the program using command-line arguments. Here are the available options:

- `--port`: The serial port name of your 3D printer (default: "/dev/tty.usbserial-1140").
- `--baud-rate`: The baud rate for serial communication (default: 250000).
- `--log-file`: The name of the log file for storing printer data (default: "printer_log.txt").
