#pragma once

template <typename RandIt, typename T>
RandIt BinarySearch(RandIt first, RandIt last, const T& value) {
    std::size_t size = last - first;
    RandIt result;
    if (size <= 1) {
         return first;
    } else {
        T middle = (size) / 2;
        if (*(first + middle) > value)
            result = BinarySearch(first, first + middle, value);
        else
            result = BinarySearch(first + middle, last, value);
    }

    if (*result != value)
        return last;
    else
        return result;
}
