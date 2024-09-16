#define DOCTEST_CONFIG_IMPLEMENT_WITH_MAIN
#include "doctest.h"
#include <array>
#include <cstdlib>

std::array<int, 3> spin(unsigned int seed) {
  std::srand(seed);
  return {int((std::rand() % 10) + 1), int((std::rand() % 10) + 1), int((std::rand() % 10) + 1)};
}

TEST_CASE("check if spin worked") {
  std::srand(0);
  std::array<int, 3> expected = {int((std::rand() % 10) + 1), int((std::rand() % 10) + 1), int((std::rand() % 10) + 1)};
  std::array<int, 3> result = spin(0);
  CHECK(result == expected);
}
