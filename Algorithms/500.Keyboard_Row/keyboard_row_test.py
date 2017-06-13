#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: keyboard_row_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import keyboard_row


class KeyboardRowTestCase(unittest.TestCase):
    def setUp(self):
        self.solution = keyboard_row.Solution()

    def tearDown(self):
        self.solution = None

    def test_example(self):
        input = ["Hello", "Alaska", "Dad", "Peace"]
        output = ["Alaska", "Dad"]
        self.assertItemsEqual(self.solution.findWords(input), output)
        
        
if __name__ == "__main__":
    unittest.main()


