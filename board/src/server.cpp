#include <server.h>
#include "moisture.h"
#include "AsyncJson.h"
#include "ArduinoJson.h"

AsyncWebServer server(80);

void notFound(AsyncWebServerRequest *request) {
    request->send(404, "text/plain", "Not found");
}


void route(struct con_config config)
{
    Serial.println("routing");   
    server.on("/", HTTP_ANY, [] (AsyncWebServerRequest *request) {
        request->send(200, "text/plain", "Active");
    });

    server.on("/humidity", HTTP_GET, [](AsyncWebServerRequest *request) {
        float moist = moisture_read();
        request->send(200, "text/plain", String(moist));
    });

    server.on("/server", HTTP_POST, [config](AsyncWebServerRequest *request) {
        request->send(201, "text/plain", "Server Added");
    });

    server.onNotFound(notFound);    
    server.begin();
}
