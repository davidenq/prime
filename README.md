# Large primes

## Instructions (rules) given

Implement a system that can find large prime numbers. You can only use the basic arithmetic operators in your language for this task. The system should be able to find prime numbers of over 1000 bits length, in less than one second. Since there is no way of doing that with certainty, the system should use a probabilistic approach, and allow the user to say how large the probability of error should be.

The core functionality should be exposed in a library that can be used from other programs in the same environment. The system should also expose an interface that would allow a user to use the functionality directly.


## Check list following the above instructions

- [ ] Find large prime numbers
- [x] You can only use the basic arithmetic operators in your language for this task
- [ ] The system should be able to find prime numbers of over 1000 bits length
- [x] In less than one second
- [x] Running more that two instances to talk with each other.
- [x] The system should use a probabilistic approach
- [x] Allow the user to say how large the probability of error should be
- [x] The core functionality should be exposed in a library that can be used from other programs in the same environment
- [x] The system should also expose an interface that would allow a user to use the functionality directly.

## Dependencies

In order to test the implementation you need to below dependencies.

**_IMPORTANT_**

This implementation has been proved only on `Linux version 5.3.0-42-generic (buildd@lcy01-amd64-026) (gcc version 9.2.1 20191008 (Ubuntu 9.2.1-9ubuntu2))`



`go > 1.10`
`nodejs > 12 Only if you want to test ffi` 
`python > 3`
`Docker`
`docker-compose`
`make`


## How to run

1. Running the http server
`$ > make run-server` It will running a http server as an API REST on port 8000.

2. Execute some tests from external client such as `Postman`, `curl` in command line and so.

`curl -X GET 'http://localhost:8000/check?prime=284&samples=100&umbral=90'`



### Request
endpoint: `/check`
method:   `GET`
body: `(JSON)`

| Key  |Type    | Limit | Required | Observations  |
|------|--------|-------|--------- | --------------|
| prime | string | -    |yes       | It is a value as a pseudoprime to be proved  |
| samples | string | 10 - 1000 | yes | none |
| umbral | string | 0 - 100 | yes | it must be the url where the client will execute the requests |

For instance,
```
curl -X GET 'http://localhost:8000/check?prime=284&samples=100&umbral=90'`
```

### Response

body: `(JSON)`

| Key  | Type |
|------|------|
| message | information about the solution |
| status| http status code |
| code| http code |

For instance,
```
{
    "status": "ok",
    "code": 200,
    "message": "284 is not a prime number. Pass with 0 samples. The margin error is 100.00",
}
```

## Background and Experience

Implementing some code to search prime numbers until now has been a easy task when the number is not so large. But when the requirement is searching a really large number, it  could be a really challenge because it's important to think in a different way, having in mind the performance and how to handle really large numbers in declarations in code. Those things have been really important to have in mind. For that reason the first thing made before to develop the solution, has been reading widely about how mathematicians have proposed mechanims to search large prime numbers and how I could handle large numbers, more than 128bytes (1024bits).

There are different questions that need to be answered before to start the development:
- Could I develop my own algorithm?
- What of them (theorems) have the best performance? 

After read widely and at the first time, I have tried to implement the Fermat's Litte Theorem based on (3). The first cases of test passed without errors (see `main_test.go or running make tests-primes if you want to know the status`), but I realized that the equation implemented was "wrong". I say wrong between quotes because it was a bad consideration don't use the numbers in the right way. Nevertheless, although the tests passed without errors for large numbers (not very very large numbers), the equation work, in a certain way, very well for non large numbers (less than 100 digits).

What was the bad consideration?

Well, Fermat's Little Theorem says:

Fermat Little Theorem says:
`p divides b^(p - 1)` in another words `b^(prime-1) congruent 1 (mod p)`
Where:
`p` is a pseudoprime, which must be proved.
`b` is a random number between 0 < b < p - 1

My bad consideration was that `p` must be a real prime number that has been proved and `b` the number which must be proved. I can't prove theoretically that this considerations is right even so it works well in certain cases. It could be a conjecture? I don't know.

Also I found some behaviours to test. For instance, when values are less than 1000, and the prime value is 2, 3 or 5, all the values passes correctly with an margin of error 0%. But when the values are more than 1000, the prime values to reduce a range error must be 2 and 3. In those cases, the margin of error is arround 40% in certain cases and in others, 0%. It means those values with 0% margin of error passed correctly. There are a lot of cases that I have tested with differente behaviours that I have not documented because I have not a lot of time for this.

Finally to reinforce the tests and check if the number is a prime number, a trial division as a first flow to test the number quickly has been implemented as well.


## Resources

Below there are some external links that have been used to understand in a better way how to use Baillie - PSW theorem to find large prime numbers (more than 1000 bits as has been required). This theorem has not been implemented yet. Fermat's Little theorem and Trial Division have been implemented.

1. [Average case error estimates for the strong probable prime test](https://math.dartmouth.edu/~carlp/PDF/paper88.pdf)
2. [bigprimes.org](https://bigprimes.org/how-it-works) to take a decision. [Baillie - PSW](https://en.wikipedia.org/wiki/Baillie%E2%80%93PSW_primality_test)
3. [The Prime Facts: From Euclid to AKS](https://www.scottaaronson.com/writings/prime.pdf) by [Scott Aaronson](https://www.scottaaronson.com/)
4. [Fermat Little Theorem](https://en.wikipedia.org/wiki/Fermat%27s_little_theorem)
5. [Fermat primality test](https://www.khanacademy.org/computing/computer-science/cryptography/random-algorithms-probability/v/fermat-primality-test-prime-adventure-part-10)
6. [Lucas pseudoprime](https://en.wikipedia.org/wiki/Lucas_pseudoprime)
7. [Baillie-PWS Primality Test](https://mathworld.wolfram.com/Baillie-PSWPrimalityTest.html)
8. [Strong Pseudoprime](https://mathworld.wolfram.com/StrongPseudoprime.html)
