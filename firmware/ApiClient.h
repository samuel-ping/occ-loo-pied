#ifndef ApiClient_h
#define ApiClient_h
#include "Arduino.h"

class ApiClient {
  public:
    ApiClient();
    bool isOccupied();
  private:
    int _pin;
    String makeGetRequest();
};

#endif