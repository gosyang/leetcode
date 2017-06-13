#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: distribute_candies_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import distribute_candies


class DistributeCandiesTestCase(unittest.TestCase):
    def setUp(self):
        self.solution = distribute_candies.Solution()

    def tearDown(self):
        self.solution = None

    def test_example(self):
        self.assertEqual(self.solution.distributeCandies([1, 1, 2, 2, 3, 3]), 3)
        self.assertEqual(self.solution.distributeCandies([1, 1, 2, 3]), 2)
        
        
if __name__ == "__main__":
    unittest.main()


