#include "Arduino.h"
#include <ESP8266WiFi.h>
#include "server.h"
#include <ESP8266HTTPClient.h>
#include "moisture.h"

void setup()
{
    Serial.begin(115200);
    Serial.println();

    struct con_config config =  connect();
    route(config);
    periodic_read(config.IP, config.serverAdr);
}

void loop()
{
}
