#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: fizz_buzz_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import fizz_buzz


class FizzBuzzTestCase(unittest.TestCase):
    def setUp(self):
        self.solution = fizz_buzz.Solution()

    def tearDown(self):
        self.solution = None

    def test_example(self):
        output = [
            "1",
            "2",
            "Fizz",
            "4",
            "Buzz",
            "Fizz",
            "7",
            "8",
            "Fizz",
            "Buzz",
            "11",
            "Fizz",
            "13",
            "14",
            "FizzBuzz"
        ]
        self.assertItemsEqual(self.solution.fizzBuzz(15), output)
    
    def test_one(self):
        self.assertItemsEqual(self.solution.fizzBuzz(1), ["1"])
        
        
if __name__ == "__main__":
    unittest.main()


