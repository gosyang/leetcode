/***************************************************************************
 * 
 * Copyright (c) 2016 , Inc. All Rights Reserved
 * 
 **************************************************************************/
 
 /**
 * @file Counting_Bits.cpp
 * @author baidu(baidu)
 * @date 2016/10/03 08:59:20
 * @brief 
 *  
 **/
#include <iostream>
#include <vector>
#include <math.h>

using namespace std;

vector<int> countBits(int num) {
    vector<int> res;
    int n = 2;
    for (int i = 0; i <= num; ++i) {
        if (i % n == 0 && i >= 4) {
            n = n * 2;
        }
        if (i == 0) {
            // 这种不对res[i] = 0;？
            res.push_back(0);
        } else if (i == 1) {
            res.push_back(1);
        } else {
            res.push_back(res[i - n] + 1);
        }
    }
    return res;
}

int main() {
    int n;
    cin >> n;
    vector<int> a = countBits(n);
    for (vector<int>::iterator it = a.begin(); it != a.end(); ++it) {
        cout << *it << endl;
    }
    return 0;
}
