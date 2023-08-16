#pragma once

#include <drogon/HttpSimpleController.h>

#include <functional>

namespace controller {

class TemperatureController
    : public drogon::HttpSimpleController<TemperatureController> {
 public:
  void asyncHandleHttpRequest(
      const drogon::HttpRequestPtr &req,
      std::function<void(const drogon::HttpResponsePtr &)> &&callback) override;
  PATH_LIST_BEGIN
  PATH_ADD("/temperature", drogon::Get);
  PATH_LIST_END
};

}  // namespace controller
