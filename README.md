## Find the Running Median with Golang ##
 
One of quite tricky questions on Software Developers position interview - finding a running median.

It could be described with various variants but the main conditions of the task is:
1. You have a numbers stream wich consistantly grow or it could be a sliding window (ex. for each iteration add one number from right and delete another from left).
2. You should find a median of this stream and finding it on each iteration.

The obvious solution is saving all this data in the sorted array (maybe - sort it on each iteration or, for better performance, keep it sort from the start). It's quite easy way and it works. But this way is quite slow (i did some benchmarks - performance difference about 400 times) , your solution doesn't pass performance test.

There is a canonical way how to solve this codding challange! 
We should use two heaps (maximum and minimum) to keep two halfs of our input numbers. The MAXIMUM half of input numbers we keep in MINIMAL heap (in the root we always will have minimal number from maximum half). The MINIMUM half of input numbers we keep in MAXIMAL heap (in the root we always will have maximal number from minimal half). 

>### The algorithm:
>First iteration. We have two empty heaps. Add number in any heap and set current median equal to this number.

>Second iteration. Add new number in another heap. If the root of min heap less then root of max heap - changing roots of heaps. Set current median as average of roots.

>Next iterations:
>1. If new input number bigger (or equal) than current median - add it in minimal heap. If not - in maximal heap.
>2. Check if one heap elements number more than another on 2(and more) - get element from bigger heap and add in to less heap.
>3. Update the current median. (If heaps has equal size - median is summ of their roots devided by 2. Another case - median is the root of minimal heap)
