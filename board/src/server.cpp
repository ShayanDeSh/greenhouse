#include <server.h>
#include "moisture.h"

AsyncWebServer server(80);

void notFound(AsyncWebServerRequest *request) {
        request->send(404, "text/plain", "Not found");
}

void route()
{
    Serial.println("routing");   
    server.on("/", HTTP_ANY, [] (AsyncWebServerRequest *request) {
        request->send(200, "text/plain", "Active");
    });

    server.on("/humidity", HTTP_GET, [](AsyncWebServerRequest *request) {
        float moist = moisture_read();
        request->send(200, "text/plain", String(moist));
    });

    server.onNotFound(notFound);    
    server.begin();
}


