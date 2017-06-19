#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: ransom_note_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import ransom_note


class RansomNoteTestCase(unittest.TestCase):
    def setUp(self):
        self.solution = ransom_note.Solution()

    def tearDown(self):
        self.solution = None

    def test_example(self):
        self.assertFalse(self.solution.canConstruct("a", "b"))
        self.assertFalse(self.solution.canConstruct("aa", "ab"))
        self.assertTrue(self.solution.canConstruct("aa", "aab"))
        
        
if __name__ == "__main__":
    unittest.main()


