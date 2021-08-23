#include "test_runner.h" 
#include "profile.h"
#include "../LinearSearch.h"
#include "../BinarySearch.h"

#include <algorithm>
#include <vector>
#include <ctime>

const int64_t SIZE_SEARCH_ARR = 10000000;
const int64_t SIZE_NUMB = 1000;

void TestSearch();
void TestStandardSearch(const std::vector<int>& seach_arr, int search_el);
void TestLinearSearch(const std::vector<int>& seach_arr, int search_el);
void TestBinarySearch(const std::vector<int>& seach_arr, int search_el);

int main() {
    TestRunner tr;

    RUN_TEST(tr, TestSearch);

	return 0;
}

void TestSearch() {
    srand(time(NULL));

    std::vector<int> seach_arr(SIZE_SEARCH_ARR);
    int search_el = std::rand() % SIZE_SEARCH_ARR;
    for (auto& el: seach_arr)
        el = std::rand() % SIZE_NUMB;
    std::sort(seach_arr.begin(), seach_arr.end());

    TestStandardSearch(seach_arr, search_el);
    TestLinearSearch(seach_arr, search_el);
    TestBinarySearch(seach_arr, search_el);
}

void TestStandardSearch(const std::vector<int>& seach_arr, int search_el) {
    LOG_DURATION("Search time: standard search algorithm");
    auto pos = std::find(seach_arr.begin(), seach_arr.end(), seach_arr[search_el]);

    ASSERT_EQUAL(*pos, seach_arr[search_el]);
}

void TestLinearSearch(const std::vector<int>& seach_arr, int search_el) {
    LOG_DURATION("Search time: linear search algorithm");
    auto pos = LinearSearch(seach_arr.begin(), seach_arr.end(), seach_arr[search_el]);

    ASSERT_EQUAL(*pos, seach_arr[search_el]);
}

void TestBinarySearch(const std::vector<int>& seach_arr, int search_el) {
    LOG_DURATION("Search time: binary search algorithm");
    auto pos = BinarySearch(seach_arr.begin(), seach_arr.end(), seach_arr[search_el]);

    ASSERT_EQUAL(*pos, seach_arr[search_el]);
}
