"""
Given an array of strings, group anagrams together.

Example:

Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
Output:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
Note:

All inputs will be in lowercase.
The order of your output does not matter.
"""

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        hashes = {}
        for s in strs:
            sHash = tuple(sorted(s))
            if sHash in hashes:
                hashes[sHash].append(s)
            else:
                hashes[sHash] = [s]
        return [v for k, v in hashes.items()]