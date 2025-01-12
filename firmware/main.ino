#include <ESP8266WiFi.h>

#include "ApiClient.h"
#include "arduino_secrets.h"

const int DELAY = 1000;
ApiClient apiClient;

void setup() {
  Serial.begin(9600);
  pinMode(LED_BUILTIN, OUTPUT);

  WiFi.begin(SECRET_SSID, SECRET_PASSWORD);
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
      bool occupied = apiClient.isOccupied();
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

