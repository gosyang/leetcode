/***************************************************************************
 * 
 * Copyright (c) 2015, Inc. All Rights Reserved
 * 
 **************************************************************************/
 
 /**
 * @file Nim_Game.cpp
 * @author yanghanlin 
 * @date 2015/12/06 20:56:16
 * @brief 如果拿4必然拿不到，而如果每次让对手从4的倍数开始拿，可以保证你下次还让他继续
 *        如果扩展下，每次轮流拿n个，则每次让对手从n+1的倍数开始拿即可
 *  
 **/
#include <iostream>
bool canWinNim(int n) {
    if (n < 1) {
        return false;
    } else if (n % 4 == 0) {
        return false;
    } else {
        return true;
    }
}

int main() {
    int n;
    std::cin >> n;
    std::cout << canWinNim(n);
}


/* vim: set ts=4 sw=4 sts=4 tw=100 */
