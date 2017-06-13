#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: number_ofboomeranges_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import number_of_boomeranges


class NumberOfBoomerangesTestCase(unittest.TestCase):
    def setUp(self):
        self.solution = number_of_boomeranges.Solution()

    def tearDown(self):
        self.solution = None

    def test_example(self):
        self.assertEqual(self.solution.numberOfBoomerangs([[0, 0], [1, 0], [2, 0]]), 2)

    def test_empty(self):
        self.assertEqual(self.solution.numberOfBoomerangs([]), 0)

    def test_single_one(self):
        self.assertEqual(self.solution.numberOfBoomerangs([[-10000, 10000]]), 0)
        
        
if __name__ == "__main__":
    unittest.main()


