---
sort: 3
spin: 3
span: 
suit: 3
description: 
---
# Primes Platform

We analyze the inference process of dynamic and observe that the dynamic tensor shapes and control flow hinder by prime numbers

{% include list.liquid all=true %}

See [this paper](https://github.com/user-attachments/files/17277693/s11390-021-1161-y.pdf) for a discussion of the challenges imposed by dynamic shapes and one project's approach to addressing them.

## Dynamic Shapes

[Prime hexagon](https://youtu.be/fQL4KRH3wUQ) is a mathematical structure that is formed when integers are sequentially added to a field of tessellating equilateral triangles

```note
This is not easy as they are linked to the nature of prime numbers, and nothing is easy about the nature of prime numbers.  But I begin with this assumption: if the hexagons participate in the Universe in any way other than haphazardly, they must be demonstrably ***congruent to something organized*** _([T. Gallion](https://www.hexspin.com/minor-hexagons/))_.
```

```txt
s p i n
0 0 0 0
1 0 0 0
2 0 1 0  ◄--- 1st prime
3 1 1 0  ◄--- 2nd prime
--------
5 2 1 0  ◄--- 3rd prime
7 3 1 0
11 4 1 0
13 5 1 0
17 0 1 1 ◄--- 7th prime
19 1 1 1 ◄--- 8th prime
```

The program used to demonstrate functions of dynamic shapes where the path of the integers is changed whenever a prime is encountered.

```note
- Create a program that includes dynamic shapes in program inputs and outputs
- Import that program into IREE's compiler
- Compile that program to an IREE VM bytecode module
- Load the compiled program using IREE's high level runtime C API
- Call exported functions on the loaded program
```

***17 = 7th prime = (18 - 11) th prime***

[![Prime hexagon](https://user-images.githubusercontent.com/8466209/226643815-f8c926cd-ae36-43d3-b271-8df07272f429.png)](https://youtu.be/fQL4KRH3wUQ)

To match the different shape, we will need a common shape, and the result is a two-dimensional array. This would later be multilinear dimensions.

```txt
p r i m e s
1 0 0 0 0
2 1 0 0 0
3 2 0 1 0 2 ◄--- 1st prime
4 3 1 1 0 3 ◄--- 2nd prime
5 5 2 1 0 5 ◄--- 3rd prime
6 7 3 1 0
7 11 4 1 0
8 13 5 1 0
9 17 0 1 1 17 ◄--- 7th prime
10 19 1 1 1 19 ◄--- 8th prime
-----
11 23 2 1 1 23 ◄--- 9th prime √
```

You may learn that sets of algebraic objects has a multilinear relationship related to a vector space called _[tensor](https://en.wikipedia.org/wiki/Tensor)_. Tensors may map between different objects such as vectors, scalars, and even other tensors.

```note
Tensors are multi-dimensional arrays with a uniform type (e.g. int32, float32) and a shape. Shapes consist of a rank and a list of dimensions and may be static (i.e. fully known and fixed) or varying degrees of dynamic. See these references:
- PyTorch: [Compiler dynamic shapes](https://pytorch.org/docs/stable/torch.compiler_dynamic_shapes.html), [torch.Tensor](https://pytorch.org/docs/stable/tensors.html)
- TensorFlow: [Introduction to Tensors](https://www.tensorflow.org/guide/tensor)

All tensors are immutable like Python numbers and strings: you can never update the contents of a tensor, only create a new one. [Dynamic shapes](https://github.com/iree-org/iree/tree/main/samples/dynamic_shapes) are useful for passing variable sized batches as input, receiving variable length sentences of text as output.
```

[![300px-Components_stress_tensor svg](https://user-images.githubusercontent.com/8466209/211590875-9cae3c47-bbdc-43d8-acc5-d8c64c802967.png)](https://en.wikipedia.org/wiki/Tensor)

Connecting nodes seems a trivial operation, but it hides some difficulties related to the concept of tensor’s shape in both its variants: static and dynamic.

```note
Every tensor has a name, a type, a rank and a shape.
- The name uniquely identifies the tensor in the computational graphs (for a complete understanding of the importance of the tensor name and how the full name of a tensor is defined, I suggest the reading of the article [Understanding Tensorflow using Go](https://pgaleone.eu/tensorflow/go/2017/05/29/understanding-tensorflow-using-go/)).
- The type is the data type of the tensor, e.g.: a tf.float32, a tf.int64, a tf.string, …
- The rank, in the [ Tensorflow](https://pgaleone.eu/tensorflow/2018/07/28/understanding-tensorflow-tensors-shape-static-dynamic/#) world (that’s different from the [ mathematics](https://pgaleone.eu/tensorflow/2018/07/28/understanding-tensorflow-tensors-shape-static-dynamic/#) world), is just the number of dimension of a tensor, e.g.: a scalar has rank 0, a vector has rank 1, …
- The shape is the number of elements in each dimension, e.g.: a scalar has a rank 0 and an empty shape (), a vector has rank 1 and a shape of (D0), a matrix has rank 2 and a shape of (D0, D1) and so on.

So you might wonder: what’s difficult about the shape of a tensor? It just looks easy, is the number of elements in each dimension, hence we can have a shape of () and be sure to work with a scalar, a shape of (10) and be sure to work with a vector of size 10, a shape of (10,2) and be sure to work with a matrix with 10 rows and 2 columns. Where’s the difficulty? _([tensors shape: static and dynamic](https://pgaleone.eu/tensorflow/2018/07/28/understanding-tensorflow-tensors-shape-static-dynamic/))_
```

```txt
p r i m e s
1 0 0 0 0
2 1 0 0 0
3 2 0 1 0 2 ◄--- 1st prime
4 3 1 1 0 3 ◄--- 2nd prime
5 5 2 1 0 5 ◄--- 3rd prime
6 7 3 1 0
7 11 4 1 0
8 13 5 1 0
9 17 0 1 1 17 ◄--- 7th prime
10 19 1 1 1 ∆1 ◄--- 8th prime ◄--- Terminating Digit
-----
11 23 2 1 1 √
```

The difficulties (and the cool stuff) arises when we dive deep into the Tensorflow peculiarities, and we find out that there’s no constraint about the definition of the shape of a tensor.

```note
Tensorflow, in fact, allows us to represent the shape of a Tensor in 3 different ways:
- Fully-known shape: that are exactly the examples described above, in which we know the rank and the size for each dimension.
- Partially-known shape: in this case, we know the rank, but we have an unknown size for one or more dimension (everyone that has trained a model in batch is aware of this, when we define the input we just specify the feature vector shape, letting the batch dimension set to None, e.g.: (None, 28, 28, 1).
Unknown shape and known rank: in this case we know the rank of the tensor, but we don’t know any of the dimension value, e.g.: (None, None, None).
- Unknown shape and rank: this is the toughest case, in which we know nothing about the tensor; the rank nor the value of any dimension.

Tensorflow, when used in its non-eager mode, separates the graph definition from the graph execution. This allows us to first define the relationships among nodes and only after executing the graph.
```

***(17+13) + (11+19) = (7+11) + (19+23) = 60***

![image](https://github.com/eq19/maps/assets/8466209/a95b5e7d-37e4-47ff-aeb2-55361b98a37b)

Remember that the success of a machine learning model depends on the fact that both the train and the test set are from the same distribution; however, it seems that this constraint is violated in the case of [prime numbers](https://github.com/zetbaitsu/tensorflow_sample/blob/master/prime.py).

```note
Machine learning method just learn a function mapping x to f(w; x). it then define a loss function as loss(y, f(w;x)). x is a vector representing features of something. 
```

![image](https://github.com/eq19/maps/assets/8466209/922fac71-d793-477c-a56c-ec21e552d695)

It doesn’t do learning in the required way to learn from examples, but at least it was proposing mathematical theorems and representing mathematical ideas in a more natural format such as _[terminating digits](https://primesdemystified.com/#deepsymmetries)_.

```txt
p r i m e s
1 0 0 0 0 0
2 1 0 0 0 1
3 2 0 1 0 2
4 3 1 1 0 3
5 5 2 1 0 5
6 7 3 1 0 7
7 11 4 1 0 11
8 13 5 1 0 13
9 17 0 1 1 17 ◄--- 7th prime
10 19 1 1 1 ∆1 ◄--- 0th ∆prime ◄--- Terminating Digit #0 √
-----
11 23 2 1 1 ∆2 ◄--- 1st ∆prime ◄--- Terminating Digit #1 √
+29 rows √
-----
41 √
```

In order to maintain the 36 symmetry (whether it is an addition zone or not), with this prime number 19 was found at least seven (7) pairs of _[truncated patterns](https://www.eq19.com/multiplication/file19.html#truncated-patterns)_.

```py
total_parameters = 0
for variable in tf.trainable_variables():
    # shape is an array of tf.Dimension
    shape = variable.get_shape()
    print(shape)
    print(len(shape))
    variable_parameters = 1
    for dim in shape:
        print(dim)
        variable_parameters *= dim.value
    print(variable_parameters)
    total_parameters += variable_parameters
print(total_parameters)
```

***π(6+11) = π(17) = 7***

![](https://user-images.githubusercontent.com/36441664/274093531-0878e3e5-6be3-448e-9ad4-3b34523c1e9c.jpg)

By which we finally found if this behaviour is cascaded bilaterally within the correlation between ***61*** as [the 18th prime](https://gist.github.com/eq19/e9832026b5b78f694e4ad22c3eb6c3ef) and ***67*** as [the 19th prime](https://gist.github.com/eq19/c9bdc2bbe55f2d162535023c8d321831).

```note
The solution is not only to prove Re(z)= 1/2 but also to calculate ways for the imaginary part of the complex root of ζ(z)=0 and also to solve the Functional equations of Riemann
```

***18 + 19 = π(61) + π(67) = 37***

![](https://camo.githubusercontent.com/492732f872ab84bf8c013f106230ea47d355661f5faf6a6e3b0661b30cb9bc28/68747470733a2f2f75706c6f61642e77696b696d656469612e6f72672f77696b6970656469612f636f6d6d6f6e732f662f66632f5269656d616e6e5a6574615f5a65726f732e737667)

```txt
p r i m e s
1 0 0 0 0 0
2 1 0 0 0 1
3 2 0 1 0 2
4 3 1 1 0 3
5 5 2 1 0 5
6 7 3 1 0 7
7 11 4 1 0 11
8 13 5 1 0 13
9 17 0 1 1 17 ◄--- 7th prime
10 19 1 1 1 ∆1 ◄--- 0th ∆prime ◄--- Terminating Digit #0 (spin 18) √
-----
11 23 2 1 1 ∆2 ◄--- 1st ∆prime ◄--- Terminating Digit #1 (spin 19) √
+29 rows
-----
41
```

This polarity is happened per ***six (6) cycles*** by the polar of ***six (6) to one (1)*** and ***six (6) to seven (7)*** that leads to the prime number ***61 and 67***.

```note
The above ***characteristics of primes in the hexagon suggests 0 family numbers split more than twin primes***. I speculate these numbers split all primes. That is, all primes have a partner (of the opposite family) equidistant from such a number. For instance, ***0 family member 18 splits twin primes 17 and 19***, but is also 5 more than 13 and 5 less than 23, and it is also 11 more the 7, and 11 less than 29, etc. _([Hexspin](https://www.hexspin.com/cell-types/))_
```

[![](https://user-images.githubusercontent.com/8466209/219239425-90f075fa-fe8a-4f80-b3ce-7b2053956c6b.png)](http://www.hexspin.com/0-1-and-negative-numbers/)

So it is not strictly true that neural networks cannot be trained to recognize prime numbers, but there are significant challenges associated with this task.

## Central Polarity

```note
The Prime Spiral Sieve possesses remarkable structural and numeric symmetries. For starters, the intervals between the prime roots (and every subsequent row or rotation of the sieve) are perfectly balanced, with a period ***eight (8) difference sequence*** of: {6, 4, 2, 4, 2, 4, 6, 2} 
_([Primesdemystified](https://primesdemystified.com/#deepsymmetries))_.
```

[![image](https://user-images.githubusercontent.com/8466209/219260933-4331d79b-5815-4566-82e3-1a485bb2c61f.png)](https://primesdemystified.com/#deepsymmetries)

```txt
p r i m e s
1 0 0 0 0 0
2 1 0 0 0 1
3 2 0 1 0 2
4 3 1 1 0 3
5 5 2 1 0 5
6 7 3 1 0 7
7 11 4 1 0 11
8 13 5 1 0 13
9 17 0 1 1 17 ◄--- 7th prime
10 19 1 1 1 ∆1 ◄--- 0th ∆prime ◄--- Terminating Digit #0 (spin 18)
-----
11 23 2 1 1 ∆2 ◄--- 1st ∆prime ◄--- Terminating Digit #1 (spin 19)
+29 rows
-----
41
+59 rows √
```

```note
Speaking of the Fibonacci number sequence, there is symmetry mirroring the above in the relationship between the terminating digits of Fibonacci numbers and their index numbers equating to members of the array populating the Prime Spiral Sieve.
```

[![11's additive sums](https://user-images.githubusercontent.com/8466209/221473004-867a1b50-f91f-470d-9922-e5e4f543a590.png)](https://primesdemystified.com/#deepsymmetries)

***Fibonacci level-1 (29) x Fibonacci level-2 (59) = 10x10 = 💯***

```txt
p r i m e s
1 0 0 0 0 0
2 1 0 0 0 1
3 2 0 1 0 2
4 3 1 1 0 3
5 5 2 1 0 5
6 7 3 1 0 7
7 11 4 1 0 11
8 13 5 1 0 13
9 17 0 1 1 17 ◄--- 7th prime
10 19 1 1 1 ∆1 ◄--- 0th ∆prime ◄--- Terminating Digit #0 ◄- Fibonacci Index #18 √
-----
11 23 2 1 1 ∆2 ◄--- 1st ∆prime ◄--- Terminating Digit #1 ◄- Fibonacci Index #19 √
..
..
40 163 5 1 0 ∆31 ◄- 11th ∆prime ◄- Terminating Digit #11 ◄- Fibonacci Index #29 √
-----
41
+59 rows ◄--- total 41+59 = 💯 rows = 10x10 rows √
```

## Numeric Symmetries

***(59² − 31²) = 360 x 7***

![Squares_Distribution](https://github.com/eq19/maps/assets/8466209/0f7725e3-52b0-4e5d-93bb-b12358bdf1d4)

```txt
p r i m e s
1 0 0 0 0 0
2 1 0 0 0 1
3 2 0 1 0 2
4 3 1 1 0 3
5 5 2 1 0 5
6 7 3 1 0 7
7 11 4 1 0 11
8 13 5 1 0 13
9 17 0 1 1 17 ◄--- 7th prime
10 19 1 1 1 ∆1 ◄--- 0th ∆prime ◄--- Fibonacci Index #18
-----
11 23 2 1 1 ∆2 ◄--- 1st ∆prime ◄--- Fibonacci Index #19
..
..
40 163 5 1 0 ∆31 ◄- 11th ∆prime ◄-- Fibonacci Index #29
-----
41 167 0 1 1 ∆0
42 173 0 -1 1 ∆1
43 179 0 1 1 ∆2 ◄--- ∆∆1
44 181 1 1 1 ∆3 ◄--- ∆∆2 ◄--- 1st ∆∆prime ◄--- Fibonacci Index #30 ✔️
..
..
100 521 0 -1 2 ∆59 ◄--- ∆∆17 ◄--- 7th ∆∆prime ◄--- Fibonacci Index #36 ✔️
-----
```

![](https://user-images.githubusercontent.com/8466209/219261961-10e4d77f-ead3-43d4-9407-f01d83f1f204.png)

```txt
p r i m e s
1 0 0 0 0 0
2 1 0 0 0 1
3 2 0 1 0 2
4 3 1 1 0 3
5 5 2 1 0 5
6 7 3 1 0 7
7 11 4 1 0 11
8 13 5 1 0 13
9 17 0 1 1 17 ◄--- 7th prime 👉 7s ✔️
10 19 1 1 1 ∆1 ◄--- 0th ∆prime ◄--- Fibonacci Index #18
-----
11 23 2 1 1 ∆2 ◄--- 1st ∆prime ◄--- Fibonacci Index #19
..
..
40 163 5 1 0 ∆31 ◄- 11th ∆prime ◄-- Fibonacci Index #29 👉 11s ✔️
-----
41 167 0 1 1 ∆0
42 173 0 -1 1 ∆1
43 179 0 1 1 ∆2 ◄--- ∆∆1
44 181 1 1 1 ∆3 ◄--- ∆∆2 ◄--- 1st ∆∆prime ◄--- Fibonacci Index #30
..
..
100 521 0 -1 2 ∆59 ◄--- ∆∆17 ◄--- 7th ∆∆prime ◄--- Fibonacci Index #36  👉 7s ✔️
-----
```

```note
These positions: 1 7 11 13 17 19 23 29. We refer to this basic system as MEC 30 - "Mathematical Elementary Cell 30".
- By repeating the positions we show the function of the basic system in the next step. If we extend the 30th order of the MEC, for example, to the number 120, the result is 4 times a 30th order and thus 4 × 8 = 32 prime positions.
- Hypothetical assumption: If the product of the primes (except 2, 3, 5,) would not fall into the prime positions, thus be divided by 2, 3 or 5, the information would have 120 = 32 primes in 32 prime positions.
- Prime positions (not the primes) 1, 7, 11, 13, 17, 19, 23, 29, / 1, 7, 11, 13, 17, 19, 23, 29, / 1, 7, 11, 13, 17 , 19, 23, 29, / 1, 7, 11, 13, 17, 19, 23, 29,
- The 30th order is repeated in the number space 120 = 4 times, 4 × 8 = 32 prime positions, thus 4 terms. From our considerations and also from the graphic see 2 However, we can conclude that the distribution of prime numbers must have a static base structure, which is also confirmed logically in the further course. 

This static structure is altered by the products of the primes themselves, since these products must fall into the prime positions since they are not divisible by 2, 3 and 5.
```

![](https://user-images.githubusercontent.com/36441664/74366957-992db780-4e03-11ea-8f26-cca32bd26003.png)

```note
The numbers not divisible by 2, 3 or 5 are highlighted. We call them prime positions, hence 1, 7, 11, 13, 17, 19, 23, 29. Important for our work is that in the following the term prime refers only to prime numbers that are in the prime positions. ***So primes 2, 3 and 5 are always excluded***.
```

```txt
p r i m e s
1 0 0 0 0 0
2 1 0 0 0 1 ◄--- #29 ✔️
3 2 0 1 0 2
4 3 1 1 0 3
5 5 2 1 0 5
6 7 3 1 0 7 ◄--- #23 ✔️
7 11 4 1 0 11 ◄--- #19 ✔️
8 13 5 1 0 13 ◄--- # 17 ✔️
9 17 0 1 1 17 ◄--- 7th prime 👉 7s
10 19 1 1 1 ∆1 ◄--- 0th ∆prime ◄--- Fibonacci Index #18
-----
11 23 2 1 1 ∆2 ◄--- 1st ∆prime ◄--- Fibonacci Index #19
..
..
40 163 5 1 0 ∆31 ◄- 11th ∆prime ◄-- Fibonacci Index #29 👉 11
-----
41 167 0 1 1 ∆0
42 173 0 -1 1 ∆1
43 179 0 1 1 ∆2 ◄--- ∆∆1
44 181 1 1 1 ∆3 ◄--- ∆∆2 ◄--- 1st ∆∆prime ◄--- Fibonacci Index #30
..
..
100 521 0 -1 2 ∆59 ◄--- ∆∆17 ◄--- 7th ∆∆prime ◄--- Fibonacci Index #36  👉 7s
-----
```

```note
In this one system, reproduced as an icon, we can show the distribution profile of the primes as well as their products over a checkerboard-like model in the 4.
- We show this fundamental causal relationship in the MEC 30 mathematically accurate in the table 13 , The organization of this table is based on the well-known idea of Christian Goldbach. That every even number should consist of the sum of two primes.
- All pairs of prime numbers without "1", 2, 3, 5, we call henceforth Goldbach pairs. The MEC 30 transforms this idea of Christian Goldbach into the structure of a numerical double-strand, into an opposing member of the MEC 30 scale.
- We call this double strand a convolution, which results in an opposite arrangement. It represents the natural vibration, thus also the redundant vibrations in the energy transfer. In the 6 For example, in the graph, the even number 60 is folded. At folding of the even number 60 6 result in 8 prime pairs.
- In this case, among the 8 pairs of prime pairs there are only 6 Goldbach pairs. 2 prime positions in the prime position pairs carry products of the factors "1 × 1" and 7 × 7. Thus, 2 prime pairs do not fulfill the requirements of the Goldbach pairs. In general, any even number larger than 30 can be represented graphically within a cycle (MEC 30) as a specific cyclic convolution. This characteristic convolution of the even numbers is a fundamental test element in the numerical table. The result Even the even numbers to infinity occupy a fixed position within the 30s system MEC 30. The even numbers thus have 15 positions: 30/2 = 15 even positions of the MEC 30.
- There are therefore only 15 even positions for all even numbers to infinity. Every even number has a specific convolution due to its position in the 30s system. First, we have to determine the positions of the even numbers in the 30s system to make them one in the following graph 7 attributable to the 15 specific folds.
```

![](https://user-images.githubusercontent.com/36441664/74591945-2b75cb80-504f-11ea-85dd-14d0a803ee6b.png)

```txt
p r i m e s
1 0 0 0 0 0
2 1 0 0 0 1 ◄--- #29 ◄--- #61 ✔️
3 2 0 1 0 2
4 3 1 1 0 3
5 5 2 1 0 5
6 7 3 1 0 7 ◄--- #23
7 11 4 1 0 11 ◄--- #19
8 13 5 1 0 13 ◄--- # 17 ◄--- #49 ✔️
9 17 0 1 1 17 ◄--- 7th prime 👉 7s
10 19 1 1 1 ∆1 ◄--- 0th ∆prime ◄--- Fibonacci Index #18
-----
11 23 2 1 1 ∆2 ◄--- 1st ∆prime ◄--- Fibonacci Index #19 ◄--- #43 ✔️
..
..
40 163 5 1 0 ∆31 ◄- 11th ∆prime ◄-- Fibonacci Index #29 👉 11
-----
41 167 0 1 1 ∆0
42 173 0 -1 1 ∆1
43 179 0 1 1 ∆2 ◄--- ∆∆1
44 181 1 1 1 ∆3 ◄--- ∆∆2 ◄--- 1st ∆∆prime ◄--- Fibonacci Index #30
..
..
100 521 0 -1 2 ∆59 ◄--- ∆∆17 ◄--- 7th ∆∆prime ◄--- Fibonacci Index #36  👉 7s
-----
```

## Palindromic Sequence 

```note
In the matrix pictured below, we list the first 24 elements of our domain, take their squares, calculate the modulo 90 congruence and digital roots of each square, and display the digital root factorization dyad for each square (and map their ***collective bilateral 9 sum symmetry***). _([PrimesDemystified](https://primesdemystified.com/))_
```

***7 x π(89) = 7 x 24 = 168 = π(1000)***

[![collective bilateral 9 sum symmetry](https://user-images.githubusercontent.com/8466209/200229388-03811f48-2492-4845-b15f-85259cd93717.png)](https://www.eq19.com/multiplication/#elementary-structure)

```txt
p r i m e s
1 0 0 0 0 0
2 1 0 0 0 1 ◄--- #29 ◄--- #61
3 2 0 1 0 2 👉 2 ✔️
4 3 1 1 0 3 👉 61 - 1 = 60 ✔️
5 5 2 1 0 5
6 7 3 1 0 7 ◄--- #23
7 11 4 1 0 11 ◄--- #19
8 13 5 1 0 13 ◄--- # 17 ◄--- #49
9 17 0 1 1 17 ◄--- 7th prime 👉 7s
10 19 1 1 1 ∆1 ◄--- 0th ∆prime ◄--- Fibonacci Index #18
-----
11 23 2 1 1 ∆2 ◄--- 1st ∆prime ◄--- Fibonacci Index #19 ◄--- #43
..
..
40 163 5 1 0 ∆31 ◄- 11th ∆prime ◄-- Fibonacci Index #29 👉 11
-----
41 167 0 1 1 ∆0
42 173 0 -1 1 ∆1
43 179 0 1 1 ∆2 ◄--- ∆∆1
44 181 1 1 1 ∆3 ◄--- ∆∆2 ◄--- 1st ∆∆prime ◄--- Fibonacci Index #30
..
..
100 521 0 -1 2 ∆59 ◄--- ∆∆17 ◄--- 7th ∆∆prime ◄--- Fibonacci Index #36  👉 7s
-----
```

```note
The color spin addresses for numbers are generally straightforward – a composite number takes the spin of the prior prime.  4 spins blue because 3 spins blue.  8 is red because 7 is red.  However, twin primes, and the 0 type numbers between them, are open to some interpretation.
```

[![base](https://user-images.githubusercontent.com/8466209/249753163-6cfbcecf-3713-409b-8d8b-5fa5cf8489ac.png)](https://www.hexspin.com/finding-a-number-in-the-hexagon/)

***(43 - 19)the prime = 24th prime = 89***

```
$True Prime Pairs:
(5,7), (11,13), (17,19)
 
layer | node | sub |  i  |  f.                                       MEC 30 / 2
------+------+-----+-----+------      ‹------------------------------ 0 {-1/2}
      |      |     |  1  | --------------------------
      |      |  1  +-----+                           |    
      |  1   |     |  2  | (5)                       |
      |      |-----+-----+                           |
      |      |     |  3  |                           |
  1   +------+  2  +-----+----                       |
      |      |     |  4  |                           |
      |      +-----+-----+                           |
      |  2   |     |  5  | (7)                       |
      |      |  3  +-----+                           |
      |      |     |  6  |                          11s ‹-- ∆28 = (71-43) √
------+------+-----+-----+------      } (36)         |
      |      |     |  7  |                           |
      |      |  4  +-----+                           |
      |  3   |     |  8  | (11)                      |
      |      +-----+-----+                           |
      |      |     |  9  |‹-- ∆9 = (89-71) / 2 √     |
  2   +------|  5  +-----+-----                      |
      |      |     |  10 |                           |
      |      |-----+-----+                           |
      |  4   |     |  11 | (13) --------------------- 
      |      |  6  +-----+            ‹------------------------------ 15 {0}
      |      |     |  12 |---------------------------
------+------+-----+-----+------------               |
      |      |     |  13 |                           |
      |      |  7  +-----+                           |
      |  5   |     |  14 | (17)                      |
      |      |-----+-----+                           |
      |      |     |  15 |                           7s ‹-- ∆24 = (43-19) √
  3   +------+  8  +-----+-----       } (36)         |
      |      |     |  16 |                           |
      |      |-----+-----+                           |
      |  6   |     |  17 | (19)                      |
      |      |  9  +-----+                           |
      |      |     |  18 | -------------------------- 
------|------|-----+-----+-----  ‹----------------------------------- 30 {+1/2}
```

```note
The number 120 has 32 prime positions minus 5 prime number products = 27 prime numbers. The information of the prime number products translates our theory into a checkerboard-like pattern using the finite 8 prime positions from the MEC 30, we call it Ikon. 8 × 8 primary positions = 64 primary positions of the checkerboard icon.
- First, there are two main features that we use. To Ikon 1: The primes information and their products. In this left icon, the redundants (the doubles) are to be determined through the number information in the positions Impeccable.
- Second: The product positions. In the icon, the cyclic behavior is shown in identical 8 horizontal and 8 vertical orders, we call these orders templates that would not be visible through the pure number information. The cyclical behavior of the 8 × 8 product positions continues indefinitely.
- Since the prime positions are finite, a total of 8 positions in the 30th order, an already revolutionary system opens up, the entire infinite distribution of the prime number products in an icon as a "checkerboard pattern". represent and thus obtain mathematically exact results.
- The three and 4 , Square Graphics (Ikon) will now be in the following, larger graphic 5 transfer. As stated above, we use the properties of the numbers, they consist of one information and one position. Thus we are able to calculate the redundant product positions by means of identical information in different positions.
- And subtracting them from the total prime positions gives us the number of prime numbers. This succeeds due to the self-similarity of the 30th order of the MEC 30, as shown in the graph 5 is articulated. At the top of the following larger graphic 5 the self-similarity of the 30th order (MEC 30) can be seen.
- This results in a fundamental causal relation to the primes, systemically the products are entered into the position system. Therefore, the distribution of primes products also determines the distribution of primes themselves. The reason lies in the one-system, since the prime number as a number itself also consists of an information and a position.

We apply the same principle as above for the determination of the prime position. Only with the difference that we move in the even positions of the MEC 30.
```

***7 x π(89) = 7 x 24 = 168 = π(1000)***

![](https://user-images.githubusercontent.com/8466209/244855108-b80b3c2e-1a13-4b55-828a-72ed15b448dd.png)

```tip
"Not every beautiful theory has a physical application, but all fundamental physical theories found so far–without exception–have a type of beauty or symmetry." – Michio Kaku, The God Equation
```

[![Theory of Everything](https://github.com/eq19/maps/assets/8466209/6e26ada0-c545-4771-810b-9f721512a03c)](https://www.primesdemystified.com/Factorization.html)

If there is a function that can tell us weather x is a prime number, then nn may find it. But I don't think there is such function (maybe it exists because no mathematician prove there isn't).
