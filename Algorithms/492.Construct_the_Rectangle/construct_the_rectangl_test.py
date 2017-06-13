#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: construct_the_rectangle_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import construct_the_rectangle


class ConstructTheRectangleTestCase(unittest.TestCase):
    def setUp(self):
        self.solution = construct_the_rectangle.Solution()

    def tearDown(self):
        self.solution = None

    def test_9(self):
        self.assertEqual(self.solution.constructRectangle(9), [3, 3])

    def test_2(self):
        # 这个需要顺序对,所以不能 itemequal
        self.assertEqual(self.solution.constructRectangle(2), [2, 1])
        
        
if __name__ == "__main__":
    unittest.main()


