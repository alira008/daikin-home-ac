#include "temperature-controller.hpp"

#include <iostream>
#include <string>

namespace controller {

void TemperatureController::asyncHandleHttpRequest(
    const drogon::HttpRequestPtr &req,
    std::function<void(const drogon::HttpResponsePtr &)> &&callback) {
  std::string value = req->getParameter("value");
  Json::Value json_res;
  json_res["success"] = value;
  auto resp = drogon::HttpResponse::newHttpJsonResponse(json_res);
  callback(resp);
}

}  // namespace controller
