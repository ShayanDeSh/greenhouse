#include "connection.h"
#include "ArduinoJson.h"
#include <ESP8266HTTPClient.h>

struct con_config config;

String con() {
    StaticJsonDocument<300> doc;
    doc["Ip"] = config.IP;
    doc["Humidity"] = 0;
    doc["NeedWater"] = false;
    doc["Threshhold"] = 120;
    String a ;
    serializeJsonPretty(doc, a);
    return a;
}

struct con_config connect() 
{
    config.SSID = "shayan2";
    config.pass = "pwolpcgeplir";
    config.serverAdr = "http://192.168.1.41:9000";
    WiFi.begin(config.SSID, config.pass);
    Serial.print("Connecting");
    while (WiFi.status() != WL_CONNECTED)
    {
        delay(500);
        Serial.print(".");
    }
    Serial.println();
    config.IP = WiFi.localIP().toString();
    Serial.print("Connected, IP address: ");
    Serial.println(config.IP);
    auto out = con();
    HTTPClient http;
    http.begin(config.serverAdr + "/v1/plants/add");
    http.addHeader("Content-Type", "application/json"); 
    int x  = http.POST(out);
    Serial.println(out);
    Serial.println(x);
    return config;
}

