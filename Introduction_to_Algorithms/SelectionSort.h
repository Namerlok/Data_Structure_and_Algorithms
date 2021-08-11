#pragma once

#include <iterator>
#include <utility>
#include <iostream>

template <typename RandomIt>
void SelectionSort(RandomIt beg, RandomIt end) {
    for (auto it = beg; std::next(it) != end; ++it) {
        auto min_el = std::move_if_noexcept(*it);
        RandomIt sh_min_it = std::next(it);
        while (sh_min_it != end) {
            if (*sh_min_it < min_el)
                std::swap(*sh_min_it, min_el);
            ++sh_min_it;
        }
        *it = std::move_if_noexcept(min_el);
    }
}

template <typename RandomIt, typename Compare>
void SelectionSort(RandomIt beg, RandomIt end, Compare cmp) {
    for (auto it = beg; std::next(it) != end; ++it) {
        auto min_el = std::move_if_noexcept(*it);
        RandomIt sh_min_it = std::next(it);
        while (sh_min_it != end) {
            if (cmp(*sh_min_it, min_el))
                std::swap(*sh_min_it, min_el);
            ++sh_min_it;
        }
        *it = std::move_if_noexcept(min_el);
    }
}

