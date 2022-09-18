import numpy as np
import matplotlib.pyplot as plt
# from mpltools import annotation

# some default font sizes for plots
plt.rcParams['font.size'] = 12
plt.rcParams['font.family'] = 'sans-serif'
plt.rcParams['font.sans-serif'] = ['Arial', 'Dejavu Sans']

def random_plot(X, Y, labels):
    X = np.array(X)
    Y = np.array(Y)


    fig = plt.figure(figsize=(20, 10))
    ax1 = fig.add_subplot(111)
    ax1.margins(0.1)
    ax1.grid(True)

    color = ['k', 'r', 'b', 'g', 'c', 'y']
    shape = ['-','.']
    labels = ["original Impact Earth Program", "New ImpactEarth (based on golang)", "New ImpactEarth with redis(based on golang)"]

    for i in range(len(X)):
        ax1.plot(X[i], Y[i], '%s%s-'%(color[i%6], shape[0]), 
            label=labels[i])


    ax1.set_xlabel('http requests in one times', fontsize=16)
    ax1.set_ylabel('time cost (ms)', fontsize=16)
    ax1.set_title('HTTP Request times', fontsize=16)
    ax1.legend(loc='best', fontsize=14)

    plt.savefig("../img/HTTP_Request_Time_random_with_valid_parameter.png")