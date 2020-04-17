#include "Arduino.h"
#include <ESP8266WiFi.h>
#include "connection.h"
#include "server.h"

void setup()
{
    Serial.begin(115200);
    Serial.println();

    connect();
    route();
}

void loop()
{
}
