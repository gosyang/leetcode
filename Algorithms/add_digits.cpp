/***************************************************************************
 * 
 * Copyright (c) 2015, Inc. All Rights Reserved
 * 
 **************************************************************************/
 
 /**
 * @file add_digits.cpp
 * @author yanghanlin 
 * @date 2015/12/06 20:56:16
 * @brief 一下子想不到方法就使用列举法，一个一个列出来，123456789，10，11，12
 *        然后发现，从10，11 开始了 %9 的循环，特殊处理0
 *  
 **/
#include <iostream>

int addDigits(int num) {
    if (num == 0) {
        return 0;
    } else if (num % 9 == 0) {
        return 9;
    } else {
        return num % 9;
    }
}

int main() {
    int n;
    std::cin >> n;
    std::cout << addDigits(n);
}


/* vim: set ts=4 sw=4 sts=4 tw=100 */
