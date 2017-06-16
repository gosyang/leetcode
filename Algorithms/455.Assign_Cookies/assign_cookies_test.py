#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: assign_cookies_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import assign_cookies


class AssignCookiesTestCase(unittest.TestCase):
    def setUp(self):
        self.solution = assign_cookies.Solution()

    def tearDown(self):
        self.solution = None

    def test_example(self):
        self.assertEqual(self.solution.findContentChildren([1, 2, 3], [1, 1]), 1)
        self.assertEqual(self.solution.findContentChildren([1, 2], [1, 2, 3]), 2)

    def test_child_none(self):
        self.assertEqual(self.solution.findContentChildren([], []), 0)
        
    def test_child_more(self):
        self.assertEqual(self.solution.findContentChildren([1, 1, 1], [1]), 1)
        
    def test_one_cookie_small(self):
        self.assertEqual(self.solution.findContentChildren([2, 2, 2], [1]), 0)
        
    def test_cookie_more(self):
        self.assertEqual(self.solution.findContentChildren([2, 2], [3, 4, 5]), 2)
    
if __name__ == "__main__":
    unittest.main()


