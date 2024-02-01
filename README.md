# transductive-experimental-design

[![Build](https://github.com/h-waldschmidt/transductive-experimental-design/actions/workflows/build.yml/badge.svg)](https://github.com/h-waldschmidt/transductive-experimental-design/actions/workflows/build.yml)
[![Test](https://github.com/h-waldschmidt/transductive-experimental-design/actions/workflows/test.yml/badge.svg)](https://github.com/h-waldschmidt/transductive-experimental-design/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/h-waldschmidt/transductive/branch/main/graph/badge.svg?token=CJ4UBDX0C8)](https://codecov.io/gh/h-waldschmidt/transductive)
[![Go Report Card](https://goreportcard.com/badge/github.com/h-waldschmidt/transductive)](https://goreportcard.com/report/github.com/h-waldschmidt/transductive)

Transductive experimental design (TED) by Kai Yu et. al. [[1]](#1) selects the most informative points from a dataset to solve a regression problem.
The data doesn't need to be labeled, meaning that TED can be used for active learning.

This library currently only supports the sequential version of TED (the alternating version is still a work in progress). 

Additionally, an implementation of k-means++ is provided for comparison.
In future updates, implementations of various optimal design algorithms will be provided for comparison.

Examples can be found in [examples](https://github.com/h-waldschmidt/transductive/tree/main/examples)

## References

<a id="1">[1]</a> 
Yu Kai, Jinbo Bi, and Volker Tresp.
"Optimized placement of charging stations for electric cars"
Proceedings of the 23rd international conference on Machine learning. 2006.
