#ifndef ApiClient_h
#define ApiClient_h

class ApiClient {
  public:
    ApiClient();
    bool isOccupied();
    void setOccupiedRequest(bool occupied);
  private:
    int _pin;
    String getOccupiedRequest();
};

#endif