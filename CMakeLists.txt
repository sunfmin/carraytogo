CMAKE_MINIMUM_REQUIRED(VERSION 3.14 FATAL_ERROR)
set(CMAKE_CXX_STANDARD 17)

project(felix)

add_library(hello SHARED
    hello.cpp
)
