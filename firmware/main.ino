#include <Adafruit_NeoPixel.h>
#include <ESP8266WiFi.h>

#include "ApiClient.h"
#include "arduino_secrets.h"

#define TOUCH_PIN D5
#define LED_PIN D6

const int NUM_PIXELS = 16;
Adafruit_NeoPixel strip = Adafruit_NeoPixel(NUM_PIXELS, LED_PIN, NEO_GRB + NEO_KHZ800);
const uint32_t COLOR_RED = strip.Color(200, 40, 85);
const uint32_t COLOR_GREEN = strip.Color(50, 168, 82);
const int BRIGHTNESS = 2; // (255 max)


static unsigned long builtinLedLastToggle = 0;

const int DELAY = 1000;
unsigned long lastPollTime = 0;
const unsigned long POLL_INTERVAL = 1000;  // Poll every 1 second
ApiClient apiClient;

int touchState = 0;  // if TTP223 is on
bool occupied = false;

void setup() {
  Serial.begin(9600);

  pinMode(LED_BUILTIN, OUTPUT);
  pinMode(TOUCH_PIN, INPUT);
  pinMode(LED_PIN, OUTPUT);
  strip.setBrightness(BRIGHTNESS);
  strip.begin();

  connectToWifi();
}

void loop() {
  if (WiFi.status() == WL_CONNECTED) {

    // Poll status
    if (millis() - lastPollTime > POLL_INTERVAL) {
      lastPollTime = millis();
      occupied = apiClient.isOccupied();
    }

    // Handle touch button
    if (digitalRead(TOUCH_PIN) == HIGH) {
      if (!occupied) {
        apiClient.setOccupiedRequest(true);
        occupied = true;
      }
    } else {
      if (occupied) {
        apiClient.setOccupiedRequest(false);
        occupied = false;
      }
    }
  } else {
    Serial.println("Wifi disconnected");
    connectToWifi();
  }

  // Update LED ring based on status
  if (occupied) {
    strip.fill(COLOR_RED, 0, NUM_PIXELS);
  } else {
    strip.fill(COLOR_GREEN, 0, NUM_PIXELS);
  }
  strip.show();

  delay(DELAY);
}

void connectToWifi() {
  Serial.print("Connecting to Wifi");

  WiFi.begin(SECRET_SSID, SECRET_PASSWORD);
  while (WiFi.status() != WL_CONNECTED) {
    flashBuiltinLed();
    delay(DELAY);
    Serial.print(".");
  }

  digitalWrite(LED_PIN, HIGH);
  Serial.print("Wifi connected with IP Address: ");
  Serial.println(WiFi.localIP());
}

void flashBuiltinLed() {
  static bool ledState = digitalRead(LED_PIN);

  // Toggle the LED every 300ms
  if (millis() - builtinLedLastToggle >= 300) {
    ledState = !ledState;
    digitalWrite(LED_PIN, ledState);
    builtinLedLastToggle = millis();
  }
}