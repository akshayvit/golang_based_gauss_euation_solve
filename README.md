Guassian Elimination implemented to solve 3,4 degree equations and serve it in web service with random generated equations

<button class="btn btn-success"><a href="https://jimkang.medium.com/install-go-on-mac-with-homebrew-5fa421fc55f5" target="blank">Get started with go in MAC</a></button>

<h1>Execute through thr following command:</h1>

go run main.go

<hr>
Test at : http://localhost:8080/show3 : For 3 degree equations


For 4 degree equations : http://localhost:8080/show4


<hr>

<p>
What is the Gauss Elimination Method?

In mathematics, the <strong>Gaussian elimination method</strong> is known as the <strong>row reduction algorithm</strong> for solving linear equations systems. It consists of a sequence of operations performed on the corresponding matrix of coefficients. We can also use this method to estimate either of the following:

The rank of the given matrix
The determinant of a square matrix
The inverse of an invertible matrix
To perform row reduction on a matrix, we have to complete a sequence of elementary row operations to transform the matrix till we get 0s (i.e., zeros) on the lower left-hand corner of the matrix as much as possible. That means the obtained matrix should be an upper triangular matrix. There are three types of elementary row operations; they are:

Swapping two rows and this can be expressed using the notation ↔, for example, R2 ↔ R3
Multiplying a row by a nonzero number, for example, R1 → kR2 where k is some nonzero number
Adding a multiple of one row to another row, for example, R2 → R2 + 3R1
Learn more about the elementary operations of a matrix here.

<strong>The obtained matrix will be in row echelon form. The matrix is said to be in reduced row-echelon form when all of the leading coefficients equal 1, and every column containing a leading coefficient has zeros elsewhere.</strong>

</p>
<hr>
[ Reference for the above method explanations here is : https://byjus.com/maths/gauss-elimination-method/#:~:text=In%20mathematics%2C%20the%20Gaussian%20elimination,the%20corresponding%20matrix%20of%20coefficients. ]
<hr>
<hr>

There are two major functions :

<strong>Forward elimination</strong> : To convert to row-echeleon method.

<strong>Backward Substituiton</strong> : To convert to reduced row echeleon form. [ Where the actual solution for finite solution based consistent equations are obtained ]
