#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: longest_uncommon_subsequence_II_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import longest_uncommon_subsequence_II


class LongestUncommonSubsequenceIITestCase(unittest.TestCase):
    def setUp(self):
        self.solution = longest_uncommon_subsequence_II.Solution()

    def tearDown(self):
        self.solution = None

    def test_example(self):
        self.assertEqual(self.solution.findLUSlength(["aba", "cdc", "eae"]), 3)

    def test_longest_same(self):
        self.assertEqual(self.solution.findLUSlength(["aba", "aba", "c"]), 1)

    def test_longest_same_and_short_sub(self):
        self.assertEqual(self.solution.findLUSlength(["aba", "aba", "a"]), -1)

    def test_longest_same_and_short_not_sub(self):
        """
        没想到的case, 主要对子序列的理解有错误, 只要在里面符合顺序不一定要挨着
        """
        self.assertEqual(self.solution.findLUSlength(["aabbcc", "aabbcc", "cb", "abc"]), 2)

    def test_same_length_same_item(self):
        """
        没想到的case, 没想好多个相同的长度但是不同内容
        """
        self.assertEqual(self.solution.findLUSlength(["a","b","c","d","e","f","a","b","c","d","e","f"]), -1)

    def test_same_length_one_different(self):
        """
        没想到的case
        """
        self.assertEqual(self.solution.findLUSlength(["a","b","c","d","e","f","a","b","c","d","e"]), 1)
    
    def test_same_in_group(self):
        strs = ["mibqsxxnne", "mibqsxxnne", "mibqsxxnne", "mibqsxxnne", "mibqsxxnne", "qxfwgjpyuk",
         "qxfwgjpyuk", "qxfwgjpyuk", "qxfwgjpyuk", "qxfwgjpyuk", "mdobdmiamk", "mdobdmiamk",
         "mdobdmiamk", "mdobdmiamk", "mdobdmiamk", "suvrrybhfz", "suvrrybhfz", "suvrrybhfz",
         "suvrrybhfz", "suvrrybhfz", "wmughoitpp", "wmughoitpp", "wmughoitpp", "wmughoitpp",
         "wmughoitpp", "mgjsizloja", "mgjsizloja", "mgjsizloja", "mgjsizloja", "mgjsizloja",
         "hhlefbuuwk", "hhlefbuuwk", "hhlefbuuwk", "hhlefbuuwk", "hhlefbuuwk", "uheyrvakww",
         "uheyrvakww", "uheyrvakww", "uheyrvakww", "uheyrvakww", "yfrugcirau", "yfrugcirau",
         "yfrugcirau", "yfrugcirau", "yfrugcirau", "dnapgcaipk", "dnapgcaipk", "dnapgcaipk",
         "dnapgcaipk", "dnapgcaipk"]

        self.assertEqual(self.solution.findLUSlength(strs), -1)
        
if __name__ == "__main__":
    unittest.main()


