# Common_Traps

## Total Beginner

- Opening Brace Can't Be Placed on a Separate Line

1. 语句块的左方括号不能另起一行，编译出错

- Unused Variables

1. 定义变量不能不使用，编译出错

- Unused Imports

1. 导入包不能不使用，编译出错

- Short Variable Declarations Can Be Used Only Inside Functions

1. 简洁赋值不能在函数体外，编译错误

- Redeclaring Variables Using Short Variable Declarations

1. 不能重复简洁赋值，编译错误

- Can't Use Short Variable Declarations to Set Field Values

1. 不能为结构体字段简洁赋值

- Accidental Variable Shadowing

1. 短声明在不同作用域链中可以重复，但在不同作用域中不是同一个地址.逻辑错误，需静态检查。

- Can't Use "nil" to Initialize a Variable Without an Explicit Type

1. 不能给一个不确定类型的变量赋予 nil 值，编译错误。

- Using "nil" Slices and Maps

1. 可以给切片 nil 值，但不能给map类型 nil 值，运行时panic。

- Map Capacity

1. 对Map变量使用cap函数会报错,编译错误。

- Strings Can't Be "nil"

1. 不能给字符串赋予 nil 值，编译错误。

- Array Function Arguments

1. 数组作为函数形参是值传递，函数不会改变实参初始值。逻辑错误。

- Unexpected Values in Slice and Array "range" Clauses

1. range返回的是两个值，优先返回索引。逻辑错误

- Slices and Arrays Are One-Dimensional

1. 

- Accessing Non-Existing Map Keys
- Strings Are Immutable
- Conversions Between Strings and Byte Slices
- Strings and Index Operator
- Strings Are Not Always UTF8 Text
- String Length
- Missing Comma In Multi-Line Slice/Array/Map Literals
- log.Fatal and log.Panic Do More Than Log
- Built-in Data Structure Operations Are Not Synchronized
- Iteration Values For Strings in "range" Clauses
- Iterating Through a Map Using a "for range" Clause
- Fallthrough Behavior in "switch" Statements
- Increments and Decrements
- Bitwise NOT Operator
- Operator Precedence Differences
- Unexported Structure Fields Are Not Encoded
- App Exits With Active Goroutines
- Sending to an Unbuffered Channel Returns As Soon As the Target Receiver Is Ready
- Sending to an Closed Channel Causes a Panic
- Using "nil" Channels
- Methods with Value Receivers Can't Change the Original Value

## Intermediate Beginner

- Closing HTTP Response Body
- Closing HTTP Connections
- Unmarshalling JSON Numbers into Interface Values
- Comparing Structs, Arrays, Slices, and Maps
- Recovering From a Panic
- Updating and Referencing Item Values in Slice, Array, and Map "for range" - Clauses
- "Hidden" Data in Slices
- Slice Data Corruption
- "Stale" Slices
- Type Declarations and Methods
- Breaking Out of "for switch" and "for select" Code Blocks
- Iteration Variables and Closures in "for" Statements
- Deferred Function Call Argument Evaluation
- Deferred Function Call Execution
- Failed Type Assertions
- Blocked Goroutines and Resource Leaks

## Advanced Beginner

- Using Pointer Receiver Methods On Value Instances
- Updating Map Value Fields
- "nil" Interfaces and "nil" Interfaces Values
- Stack and Heap Variables
- GOMAXPROCS, Concurrency, and Parallelism
- Read and Write Operation Reordering
- Preemptive Scheduling 