#include "test_runner.h"

#include <unordered_map>
#include <algorithm>
#include <iostream>
#include <vector>
#include <numeric>

// my first solution
class Solution {
public:
    std::vector<int> twoSum(std::vector<int>& nums, int target) {
        std::vector<int> result(2);
        result.assign(result.size(), -1);
        std::vector<int> indexs(nums.size());
        std::iota(indexs.begin(), indexs.end(), 0);
        std::sort(indexs.begin(), indexs.end(),
            [&nums](const auto& left, const auto& right){
                return nums[left] < nums[right];
            }
        );
        auto beg_search_it = indexs.begin();
        auto end_search_it = std::upper_bound(indexs.begin(), indexs.end(), target - nums[indexs.front()],
                                              [&nums](const auto& left, const auto& right){
                                                  return left < nums[right];
                                              }
        );
        --end_search_it;
        while (end_search_it > beg_search_it) {
            int summ_ = nums[*beg_search_it] + nums[*end_search_it];
            if (summ_ == target) {
                result[0] = *beg_search_it;
                result[1] = *end_search_it;
                return result;
            } else if (summ_ < target) {
                ++beg_search_it;
            } else if (summ_ > target) {
                --end_search_it;
            }
        }
        return result;
    }
};

// my second solution
/*
class Solution {
public:
    std::vector<int> twoSum(std::vector<int>& nums, int target) {
        std::unordered_map<int , int> m;
        int index = 0;
        for (auto num: nums ) {
            auto iter = m.find(target - num);
            if (iter != m.end()) return {iter->second , index};
            m.emplace (num , index++);
        }
        return {};
    }
};
*/

void TestSolution();

int main() {
    TestRunner tr;
    RUN_TEST(tr, TestSolution);
    return 0;
}

void TestSolution() {
    {
        std::vector<int> nums{21, 17, 12, 25};
        std::vector<int> result{2, 1};
        Solution ss;
        ASSERT_EQUAL(ss.twoSum(nums, 29), result);
    }
    {
        std::vector<int> nums{0, 4, 3, 0};
        std::vector<int> result{0, 3};
        Solution ss;
        ASSERT_EQUAL(ss.twoSum(nums, 0), result);
    }
    {
        std::vector<int> nums{2, 7, 11, 15};
        std::vector<int> result{0, 1};
        Solution ss;
        ASSERT_EQUAL(ss.twoSum(nums, 9), result);
    }
    {
        std::vector<int> nums{3, 2, 4};
        std::vector<int> result{1, 2};
        Solution ss;
        ASSERT_EQUAL(ss.twoSum(nums, 6), result);
    }
    {
        std::vector<int> nums{-3, 4, 3, 90};
        std::vector<int> result{0, 2};
        Solution ss;
        ASSERT_EQUAL(ss.twoSum(nums, 0), result);
    }
}
