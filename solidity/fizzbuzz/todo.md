# ToDo
Please create the following 2 programs.
- You can use any libraries/frameworks.
- You can search for anything on google.
- You can refer to your past code.
- But keep sharing your screen and what you are doing.


## FizzBuzz Contract
- It has two functions.
  - increment()
    - Increment a internal counter.
  - getFizzBuzz()
    - Convert a value of the internal counter to string, and return it.
    - If a value of the internal counter equals to multiples of 3, return "Fizz".
    - If a value of the internal counter equals to multiples of 5, return "Buzz".
    - If a value of the internal counter equals to multiples of 3 and 5, return "FizzBuzz".
- Emit `Increment` event with a sender-address and an internal value after increment.

## Frontend Application
- Connect your ethereum account by Metamask
- Create a button to call `increment()` function.
- Create a button to call `getFizzBuzz()` function and show the result on the screen.