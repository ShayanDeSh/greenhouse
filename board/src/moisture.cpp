#include "moisture.h"

#define SensorPin A0

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
