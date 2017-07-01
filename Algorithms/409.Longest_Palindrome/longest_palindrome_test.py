#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: longest_palindrome_test.py
Author: gosyang
Date: 2017/06/29 16:10:09
"""
import unittest

import longest_palindrome


class LongestPalindromeTestCase(unittest.TestCase):
    def setUp(self):
        self.solution = longest_palindrome.Solution()

    def tearDown(self):
        self.solution = None

    def test_example(self):
        self.assertEqual(self.solution.longestPalindrome("abccccdd"), 7)
    
    def test_empty(self):
        self.assertEqual(self.solution.longestPalindrome(""), 0)

    def test_Aa(self):
        self.assertEqual(self.solution.longestPalindrome("Aa"), 1)
        
    def test_abcABC(self):
        self.assertEqual(self.solution.longestPalindrome("abcABC"), 1)

    def test_long(self):
        self.assertEqual(self.solution.longestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"), 983)
        
if __name__ == "__main__":
    unittest.main()


