#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: reverse_string_II_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import reverse_string_II


class ReverseStringIITestCase(unittest.TestCase):
    def setUp(self):
        self.solution = reverse_string_II.Solution()

    def tearDown(self):
        self.solution = None

    def test_example(self):
        self.assertEqual(self.solution.reverseStr("abcdefg", 2), "bacdfeg")
    
    def test_less_than_k(self):
        self.assertEqual(self.solution.reverseStr("abc", 4), "cba")
    
    def test_more_than_k_less_than_2k(self):
        self.assertEqual(self.solution.reverseStr("abcde", 3), "cbade")
    
    def test_2k(self):
        self.assertEqual(self.solution.reverseStr("abcd", 2), "bacd")

    def test_k(self):
        self.assertEqual(self.solution.reverseStr("abcd", 4), "dcba")
        
        
if __name__ == "__main__":
    unittest.main()


