cmake_minimum_required(VERSION 3.24)
include(FetchContent)
# include(CMakePrintHelpers)

set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${CMAKE_SOURCE_DIR}/lib)
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_SOURCE_DIR}/bin)
set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${CMAKE_SOURCE_DIR}/lib)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS ON)
set(CMAKE_CXX_FLAGS_DEBUG
    "${CMAKE_CXX_FLAGS_DEBUG} -g3 -Wall -Wextra -O1 -Wformat=2")
set(CMAKE_CXX_FLAGS_RELEASE "${CMAKE_CXX_FLAGS_RELEASE} -O3")
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

project(daikin-home-ac)

set(OPENSSL_ROOT_DIR /usr/local/opt/openssl)

# LIST(APPEND CMAKE_MODULE_PATH "${CMAKE_CURRENT_LIST_DIR}/cmake") Add packages
include("cmake/drogon.cmake")

file(GLOB_RECURSE HEADER_FILES "src/*.hpp")
file(GLOB_RECURSE SOURCE_FILES "src/*.cpp")

add_executable(daikin-service ${SOURCE_FILES} ${HEADER_FILES})
# target_include_directories(daikin-service PRIVATE )
target_link_libraries(daikin-service PRIVATE drogon)
# target_precompile_headers(daikin-service PRIVATE src/pch.hpp)
