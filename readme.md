# Simplified AES

Simplified AES (S-AES) was developed by Professor Edward Schaefer of Santa Clara University and several of his students. It is an educational rather than a secure encryption algorithm. It has similar properties and structure to AES with much smaller parameters. A good grasp of S-AES will make it easier for you to appreciate the structure and workings of AES.

The following is the Algorithm of S-AES:

> S-AES: The encryption algorithm takes a 16-bit block of plaintext as input and a 16-bit key and produces a 16-bit block of ciphertext as output. The S-AES decryption algorithm takes a 16-bit block of ciphertext and the same 16-bit key used to produce that ciphertext as input and produces the original 16-bit block of plaintext as output. The encryption algorithm involves the use of four different functions, or transformations: add key, nibble substitution (NS), shift row (SR), and mix column (MC), whose operation is explained subsequently. We can concisely express the encryption algorithm as a composition of functions:

Encryption:

![1](https://user-images.githubusercontent.com/73187712/208061567-53f86c21-3523-43c7-b75a-b0157c75b3d9.png)


where A<sub>K</sub><sub>0</sub> is applied first.

Decryption:

![2](https://user-images.githubusercontent.com/73187712/208062304-56d7c746-2ef1-4d46-b5d0-095147fb2d24.png)


![3](https://user-images.githubusercontent.com/73187712/208062353-945e06a9-ed59-4356-b5f9-7bd20061696c.png)

> Each function operates on a 16-bit state, treated as a matrix of nibbles, where one nibble equals 4 bits. The initial value of the State matrix is the 16-bit plaintext; State is modified by
> each subsequent function in the encryption process, producing after the last function the 16-bit ciphertext. As Figure 5.12a shows, the ordering of nibbles within the matrix is by column.

![4](https://user-images.githubusercontent.com/73187712/208062941-55a825f0-39c8-413b-a535-314ca91d976e.png)


## S-AES Encryption and Decryption

We now look at the individual functions that are a part of the encryption algorithm.

**ADDKEY**

> The Add Key function conssistes of the bitwise XOR of the 16-bit State Matrix and the 16-Bit round Key. Fig. 5.14 depicts this as a columnwise opeartion, but it can also be viewed as a nibble-wis or bitwise operation. The following is an example:

![fifth](https://user-images.githubusercontent.com/73187712/208072675-775b02df-f332-4408-a6d1-1a8decebc7c9.png)

> The inverse of the add key function is identical to the add key function, because XOR operation is its own inverse.

**NIBBLE SUBSTITUTION**

> The nibble substituiton function is a simple table look-up. AES defines a 4x4 matrix of nibbble values, called an S-box, that contains a permutation of all possible 4-bit values. Each individual nibble of State is mapped into a new nibble in the following way: The leftmost 2 bits of the nibble are used as a row value, and the rightmost 2 bits are used as a common value. These row and column values serve as indexes into the S-box to select a unique 4-bit output value. For example, the hexadecimal value A refernces row 2, column 2 of the S-box, which contains the value 0. Accordingly, the value of A is mapped into the value 0.
> Here is an example of the nibble substitution transformation.

![6](https://user-images.githubusercontent.com/73187712/208072834-f0f421e5-f74e-43f7-815d-cd2b71850cac.png)

The inverse nibble substitution function makes use of the inverse S-box. Note, for example, that the input 0 produces the output A, and the input A to the S-box produces 0.

![7](https://user-images.githubusercontent.com/73187712/208073339-d047077b-0d41-4704-bc30-6690a3d3ae12.png)

**SHIFT ROW**

>The Shift row function performs a one-nibble circular shift of the second row of State the first row is not altered. The following is an example.

![8](https://user-images.githubusercontent.com/73187712/208073771-c483d894-bc75-4e20-b348-99b2bf64cf4e.png)

>The inverse shift row function is identical to the shift row function, because it shifts the second row back to its original position.

**MIX COLUMN**

>The mix column function operates on each column individually. Each nibble of a column is mapped into a new value that is a function of both

![9](https://user-images.githubusercontent.com/73187712/208077072-81b8f029-6725-458c-a1f9-c94ca6727747.png)

>nibbles in that column. The transformation can be defined by the following matrix multiplication on **State**

![10](https://user-images.githubusercontent.com/73187712/208077792-2a8f274e-4b11-46ed-9154-660106b81603.png)

>Performing the matrix multiplication, we get

![11](https://user-images.githubusercontent.com/73187712/208077955-db2bc3a6-e554-4e03-8323-64d93ddb5b20.png)

where arithmetic is performed in GF(2<sup>4</sup>), and the symbol •	refers to a multiplication in GF(2<sup>4</sup>).

**KEY EXPANSION**

>For key expansion, the 16 bits of the initial key are grouped into a row of two 8-bit words. Below figure shows the expansion into sex words, by the calculation of four new words from the initial two words. The algorithm is 

![12](https://user-images.githubusercontent.com/73187712/208079014-191a7792-028a-4788-9951-0ea7f9c588dd.png)

>Rcon is a round constant, defined as follows: RC[i] = x<sup>i+2</sup>, so that RC[1] = x<sup>3</sup> = 1000 and RC[2] = x<sup>4</sup>mod(x<sup>4</sup> + x + 1) = x + 1 = 0011. RC[i] forms the left-most nibble of a byte, with the right-most nibble being all zeroes. Thus, Rcon(1) = 10000000 and Rcon(2) = 00110000.
>For example, suppose the key is 2D55 = 0010 1101 0101 0101 = w<sub>0</sub>w<sub>1</sub>. Then

![13](https://user-images.githubusercontent.com/73187712/208080308-fa3315c9-0d18-42d7-aecb-5e6520983a96.png)

