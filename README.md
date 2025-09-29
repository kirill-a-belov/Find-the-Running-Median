"# Find the Running Median in Go

**Find-the-Running-Median** is a Go implementation of an efficient algorithm to compute the median of a dynamically growing data stream.  
This project is designed for educational purposes, algorithm research, and performance benchmarking.

---

## Project Overview

Calculating a running median is a common challenge in data analysis, real-time statistics, and software engineering interviews.  
A **running median** is the median of a stream of numbers, updated with each new element. The stream can either grow continuously or operate as a sliding window.

### Naive Approach

The straightforward solution is to store all numbers in a sorted array and recalculate the median on each update.
- Works correctly, but **inefficient for large datasets**.
- Benchmarks show performance differences up to **400x slower** than optimal solutions.

### Efficient Algorithm

This project implements the **canonical solution using two heaps**:
- **Max-heap** — stores the smaller half of numbers
- **Min-heap** — stores the larger half of numbers

Algorithm steps:
1. Initialize two empty heaps. Insert the first number into any heap; it becomes the current median.
2. For the second number, insert into the other heap and balance roots if necessary. Median is the average of roots.
3. For subsequent numbers:
    - If number ≥ current median → insert into min-heap
    - Else → insert into max-heap
    - Balance heaps if size difference ≥ 2
    - Update median:
        - Equal heap sizes → median = average of roots
        - Unequal heap sizes → median = root of the larger heap

This algorithm ensures **O(log n)** insertion and median calculation per update, suitable for **real-time streaming applications**.

---

## Getting Started

Clone the repository:

```bash
git clone https://github.com/kirill-a-belov/Find-the-Running-Median.git
cd Find-the-Running-Median
```
## Status

Open-source and research-oriented.

Useful for algorithm learning, real-time data analysis, and performance optimization studies. 

Contributions and forks are welcome. 

## Tags

#Go #algorithms #running_median #data_stream #real_time"
