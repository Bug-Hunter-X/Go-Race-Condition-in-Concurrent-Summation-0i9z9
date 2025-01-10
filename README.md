# Go Race Condition Example

This repository demonstrates a race condition in a Go program that calculates the sum of numbers concurrently.  The program uses goroutines and a mutex for synchronization, but the mutex's implementation could lead to unexpected results.  The solution showcases the correct way to handle concurrency and prevent data races.

## Description

The program calculates the sum of numbers from 0 to 999 using multiple goroutines.  While it attempts to use a `sync.Mutex` to protect the shared `count` variable, there's a subtle race condition that might cause the final sum to be incorrect.  This is because of improper handling of the mutex in the goroutine's scope. 

## Solution

The solution addresses the race condition by ensuring correct usage and scope of the mutex. This ensures that only one goroutine can modify `count` at a time, preventing data races and yielding the correct result.
