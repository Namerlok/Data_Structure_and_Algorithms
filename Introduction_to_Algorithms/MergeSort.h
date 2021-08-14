#pragma once

#include <iostream>
#include <utility>
#include <vector>

template <typename RandomIt>
void MergingSortedConteiners(RandomIt beg_1, RandomIt end_1,
                             RandomIt beg_2, RandomIt end_2) {
    size_t size_1 = end_1 - beg_1, counter_1 = 0;
    size_t size_2 = end_2 - beg_2, counter_2 = 0;
    std::vector<typename RandomIt::value_type>
            con_1(std::make_move_iterator(beg_1),
                  std::make_move_iterator(end_1));
    std::vector<typename RandomIt::value_type>
            con_2(std::make_move_iterator(beg_2),
                  std::make_move_iterator(end_2));

    for (RandomIt it = beg_1; it < end_2; ++it) {
        if ((counter_2 == size_2 && counter_1 < size_1) ||
            (counter_1 < size_1 && con_1[counter_1] <= con_2[counter_2]))
            *it = std::move(con_1[counter_1++]);
        else if (counter_2 < size_2)
            *it = std::move(con_2[counter_2++]);
        else
            std::cerr << "Error: no conditions selected\n";
    }
}

template <typename RandomIt>
void MergeSort(RandomIt beg, RandomIt end) {
    size_t size = end - beg;
    if (size > 1) {
        MergeSort(beg, beg + size / 2);
        MergeSort(beg + size / 2, end);
        MergingSortedConteiners(beg, beg + size / 2,
                                beg + size / 2, end);
    }
}

