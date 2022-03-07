#!/bin/bash

API_KEY=$OPENWEATHER_API_KEY
CITY=$OPENWEATHER_CITY

curl -s --location --request GET \
'http://api.openweathermap.org/data/2.5/weather?appid='$API_KEY'&q='$CITY'&units=metric' | jq
