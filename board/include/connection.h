#include <ESP8266WiFi.h>
#include "Arduino.h"
#include <stdbool.h>

struct con_config 
{
    String SSID;
    String pass;
    String IP;
    String serverAdr;
};

struct con_config connect();

