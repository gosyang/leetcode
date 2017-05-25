#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 Baidu.com, Inc. All Rights Reserved
# 
########################################################################
 
"""
File: binary_tree_tilt_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import binary_tree_tilt


class BinaryTreeTiltTestCase(unittest.TestCase):
    def setUp(self):
        self.solution = binary_tree_tilt.Solution()

    def tearDown(self):
        self.solution = None

    def test_example(self):
        """
            1
          /   \
         2     3
        """
        root = binary_tree_tilt.TreeNode(1)
        left = binary_tree_tilt.TreeNode(2)
        right = binary_tree_tilt.TreeNode(3)
        root.left = left
        root.right = right
        self.assertEqual(self.solution.findTilt(root), 1)
        
        
if __name__ == "__main__":
    unittest.main()


