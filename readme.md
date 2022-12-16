# Simplified AES

Simplified AES (S-AES) was developed by Professor Edward Schaefer of Santa Clara University and several of his students. It is an educational rather than a secure encryption algorithm. It has similar properties and structure to AES with much smaller parameters. A good grasp of S-AES will make it easier for you to appreciate the structure and workings of AES.

The following is the Algorithm of S-AES:

> S-AES: The encryption algorithm takes a 16-bit block of plaintext as input and a 16-bit key and produces a 16-bit block of ciphertext as output. The S-AES decryption algorithm takes a 16-bit block of ciphertext and the same 16-bit key used to produce that ciphertext as input and produces the original 16-bit block of plaintext as output. The encryption algorithm involves the use of four different functions, or transformations: add key, nibble substitution (NS), shift row (SR), and mix column (MC), whose operation is explained subsequently. We can concisely express the encryption algorithm as a composition of functions:

Encryption:

![1](https://user-images.githubusercontent.com/73187712/208061567-53f86c21-3523-43c7-b75a-b0157c75b3d9.png)


where A ~K0~ is applied first.

![](C:/Users/Amaan/Documents/Java/AES/2)

![](C:/Users/Amaan/Documents/Java/AES/3)

> Each function operates on a 16-bit state, treated as a matrix of nibbles, where one nibble equals 4 bits. The initial value of the State matrix is the 16-bit plaintext; State is modified by
> each subsequent function in the encryption process, producing after the last function the 16-bit ciphertext. As Figure 5.12a shows, the ordering of nibbles within the matrix is by column.

![](C:/Users/Amaan/Documents/Java/AES/4)

## S-AES Encryption and Decryption

We now look at the individual functions that are a part of the encryption algorithm.

==ADDKEY==

> The Add Key function conssistes of the bitwise XOR of the 16-bit State Matrix and the 16-Bit round Key. Fig. 5.14 depicts this as a columnwise opeartion, but it can also be viewed as a nibble-wis or bitwise operation. The following is an example:
> ![](C:/Users/Amaan/Documents/Java/AES/5)

> The inverse of the add key function is identical to the add key function, because XOR operation is its own inverse.

===NIBBLE SUBSTITUTION===

> The nibble substituiton function is a simple table look-up. AES defines a 4x4 matrix of nibbble values, called an S-box, that contains a permutation of all possible 4-bit values. Each individual nibble of State is mapped into a new nibble in the following way: The leftmost 2 bits of the nibble are used as a row value, and the rightmost 2 bits are used as a common value. These row and column values serve as indexes into the S-box to select a unique 4-bit output value. For example, the hexadecimal value A refernces row 2, column 2 of the S-box, which contains the value 0. Accordingly, the value of A is mapped into the value 0.
> Here is an example of the nibble substitution transformation.
