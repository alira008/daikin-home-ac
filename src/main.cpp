#include <drogon/drogon.h>

#include <iostream>

int main() {
  drogon::app().addListener("127.0.0.1", 5040).run();

  return 1;
}
