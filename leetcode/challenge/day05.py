"""
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).

Input: [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
             Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
"""
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        total = 0
        if len(prices) < 2:
            return 0
        prior = prices[0]
        increase = False
        for i in range(1, len(prices)):
            if prices[i] > prior:
                prices = prices[i-1:]
                increase = True
                break
            prior = prices[i]
        if not increase:
            return 0
        return maxProfitSub(prices)

        # there is at least an increase in the list from the first entry

def maxProfitSub(prices: List[int]) -> int:
    length = len(prices)
    lowest, highest = prices[0], prices[1]
    profit = highest - lowest
    for i in range(2, length):
        n = prices[i]
        if n >= highest:
            highest = n
            profit = highest - lowest
        else:
            if length - i < 2:
                profit = highest - lowest
                break
            if prices[i + 1] < n:
                continue
            profit += maxProfitSub(prices[i:])
            break

    return profit