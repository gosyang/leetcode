#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: binary_watch_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import binary_watch


class BinaryWatchTestCase(unittest.TestCase):
    def setUp(self):
        self.solution = binary_watch.Solution()

    def tearDown(self):
        self.solution = None
        
    def test_bitCount(self):
        self.assertEqual(self.solution.bitCount(0), 0)
        self.assertEqual(self.solution.bitCount(1), 1)
        self.assertEqual(self.solution.bitCount(2), 1)
        self.assertEqual(self.solution.bitCount(3), 2)
        self.assertEqual(self.solution.bitCount(4), 1)
        self.assertEqual(self.solution.bitCount(5), 2)
        self.assertEqual(self.solution.bitCount(6), 2)
        self.assertEqual(self.solution.bitCount(7), 3)
        self.assertEqual(self.solution.bitCount(8), 1)

    def test_input_zero(self):
        self.assertEqual(self.solution.readBinaryWatch(0), ["0:00"])
    
    def test_input_bigger_than_all_led(self):
        self.assertEqual(self.solution.readBinaryWatch(9), [])
    
    def test_sample(self):
        sample = ["1:00", "2:00", "4:00", "8:00", "0:01",
                  "0:02", "0:04", "0:08", "0:16", "0:32"]
        self.assertItemsEqual(self.solution.readBinaryWatch(1), sample)
        # 这里不能用 assertListEqual 因为里面的顺序不同, 直接 == 比较也是false
        
        
if __name__ == "__main__":
    unittest.main()


