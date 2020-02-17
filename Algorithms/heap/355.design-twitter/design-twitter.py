#!/usr/bin/env python
# -*- coding: utf-8 -*-
# 
# Copyright (c) 2017 gosyang. All Rights Reserved
# 
"""
File: design-twitter
Author: gosyang
Date: 2020/2/17 10:48 AM
设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，能够看见关注人（包括自己）的最近十条推文。你的设计需要支持以下的几个功能：

postTweet(userId, tweetId): 创建一条新的推文
getNewsFeed(userId): 检索最近的十条推文。每个推文都必须是由此用户关注的人或者是用户自己发出的。推文必须按照时间顺序由最近的开始排序。
follow(followerId, followeeId): 关注一个用户
unfollow(followerId, followeeId): 取消关注一个用户
"""
import heapq

class Twitter(object):

    def __init__(self):
        """
        Initialize your data structure here.
        """
        # 全局的tweet id，每次递增来排序
        self.id = 0
        # key是关注者，value是set，放关注的人
        self.user_followers = {}
        # key是发布者id，value是[(allId, tweet_id)]
        self.user_tweets = {}

    def postTweet(self, userId, tweetId):
        """
        Compose a new tweet.
        :type userId: int
        :type tweetId: int
        :rtype: None
        """
        if userId not in self.user_followers:
            self.user_followers[userId] = set([userId])
        self.id += 1
        if userId not in self.user_tweets:
            self.user_tweets[userId] = []
        # 这里必须注意，list相加必须是 += []
        self.user_tweets[userId] += [(self.id, tweetId)]

    def getNewsFeed(self, userId):
        """
        Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent.
        :type userId: int
        :rtype: List[int]
        """
        if userId not in self.user_followers:
            return []
        q = []
        for followee in self.user_followers[userId]:
            # 这里要注意如果这个人没有发过推，那不能直接取
            if followee not in self.user_tweets:
                continue
            # 这里是从小到大，取最大的10个要倒着来
            q += self.user_tweets[followee][-10:]
        # TODO 这里要书记，key的用法哦
        return [t[1] for t in heapq.nlargest(10, q, key=lambda s: s[0])]


    def follow(self, followerId, followeeId):
        """
        Follower follows a followee. If the operation is invalid, it should be a no-op.
        :type followerId: int
        :type followeeId: int
        :rtype: None
        """
        # 这里要注意follow别人要注意key
        if followerId not in self.user_followers:
            self.user_followers[followerId] = set([followerId])
        self.user_followers[followerId].add(followeeId)

    def unfollow(self, followerId, followeeId):
        """
        Follower unfollows a followee. If the operation is invalid, it should be a no-op.
        :type followerId: int
        :type followeeId: int
        :rtype: None
        """
        # TODO 这里要注意dict的key不存在的情况都得判断，另外自己删自己删不掉
        if followerId not in self.user_followers or followerId == followeeId:
            return
        if followeeId in self.user_followers[followerId]:
            self.user_followers[followerId].remove(followeeId)


# Your Twitter object will be instantiated and called as such:
# obj = Twitter()
# obj.postTweet(userId,tweetId)
# param_2 = obj.getNewsFeed(userId)
# obj.follow(followerId,followeeId)
# obj.unfollow(followerId,followeeId)
