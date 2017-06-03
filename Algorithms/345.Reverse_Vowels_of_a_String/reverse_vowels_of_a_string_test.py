#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: reverse_vowels_of_a_string_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import reverse_vowels_of_a_string


class ReverseVowelsOfaStringTestCase(unittest.TestCase):
    def setUp(self):
        self.solution = reverse_vowels_of_a_string.Solution()

    def tearDown(self):
        self.solution = None

    def test_example(self):
        self.assertEqual(self.solution.reverseVowels("hello"), "holle")
        self.assertEqual(self.solution.reverseVowels("leetcode"), "leotcede")
    
    def test_empty(self):
        self.assertEqual(self.solution.reverseVowels(""), "")
    
    def test_one_blank(self):
        self.assertEqual(self.solution.reverseVowels(" "), " ")

    def test_no_volwes(self):
        self.assertEqual(self.solution.reverseVowels("qwrty"), "qwrty")
    
    def test_not_volwes_end(self):
        self.assertEqual(self.solution.reverseVowels("A."), "A.")
        
        
if __name__ == "__main__":
    unittest.main()


