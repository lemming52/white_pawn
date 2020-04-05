"""
A chef has collected data on the satisfaction level of his n dishes. Chef can cook any dish in 1 unit of time.

Like-time coefficient of a dish is defined as the time taken to cook that dish including previous dishes multiplied by its satisfaction level  i.e.  time[i]*satisfaction[i]

Return the maximum sum of Like-time coefficient that the chef can obtain after dishes preparation.

Dishes can be prepared in any order and the chef can discard some dishes to get this maximum value.

Input: satisfaction = [-1,-8,0,5,-9]
n == satisfaction.length
Output: 14
Explanation: After Removing the second and last dish,
the maximum total Like-time coefficient will be equal to (-1*1 + 0*2 + 5*3 = 14).
Each dish is prepared in one unit of time.

"""

class Solution:
    def maxSatisfaction(self, satisfaction: List[int]) -> int:
        satisfaction = sorted(satisfaction)
        coefficient = 0
        satTotal = 0
        for i in range(len(satisfaction) - 1, -1, -1):
            val = satisfaction[i]
            if val >= 0:
                satTotal += val
            else:
                if val + satTotal > 0:
                    satTotal += val
                else:
                    satisfaction = satisfaction[i+1:]
                    break
        result = 0
        for i, n in enumerate(satisfaction):
            result += (i + 1) * n
        return result
