#include <ArduinoJson.h>
#include <ESP8266HTTPClient.h>
#include <WiFiClient.h>

#include "ApiClient.h"

HTTPClient http;

ApiClient::ApiClient(const char* u) {
  apiUrl = u;
}

bool ApiClient::isOccupied() {
  String payload = getOccupiedRequest();

  StaticJsonDocument<32> doc;
  DeserializationError error = deserializeJson(doc, payload);
  if (error) {
    Serial.print(F("deserializeJson() failed: "));
    Serial.println(error.f_str());
    return false;
  }

  return doc["occupied"];
}

String ApiClient::getOccupiedRequest() {
  WiFiClient client;
  http.begin(client, apiUrl);
  
  int httpResponseCode = http.GET();
  Serial.print("GET response code: ");
  Serial.println(httpResponseCode);
  String payload = http.getString();
  Serial.println(payload);  

  http.end();

  return payload;
}

void ApiClient::setOccupiedRequest(bool occupied) {
  JsonDocument doc;
  doc["occupied"] = occupied;
  char payload[24];
  serializeJson(doc, payload);

  WiFiClient client;
  http.begin(client, apiUrl);

  http.addHeader("Content-Type", "text/json");

  Serial.print("PUT payload: ");
  Serial.println(payload);
  int httpResponseCode = http.PUT(payload);
  Serial.print("PUT response code: ");
  Serial.println(httpResponseCode);

  http.end();
}