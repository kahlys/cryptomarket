# cryptomarket

[![godoc](https://godoc.org/github.com/kahlys/cryptomarket?status.svg)](https://godoc.org/github.com/kahlys/cryptomarket) 
[![build](https://api.travis-ci.org/kahlys/cryptomarket.svg?branch=master)](https://travis-ci.org/kahlys/cryptomarket)
[![go report](https://goreportcard.com/badge/github.com/kahlys/cryptomarket)](https://goreportcard.com/report/github.com/kahlys/cryptomarket)

Basic command line client for [Cryptocurrency Market Capitalizations](https://coinmarketcap.com/)

## Installation

With a correctly configured [Go toolchain](https://golang.org/doc/install)

```
go get -u github.com/kahlys/cryptomarket/cmd/cryptomarket
```

## Usage

Select how many cryptocurrency you want to analyze.

```
$ cryptomarket -top=10

       Currency               Price ($)    Total ($)         1 Week (%)   1 Day (%)    1 Hour (%)
----   --------------------   ----------   ---------------   ----------   ----------   ----------
1      Bitcoin                8399.58      141588041044      -8.61        -1.23        -0.73     
2      Ethereum               841.92       82114067746       -12.71       -1.80        -1.14     
3      Ripple                 1.04         40690122948       +12.89       +13.69       -2.16     
4      Bitcoin Cash           1229.46      20851964333       -1.81        -5.28        -0.61     
5      Cardano                0.40         10258823270       -8.15        +4.21        -0.44     
6      Litecoin               151.72       8368261655        +1.92        -1.88        -1.60     
7      Stellar                0.39         7240394116        -8.85        -0.03        -1.31     
8      NEO                    107.69       6999525000        -12.42       -6.19        -1.17     
9      EOS                    8.81         5813619930        -11.03       -4.33        -1.24     
10     IOTA                   1.84         5101883425        -5.92        -3.57        -0.96  
```
