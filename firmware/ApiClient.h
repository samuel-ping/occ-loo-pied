#ifndef ApiClient_h
#define ApiClient_h

class ApiClient {
  const char* apiUrl;

  public:
    ApiClient(const char*);
    bool isOccupied();
    void setOccupiedRequest(bool occupied);
  private:
    int _pin;
    String getOccupiedRequest();
};

#endif