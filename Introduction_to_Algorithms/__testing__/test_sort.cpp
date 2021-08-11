#include "test_runner.h"
#include "profile.h"
#include "../InsertionSort.h"
#include "../SelectionSort.h"

#include <algorithm>
#include <vector>

const int64_t SIZE_SORT_ARR = 10000;
const int64_t SIZE_NUMB = 1000;

void TestInsertionSort();
void TestSelectionSort();

int main() {
    TestRunner tr;
    RUN_TEST(tr, TestInsertionSort);
    RUN_TEST(tr, TestSelectionSort);
    return 0;
}

void TestInsertionSort() {
    {
        std::vector<int> true_sort(SIZE_SORT_ARR);
        for (auto& el: true_sort)
            el = std::rand() % SIZE_NUMB;

        std::vector<int> result(true_sort);
        {
            LOG_DURATION("sorting time: standard sort algorithm");
            std::sort(true_sort.begin(), true_sort.end());
        }
        {
            LOG_DURATION("sorting time: insertion sort algorithm");
            InsertionSort(result.begin(), result.end());
        }
        ASSERT_EQUAL(result, true_sort);
    }
    {
        std::vector<int> true_sort(SIZE_SORT_ARR);
        for (auto& el: true_sort)
            el = std::rand() % SIZE_NUMB;

        std::vector<int> result(true_sort);
        {
            LOG_DURATION("sorting time: standard sort algorithm with compare");
            std::sort(true_sort.begin(), true_sort.end(), [](auto& l, auto& r) {
                return l > r;
            });
        }
        {
            LOG_DURATION("sorting time: insertion sort algorithm with compare");
            InsertionSort(result.begin(), result.end(), [](auto& l, auto& r) {
                return l > r;
            });
        }
        ASSERT_EQUAL(result, true_sort);
    }
}

void TestSelectionSort() {
    {
        std::vector<int> true_sort(SIZE_SORT_ARR);
        for (auto& el: true_sort)
            el = std::rand() % SIZE_NUMB;

        std::vector<int> result(true_sort);
        {
            LOG_DURATION("sorting time: standard sort algorithm");
            std::sort(true_sort.begin(), true_sort.end());
        }
        {
            LOG_DURATION("sorting time: selection sort algorithm");
            SelectionSort(result.begin(), result.end());
        }
        ASSERT_EQUAL(result, true_sort);
    }
    {
        std::vector<int> true_sort(SIZE_SORT_ARR);
        for (auto& el: true_sort)
            el = std::rand() % SIZE_NUMB;

        std::vector<int> result(true_sort);
        {
            LOG_DURATION("sorting time: standard sort algorithm with compare");
            std::sort(true_sort.begin(), true_sort.end(), [](auto& l, auto& r) {
                return l > r;
            });
        }
        {
            LOG_DURATION("sorting time: selection sort algorithm with compare");
            SelectionSort(result.begin(), result.end(), [](auto& l, auto& r) {
                return l > r;
            });
        }
        ASSERT_EQUAL(result, true_sort);
    }
}
