#pragma once

#include <iterator>
#include <utility>
#include <iostream>

template <typename RandomIt>
void InsertionSort(RandomIt beg, RandomIt end) {
    for (auto it = beg; it != end; ++it) {
        auto swap_val = std::move_if_noexcept(*it);
        RandomIt sh_it = std::prev(it);
        while (std::next(sh_it) != beg && *sh_it > swap_val) {
            *std::next(sh_it) = std::move_if_noexcept(*sh_it);
            --sh_it;
        }
        *std::next(sh_it) = std::move_if_noexcept(swap_val);
    }
}

template <typename RandomIt, typename Compare>
void InsertionSort(RandomIt beg, RandomIt end, Compare cmp) {
    for (auto it = beg; it != end; ++it) {
        auto swap_val = std::move_if_noexcept(*it);
        RandomIt sh_it = std::prev(it);
        while (std::next(sh_it) != beg && !cmp(*sh_it, swap_val)) {
            *std::next(sh_it) = std::move_if_noexcept(*sh_it);
            --sh_it;
        }
        *std::next(sh_it) = std::move_if_noexcept(swap_val);
    }
}

