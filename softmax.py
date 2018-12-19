#!/usr/bin/env python

import numpy as np

scores = [-3.0, 1.0, 0.2]

def softmax(x):
    return np.exp(x)/np.sum(np.exp(x), axis=0)

print(softmax(scores))

import matplotlib.pyplot as plt

