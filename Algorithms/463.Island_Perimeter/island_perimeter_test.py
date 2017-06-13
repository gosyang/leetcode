#!/usr/bin/env python
# -*- coding: utf-8 -*-
########################################################################
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
########################################################################
 
"""
File: island_perimeter_test.py
Author: gosyang
Date: 2017/05/15 16:10:09
"""
import unittest

import island_perimeter


class IslandPerimeterTestCase(unittest.TestCase):
    def setUp(self):
        self.solution = island_perimeter.Solution()

    def tearDown(self):
        self.solution = None

    def test_example(self):
        grid = [[0, 1, 0, 0],
                [1, 1, 1, 0],
                [0, 1, 0, 0],
                [1, 1, 0, 0]]

        self.assertEqual(self.solution.islandPerimeter(grid), 16)
    
    def test_empty(self):
        grid = []
        self.assertEqual(self.solution.islandPerimeter(grid), 0)

    def test_empty_empty(self):
        grid = [[]]
        self.assertEqual(self.solution.islandPerimeter(grid), 0)
        
    def test_wrong_case1(self):
        grid = [[0],
                [1],
                [1]]
        self.assertEqual(self.solution.islandPerimeter(grid), 6)
        
if __name__ == "__main__":
    unittest.main()


