#!/bin/bash
curl -X POST -H 'Content-Type:application/json' -d '{"Ip":"192.168.2.2","Humidity":23, "NeedWater":false, "Threshhold": 25}' http://localhost:8000/v1/plants/add

