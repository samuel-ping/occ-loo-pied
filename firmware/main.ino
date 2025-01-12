#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>
#include <WiFiClient.h>
#include <ArduinoJson.h>

#include "arduino_secrets.h"

const char* API_URL = "http://bath.room/occupied";
const int DELAY = 1000;

void setup() {
  Serial.begin(9600);
  pinMode(LED_BUILTIN, OUTPUT);

  WiFi.begin(SSID, PASSWORD);
  Serial.println("Connecting");

  while(WiFi.status() != WL_CONNECTED) {
    delay(DELAY);
    Serial.print(".");
  }

  Serial.println("");
  Serial.print("Connected to WiFi network with IP Address: ");
  Serial.println(WiFi.localIP());
}

void loop() {
   if(WiFi.status() == WL_CONNECTED) {
      bool occupied = isOccupied();
      if(occupied) {
        digitalWrite(LED_BUILTIN, LOW);
      } else {
        digitalWrite(LED_BUILTIN, HIGH);
      }
    } else {
      Serial.println("WiFi Disconnected");
    }

    delay(DELAY);
}

bool isOccupied() {
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

String makeGetRequest() {
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