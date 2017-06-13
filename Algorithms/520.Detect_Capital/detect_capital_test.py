#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: detect_capital_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import detect_capital


class DetectCapitalTestCase(unittest.TestCase):
    def setUp(self):
        self.solution = detect_capital.Solution()

    def tearDown(self):
        self.solution = None

    def test_one_lower(self):
        self.assertTrue(self.solution.detectCapitalUse("a"))

    def test_one_upper(self):
        self.assertTrue(self.solution.detectCapitalUse("Z"))

    def test_all_upper(self):
        self.assertTrue(self.solution.detectCapitalUse("USA"))

    def test_all_lower(self):
        self.assertTrue(self.solution.detectCapitalUse("sb"))

    def test_all_first_upper(self):
        self.assertTrue(self.solution.detectCapitalUse("Good"))

    def test_all_wrong(self):
        self.assertFalse(self.solution.detectCapitalUse("GooG"))
        
        
if __name__ == "__main__":
    unittest.main()


