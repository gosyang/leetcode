#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: complex_number_multiplication_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import complex_number_multiplication


class ComplexNumberMultiplicationTestCase(unittest.TestCase):
    def setUp(self):
        self.solution = complex_number_multiplication.Solution()

    def tearDown(self):
        self.solution = None

    def test_example(self):
        self.assertEqual(self.solution.complexNumberMultiply("1+1i", "1+1i"), "0+2i")
        self.assertEqual(self.solution.complexNumberMultiply("1+-1i", "1+-1i"), "0+-2i")

    def test_chunxu(self):
        self.assertEqual(self.solution.complexNumberMultiply("0+1i", "0+1i"), "-1+0i")
        
    def test_chunshi(self):
        self.assertEqual(self.solution.complexNumberMultiply("100+0i", "-100+0i"), "-10000+0i")
        
        
if __name__ == "__main__":
    unittest.main()


