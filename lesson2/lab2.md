# lab 2

## Selection

Using the following table

| Rate  | Earnings             |
| ----- | -------------------- |
| 10%   | $0 to $9,275         |
| 15%   | $9,276 to $37,650    |
| 25%   | $37,651 to $91,150   |
| 28%   | $91,151 to $190,150  |
| 33%   | $190,150 to $413,350 |
| 35%   | $413,351 to $415,050 |
| 39.6% | $415,051+            |

Create a flowchart for, and write a function for calculating your taxes two seperate ways.

1. Calculate where each earnings bracket is absolute. If you made 100,000, you paid 33% on everything.
2. Calculate where each bracket is PROGRESSIVE.
   IE you paid 10% in the first $9,275,
   15% on every dollar between $9,275 and 37,650,
   25% on every dollar between $37,651 and 91,150,
   etc.

## Repition

You have two bank accounts.
One has a return on investment of 100%. that is, each week, you doulbe the ammount in the account.
The second has a return of N dollars. That is, each week, N dollars is added to the account.

Assuming a random a random growth rate of the second account ($1000-$5000), and starting the first account at one penny ($0.01)
write a function that determines how many weeks it takes the first account to outpace the second

The expectation is that you Lookup and use the Math.random API from pkg.go.dev
