#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: find_all_numbers_disappeared_in_an_array_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import find_all_numbers_disappeared_in_an_array


class FindAllNumbersDisappearedInAnArrayTestCase(unittest.TestCase):
    def setUp(self):
        self.solution = find_all_numbers_disappeared_in_an_array.Solution()

    def tearDown(self):
        self.solution = None

    def test_example(self):
        self.assertItemsEqual(self.solution.findDisappearedNumbers([4,3,2,7,8,2,3,1]), [5, 6])

    def test_no_disappeared(self):
        self.assertItemsEqual(self.solution.findDisappearedNumbers([i for i in range(1, 100)]), [])

    def test_no_disappeared_reversed(self):
        self.assertItemsEqual(self.solution.findDisappearedNumbers([i for i in range(1, 100)][::-1]), [])
        
        
if __name__ == "__main__":
    unittest.main()


