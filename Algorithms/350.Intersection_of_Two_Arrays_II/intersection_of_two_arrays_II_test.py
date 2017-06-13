#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: intersection_of_two_arrays_II.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import intersection_of_two_arrays_II


class IntersectionOfTwoArrayIITestCase(unittest.TestCase):
    def setUp(self):
        self.solution = intersection_of_two_arrays_II.Solution()

    def tearDown(self):
        self.solution = None
        
    def test_sample(self):
        self.assertItemsEqual(self.solution.intersect([1, 2, 2, 1], [2, 2]), [2, 2])

    def test_nums1_less_then_nums2(self):
        self.assertItemsEqual(self.solution.intersect([1], [1, 1]), [1])
    
    def test_same_number_unorder(self):
        self.assertItemsEqual(self.solution.intersect([2, 1], [1, 2]), [1, 2])
        
if __name__ == "__main__":
    unittest.main()


