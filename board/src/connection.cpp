#include "connection.h"

struct con_config 
{
    String SSID;
    String pass;
    String IP;
} config;

bool connect() 
{
    config.SSID = "shayan2";
    config.pass = "pwolpcgeplir";
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
}

