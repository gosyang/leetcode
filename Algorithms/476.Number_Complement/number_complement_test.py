#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: number_complement_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import number_complement


class NumberComplementTestCase(unittest.TestCase):
    def setUp(self):
        self.solution = number_complement.Solution()

    def tearDown(self):
        self.solution = None

    def test_sample(self):
        self.assertEqual(self.solution.findComplement(5), 2)
        self.assertEqual(self.solution.findComplement(1), 0)
    
    def test_zero(self):
        self.assertEqual(self.solution.findComplement(0), 0)
    
    def test_max(self):
        self.assertEqual(self.solution.findComplement(2 ** 32 - 1), 0)

    def test_max_minuse_two(self):
        self.assertEqual(self.solution.findComplement(2 ** 32 - 3), 2)

if __name__ == "__main__":
    unittest.main()


