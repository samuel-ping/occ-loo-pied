#include <Adafruit_NeoPixel.h>
#include <ESP8266WiFi.h>

#include "ApiClient.h"
#include "arduino_secrets.h"

#define TOUCH_PIN D5
#define LED_PIN D6

const int NUM_PIXELS = 16;
Adafruit_NeoPixel strip = Adafruit_NeoPixel(NUM_PIXELS, LED_PIN, NEO_GRB + NEO_KHZ800);
const uint32_t COLOR = strip.Color(255, 0, 0);

const int DELAY = 1000;
ApiClient apiClient;

int isTouched = 0; // if TTP223 is on
bool occupied = false;

void setup() {
  Serial.begin(9600);

  pinMode(LED_BUILTIN, OUTPUT);
  pinMode(TOUCH_PIN, INPUT);
  pinMode(LED_PIN, OUTPUT);
  strip.begin();

  WiFi.begin(SECRET_SSID, SECRET_PASSWORD);
  
  Serial.print("Connecting");
  while(WiFi.status() != WL_CONNECTED) {
    delay(DELAY);
    Serial.print(".");
  }

  Serial.print("Wifi connected with IP Address: ");
  Serial.println(WiFi.localIP());
}

void loop() {
  if(WiFi.status() == WL_CONNECTED) {
    if(digitalRead(TOUCH_PIN) == HIGH) {
      strip.fill(COLOR, 0, NUM_PIXELS);

      if(!occupied) {
        apiClient.setOccupiedRequest(true);
        occupied = true;
      }
    } else {
      strip.clear();
      if(occupied) {
        apiClient.setOccupiedRequest(false);
        occupied = false;
      }
    }
  } else {
    Serial.println("Wifi disconnected");
  }
  
  strip.show();
  delay(DELAY);
}

