#include <ArduinoJson.h>
#include <ESP8266HTTPClient.h>
#include <WiFiClient.h>

#include "ApiClient.h"

const char* API_URL = "http://bath.room/occupied";

ApiClient::ApiClient() {}

bool ApiClient::isOccupied() {
  String payload = makeGetRequest();

  StaticJsonDocument<32> doc;
  DeserializationError error = deserializeJson(doc, payload);
  if (error) {
    Serial.print(F("deserializeJson() failed: "));
    Serial.println(error.f_str());
    return false;
  }

  return doc["occupied"];
}

String ApiClient::makeGetRequest() {
  WiFiClient client;
  HTTPClient http;
  http.begin(client, API_URL);
  
  int httpResponseCode = http.GET();
  Serial.print("HTTP Response code: ");
  Serial.println(httpResponseCode);
  String payload = http.getString();
  Serial.println(payload);  

  http.end();

  return payload;
}