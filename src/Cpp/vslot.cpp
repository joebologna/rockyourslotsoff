#define DOCTEST_CONFIG_IMPLEMENT_WITH_MAIN
#include "doctest.h"
#include <array>
#include <cstdlib>
#include <ctime>

// VSlot interface defines the required methods for implementing a virtual slot machine.
class VSlot {
public:
    virtual std::array<int, 3> Spin() = 0;
    virtual void Reset() = 0;
    virtual void UpdateBalance(int amount) = 0;
    virtual int GetBalance() const = 0;
    virtual ~VSlot() = default;
};

// MyVSlot class implements the VSlot interface.
class MyVSlot : public VSlot {
private:
    int balance;

public:
    MyVSlot(int64_t seed) : balance(0) {
        std::srand(seed);
    }

    std::array<int, 3> Spin() override {
        return {std::rand() % 10 + 1, std::rand() % 10 + 1, std::rand() % 10 + 1};
    }

    void Reset() override {
        balance = 0;
    }

    void UpdateBalance(int amount) override {
        balance += amount;
    }

    int GetBalance() const override {
        return balance;
    }
};

TEST_CASE("MyVSlotTest") {
    MyVSlot slotMachine(12345); // Initialize with a fixed seed for reproducibility

    SUBCASE("Spin generates three numbers between 1 and 10") {
        auto result = slotMachine.Spin();
        CHECK(result[0] >= 1);
        CHECK(result[0] <= 10);
        CHECK(result[1] >= 1);
        CHECK(result[1] <= 10);
        CHECK(result[2] >= 1);
        CHECK(result[2] <= 10);
    }

    SUBCASE("Balance updates correctly") {
        slotMachine.UpdateBalance(100);
        CHECK(slotMachine.GetBalance() == 100);

        slotMachine.UpdateBalance(-50);
        CHECK(slotMachine.GetBalance() == 50);
    }

    SUBCASE("Reset sets balance to zero") {
        slotMachine.UpdateBalance(100);
        slotMachine.Reset();
        CHECK(slotMachine.GetBalance() == 0);
    }
}
