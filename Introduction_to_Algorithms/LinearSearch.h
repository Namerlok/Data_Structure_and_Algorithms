#pragma once

template <typename InputIt, typename T>
InputIt LinearSearch(InputIt first, InputIt last, const T& value) {
    for (InputIt it = first; it != last; ++it)
        if (*it == value)
            return it;
    return last;
}
