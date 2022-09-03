/*
Copyright 2022 The Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.
*/

/*
The calc Project is a proposal to implement a calculator
specialized science. This is still under development
with several pending issues and possibilities still open.

# Components

At the moment calc has two modules:

  - cmath - with types, attributes and methods to operate with
    matrices and vectors;
  - cnumeric - with statistical functions and calculation of
    roots of polynomials (at the moment only 2nd order plinoms
    are implemented).

The project is still in the embryonic stage and does not have
user interface, but only in the development of the tools.

# Future Additions

The project provides for the addition of the following features:
 1. cmath: add resolution of eigenvalues ​​and eigenvectors;
 2. cmath: add matrix, inverse matrix, adjunct and transpose
    diagonalization.
 3. cmath: add determinant calculus for NxN matrix.
 4. calc: add GUI employ go-gtk,
    (https://github.com/mattn/go-gtk)
 5. calc: implement a numeric expression interpreter to convert
    "(a+b)xc" to "CrossProd(a.add(b),c)"
*/
package main
