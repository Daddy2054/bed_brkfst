#!/bin/bash

go build -o bookings cmd/web/*.go &&./bookings 

#-dbname=bookings -dbuser=someuser -cache=false -production=false