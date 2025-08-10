# Boyer-Moore Voting algorithm
## Algorithm:
1. Begin by initializing two variables, num and a counter count, both of which are set to 0.

2. Now we’ll begin iterating over the number and for each element x.

3. We first check to see if the count is 0, and then we assign num to x.

4. Then we check to see if the current num is equal to x, if not, we decrease the count by one, else we increment it by one.

5. Reset the counter to zero.

6. Find the frequency of the num variable by looping through the nums array.

7. Now, if the count is larger than N/2, we return num; otherwise we return -1.
