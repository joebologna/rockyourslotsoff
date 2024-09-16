cmake_minimum_required(VERSION 3.10)
project(HelloWorld)

set(CMAKE_CXX_STANDARD 11)

set(FILE_URL "https://raw.githubusercontent.com/doctest/doctest/master/doctest/doctest.h")
set(LOCAL_FILE "${CMAKE_CURRENT_SOURCE_DIR}/doctest.h")

if(NOT EXISTS ${LOCAL_FILE})
    message(STATUS "Downloading file from ${FILE_URL}...")
    execute_process(
        COMMAND curl -s -L -o ${LOCAL_FILE} ${FILE_URL}
        RESULT_VARIABLE result
    )
    
    if(result)
        message(FATAL_ERROR "Failed to download file from ${FILE_URL}")
    else()
        message(STATUS "File downloaded successfully.")
    endif()
else()
    message(STATUS "File already exists: ${LOCAL_FILE}")
endif()

enable_testing()
add_subdirectory(tests)

add_executable(HelloWorld main.cpp)
