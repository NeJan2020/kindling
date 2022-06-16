cmake_minimum_required(VERSION 3.5.1)

project(. NONE)
# set_property(TARGET agent-lib PROPERTY C_STANDARD 99)
set(CMAKE_C_FLAGS "-std=c99")

include(ExternalProject)
message(STATUS "Driver version: ${AGENT_LIBS_VERSION}")

ExternalProject_Add(
  agent-libs
  URL "https://github.com/NeJan2020/agent-libs/archive/ad379c9be53cd78200896afa9e07e4cb649806fc.tar.gz"
  URL_HASH "${AGENT_LIBS_CHECKSUM}"
  CONFIGURE_COMMAND ""
  BUILD_COMMAND ""
  INSTALL_COMMAND ""
  TEST_COMMAND "")