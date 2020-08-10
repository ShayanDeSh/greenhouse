#include "moisture.h"
#include <ESP8266HTTPClient.h>
#include "ArduinoJson.h"

#define SensorPin A0

String data(int hum, String ip) {
    StaticJsonDocument<300> doc;
    doc["Ip"] = ip;
    doc["Humidity"] = hum;
    if (hum > 120) { 
        doc["NeedWater"] = false;
    }
    else {
        doc["NeedWater"] = true;
    }
    doc["Threshhold"] = 120;
    String a;
    serializeJsonPretty(doc, a);
    Serial.println(a);
    return a;
}

float moisture_read() {
    float moisture = 0;
    for (int i = 0; i < 100; i++)
    {
        moisture += analogRead(SensorPin);
        delay(1);
    }
    moisture = moisture / 100;
    Serial.println(moisture);
    return moisture;
}

void periodic_read(String ip, String serverAdr) {
    while (true) {
        float moist = moisture_read();
        auto json = data(moist, ip);
        HTTPClient http;
        http.begin(serverAdr + "/v1/plants/update");
        http.addHeader("Content-Type", "application/json"); 
        http.POST(json);
        delay(5000);
    }
}

