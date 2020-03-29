#include "connection.h"

struct con_config 
{
    String SSID;
    String pass;
} config;

bool connect() 
{
    config.SSID = "SSID";
    config.pass = "PASS";
    WiFi.begin(config.SSID, config.pass);
    Serial.print("Connecting");
    while (WiFi.status() != WL_CONNECTED)
    {
        delay(500);
        Serial.print(".");
    }
    Serial.println();
    Serial.print("Connected, IP address: ");
    Serial.println(WiFi.localIP());
}

