"""  
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

 

Example 1:

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
"""


class Solution:
    # just sort the intervals and keep track of the latest one that's appended
    def merge(self, intervals: list[list[int]]) -> list[list[int]]:
        intervals.sort(key=lambda x: x[0])

        new_array = [intervals[0]]

        for i in range(1, len(intervals)):
            if intervals[i][0] <= new_array[-1][1]:
                new_array[-1][1] = max(intervals[i][1], new_array[-1][1])
            else:
                new_array.append(intervals[i])

        return new_array
