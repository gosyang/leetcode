#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: reshape_the_matrix_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import reshape_the_matrix


class ReshapeTheMatrixTestCase(unittest.TestCase):
    def setUp(self):
        self.solution = reshape_the_matrix.Solution()

    def tearDown(self):
        self.solution = None

    def test_sample(self):
        nums = [[1, 2], [3, 4]]
        r = 1
        c = 4
        shaped_nums = [[1, 2, 3, 4]]
        self.assertItemsEqual(self.solution.matrixReshape(nums, r, c), shaped_nums)
    
    def test_sample_impossible(self):
        nums = [[1, 2], [3, 4]]
        r = 2
        c = 4
        shaped_nums = [[1, 2], [3, 4]]
        self.assertItemsEqual(self.solution.matrixReshape(nums, r, c), shaped_nums)
        
        
if __name__ == "__main__":
    unittest.main()


