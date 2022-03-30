# Deeper understanding of structures

There are really only a few fundamental structures that build a program.
Any program you can think of that can be solved can be built using these structures.
describing these structures is best done in flowchart.

Flow diagrams/charts, and how to build a program

## Types of structures

- sequence (step by step)
  - doing one step after the next.
  - set var a to 10,
  - add some input to var a,
  - divide var a by 2.
- selection (if statements)
  - if else blocks
  - else if chaning
  - final if block
- reptition (for loops)
  - loops
  - infinite loops
  - range over loops
  - break statements and how to exit a loop.

## complecations

- stacking structures
  - using a for loop withing a foor loop
  - using an if statement within a loop
- merging structures
  - on path 1 in an if, do one for loop
  - on path 2 in an if, do a different for loop, or set of actions
- recombining
  - when done with an action, return to a specific point (p51) (explain 2-22 vs 2-29)

## notes about structures

- priming input (take in first input before the loop)
  - watch for not taking inputs at the right step.
  - taking inputs more than once, or never again after the first time.

## switch statements

- chains of if statements
- allows for selecting one of many different options with ease.
  - cases
  - break
  - default

# boolean expressions & decision making

## Comparison operators

- `<` less than
- `<=` less than or equal to
- `==` equal to
- `>=` greater than or equal to
- `>` greater than
- `!=` (sometimes <>) not equal to

## combining logic

- && logical AND operator
  - returns true if both sides of the expression are true.
  - I will be happy if I get eggs AND bacon
- || logical OR operator
  - returns true if atleast one side of the expression is true
  - I would be happy if I won a Car OR I won a million dollars.
    - note that I would be happy if I won the car and the million dollars.
- ! logical Not operator
  - says that the logic on the right is inverted.
  - I will be happy if my breakfeast does NOT contain eggs.
- Parens and chaning
  - It is possible to use multiple logical operators in a row.
  - I would be happy with (Eggs && bacon) || (pancakes && mapel syurp)

## Common errors

Each comparison in an operator bust be done every time
`I >0 AND <100` is invalid. The operator doesent know what the second side is asking for.
`I >0 AND I<100`

## Presidence

golang operator presidence,

- Not operator
- comparisons
- AND operator
- OR operator
  https://www.tutorialspoint.com/go/go_operators_precedence.htm

# APIs (Application Programming Interface) and packages

Code repition is normally bad.
common code is often combined into groupings of functions and functions often groupded into packages.
pkg.go.dev

go into details on

- strings
- math
- strconv
- fmt
