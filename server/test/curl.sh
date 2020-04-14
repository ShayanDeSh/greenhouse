#!/bin/bash
curl -X POST -H 'Content-Type:application/json' -d '{"Ip":"192.168.2.6","Humidity":2, "NeedWater":false, "Threshhold": 25}' http://localhost:8000/v1/plants/add
