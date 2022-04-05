# Logic Gates and Transistors

Logic gates and transistors are at the heart of how a computer works. I followed this [Ben Eater video](https://www.youtube.com/watch?v=sTu3LwpF6XI) while writing the code.

## Transistors
Transistors are the fundamental building block of computers. They take in two inputs. If current is flowing through both of those inputs it lets the currents pass through. Otherwise, it blocks the currents. Current is often represented with a 1 for current and a 0 for no current. (1 and 0 are some times replaced with true and false. I implemented the code using true and false.)

## Gates
---
### And Gates
And gates are made by running the first input and 1 through the transistor, and then running the second input and the output from the previous transistor into another transistor. The output of the second transistor is the result of the and gate. They have the following truth table:

| Input 1 | Input 2 | Output |
|---------|---------|--------|
|       0 |       0 |      0 |
|       1 |       0 |      0 |
|       0 |       1 |      0 |
|       1 |       1 |      0 |


### Not Gate
Not gates take in a single input and invert it. Ben Eater created one by having an on current that could go through the transistor, or around the transistor. When the input if fed into the other input for the transistor, if the input is a 1 or true, current will flow through the transistor instead of around it because flowing through the transistor is shorter than bypassing it. If the input is a 0 or false, no current will flow through the transistor so the transistor will be bypassed. The not gate has the following truth table:

| Input | Output |
|-------|--------|
|     1 |      0 |
|     0 |      1 |


### Or Gate
In an or gate, if either or both input1 and input2 are 1, then the output is 1. This can be made using two transistors in parallel. Input1 and 1 are run through the first transistor. Input2 and 1 are run through the second transistor. The outputs of both transistors are combined into the final output. Or has the following truth table:

| Input 1 | Input 2 | Output |
|---------|---------|--------|
|       0 |       0 |      0 |
|       1 |       0 |      1 |
|       0 |       1 |      1 |
|       1 |       1 |      1 |

### Nor Gate
A Nor Gate is simply the output of an or gate passed through a not gate. Nor has the following truth table:

| Input 1 | Input 2 | Output |
|---------|---------|--------|
|       0 |       0 |      1 |
|       1 |       0 |      0 |
|       0 |       1 |      0 |
|       1 |       1 |      0 |

### Xor Gate
A xor gate is 1 if input1 or input2 is 1, but not if both input1 and input2 are 1. This gate can be made by passing both inputs through and or gate, passing both inputs through an and gate, and then passing the inverted result of the and gate and the result of the or gate through a final and gate. The output is the result of the final and gate. The truth table for xor is:

| Input 1 | Input 2 | Output |
|---------|---------|--------|
|       0 |       0 |      0 |
|       1 |       0 |      1 |
|       0 |       1 |      1 |
|       1 |       1 |      0 |

### Xnor Gate
The xnor gate has the same truth table as a nor gate, except when both inputs are 1, xnor outputs a 1. Xnor can also be represented as passing the output of the nor gate for input1 and input2 as input for a not gate. The output of that not gate is also the output of the xnor gate. Xnor has the following truth table:

| Input 1 | Input 2 | Output |
|---------|---------|--------|
|       0 |       0 |      1 |
|       1 |       0 |      0 |
|       0 |       1 |      0 |
|       1 |       1 |      1 |

That covers all of the logic gates. Remember that where ever there is a 1, the code represents the value as true and where ever there is a 0, the code represents the value as false.
