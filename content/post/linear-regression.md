+++
author = "Tuffy"
title = '线性回归'
date = 2024-12-05T12:32:24+08:00
math = true 
draft = false
comments = true
toc = true
description = "{线性回归}"

+++

线性回归理解起来很简单，当参数样本只有一个特征与一个标签时就是高中学最小二乘法的那个模型，两个特征及以上时则是大学概率论上学的参数估计，只不过估计参数时选用了梯度下降法。

值得注意的是损失函数的选择，为什么是**平方误差函数**而不是差的绝对值呢？这里其实是因为我们假定偏差的分布符合高斯分布，其对应的概率密度函数就是平方和形式，而且**差的绝对值**（拉普拉斯分布）所对应的极大似然函数没法方便寻优。

而说到梯度下降，在工程实现方面其实也不难，具体见后文代码，主要谈谈背后的数学。

### 0.梯度是什么

对于一个多元函数，梯度是该函数所有偏导数的向量，指向函数增长最快的方向，大小表示增长的速率。

假设有一个函数 $f(x_1,x_2,…,x_n)$，梯度可以表示为一个向量：

$\nabla f = \left( \frac{\partial f}{\partial x_1}, \frac{\partial f}{\partial x_2}, \dots, \frac{\partial f}{\partial x_n} \right)$

这里，$\frac{\partial f}{\partial x_i}$ 是函数 $f$ 对每个变量 $x_i$ 的偏导数。

### 1. **梯度下降与泰勒展开的关系**

先解释下为什么要有梯度下降法：其实最简单的二维凸函数是抛物线 $f(x)=x^2$，很容易通过解方程 $f'(x)=0$ 求出最小值在 $x=0$ 处；只是有一些凸函数这样解方程太麻烦，便用梯度下降法来找最值。

最简单情况 ($f(x)=x^2$) 下，若将给定点 $x_0$ 加上 -$\eta\nabla f(x_0)$，就相当于一个逐渐靠近最低点的物理过程。比如取 $x_0$ = 10，$\eta = 0.2$ ，迭代 10 次左右就是差不多靠近了最低点 $x=0$

>初始值 x=10, f(x)=100<br>
>第 1 次迭代: x=6.0, f(x)=36.0<br>
>第 2 次迭代: x=3.6, f(x)=12.96<br>
>第 3 次迭代: x=2.16, f(x)=4.67<br>
>第 4 次迭代: x=1.296, f(x)=1.68<br>
>第 5 次迭代: x=0.7776, f(x)=0.60<br>
>第 6 次迭代: x=0.46656, f(x)=0.22<br>
>第 7 次迭代: x=0.27994, f(x)=0.078<br>
>第 8 次迭代: x=0.16796, f(x)=0.028<br>
>第 9 次迭代: x=0.10078, f(x)=0.0102<br>
>第 10 次迭代: x=0.06047, f(x)=0.00366<br>

不难理解这个 $\eta$ （取个名字叫学习率）取得太小则迭代次数过多，太大则会越过最低点不断震荡。

在实际中的函数没有这么简单，更复杂的算式里梯度下降沿着函数的梯度方向进行优化，而梯度本身就是函数在当前位置的泰勒展开的一阶导数。此时梯度下降法可以看作是在每一步使用泰勒展开的一阶近似来更新参数。具体来说，梯度下降可以看作是泰勒展开的一阶近似的迭代应用。

### 2. **梯度下降的解析解**

梯度下降法**没有解析解**，因为它是一种数值优化方法。（解析解和数值解的区别不知道的话何承春会头疼呢）

返回的结果就是使目标函数取得最小值的**参数**。

接下来通过一个具体的例子来解释如何理解**批量梯度下降**、**小批量梯度下降**和**随机梯度下降**。

假设我们有一个三元函数（即目标函数是有三个参数的函数），并且我们有100个样本数据。我们希望通过这些数据来优化函数的参数。

##### 1. 全批量梯度下降（Batch Gradient Descent）

全批量梯度下降会在每次迭代中使用**所有 100 个样本**来计算梯度，并更新参数。

> 1. 从初始参数 $\theta_0$ 开始。<br>
> 2. 计算所有 100 个样本点的梯度，得到梯度平均值：<br>
$\nabla f(\theta) = \frac{1}{100} \sum_{i=1}^{100} \nabla f_i(\theta)$<br>
> 3. 更新参数：<br>
> $\theta_{k+1} = \theta_k - \eta \nabla f(\theta_k)$

##### 2. 随机梯度下降（Stochastic Gradient Descent, SGD）

与全批量梯度下降不同，随机梯度下降每次迭代仅使用**一个样本**来计算梯度并更新参数。对于 100 个样本，随机梯度下降每次只随机选择一个样本进行参数更新。

> 1. 从初始参数 $\theta_0$  开始。<br>
> 2. 随机选择一个样本 $i$，计算该样本的梯度 $\nabla f_i(\theta)$。<br>
> 3. 使用该样本的梯度更新参数：<br>
>$\theta_{k+1} = \theta_k - \eta \nabla f_i(\theta_k)$
> 4. 重复这个过程，直到达到停止条件

##### 3. 小批量梯度下降（Mini-batch Gradient Descent）

小批量梯度下降是全批量梯度下降和随机梯度下降的折中方法。每次迭代时，它会从所有 100 个样本中随机选取一个**小批量**（例如，10 个样本），用这些样本来计算梯度并更新参数。

> 1. 从初始参数 $\theta_0$ 开始。<br>
> 2. 随机选择一个小批量样本，假设这个小批量包含 10 个样本 $i_1, i_2, \dots, i_{10}$。<br>
> 3. 计算这 10 个样本的梯度平均值：<br>
>$\nabla f(\theta) = \frac{1}{10} \sum_{i=1}^{10} \nabla f_{i}(\theta)$
> 4. 使用这个平均梯度更新参数：<br>
>$\theta_{k+1} = \theta_k - \eta \nabla f(\theta_k)θk+1=θk−η∇f(θk)$
> 5. 重复这个过程，直到达到停止条件。



代码实现：

```python
#%%
import random
import torch
from matplotlib import pyplot as plt
#%%
# 生成随机权重
def get_w_b():
    w = torch.normal(0, 0.01, size=(1, 1), requires_grad=True)
    b = torch.zeros(1, requires_grad=True)
    return w, b
#%%
# 生成训练数据
def synthetic_data(w, b, num_examples):
    """生成y = Xw + b + 噪声"""
    X = torch.normal(0, 1, (num_examples, len(w)))
    y = torch.matmul(X, w) + b
    y += torch.normal(0, 0.01, y.shape)
    return X, y.reshape((-1, 1))
#%%
# 生成迭代器
def data_iter(batch_size, features, labels):
    num_examples = len(features)
    indices = list(range(num_examples))
    # 这些样本是随机读取的，没有特定的顺序 
    random.shuffle(indices)
    for i in range(0, num_examples, batch_size):
        batch_indices = torch.tensor(indices[i:min(i + batch_size, num_examples)])
        yield features[batch_indices], labels[batch_indices]
#%% md
## 自己完成
#%%
# 线性回归模型
def Linear_regression(X, w, b):
    return X*w + b #return torch.matmul(X, w) + b更好
#%%
# 均方损失函数 MSE
def squared_loss(y_hat, y):
    return (y_hat - y) ** 2 / 2
#%%
# 优化器实现
def sgd(params, lr, batch_size):
    """小批量随机梯度下降"""
    with torch.no_grad():
        for param in params:
            param -= lr * param.grad / batch_size
            param.grad.zero_()
#%% md
## 测试
#%%
lr = 0.03
num_epochs = 3
net = Linear_regression
loss = squared_loss
batch_size = 10

# y = -5 * x + 0.1 
true_w = torch.tensor([-5.0])
true_b = 0.1
features, labels = synthetic_data(true_w, true_b, 1000)
w, b = get_w_b()

for epoch in range(num_epochs):
    for X, y in data_iter(batch_size, features, labels):
        l = loss(net(X, w, b), y)
        l.sum().backward()
        sgd([w, b], lr, batch_size)
    with torch.no_grad():
        train_l = loss(net(features, w, b), labels)
        print(f'epoch {epoch + 1}, loss {float(train_l.mean()):f}')
        print(f'w {w}, b {b}')
#%% md
## 绘制图形
#%%
x = features
y = labels

plt.scatter(x, y, label='Samples')
plt.plot(x.detach().numpy(), w.detach().numpy() * x.detach().numpy() + b.detach().numpy(), c='r', label='True function')
plt.plot(x, -5 * x + 0.1, c='b', label='Trained model')
plt.title("CQUPT2022212062", loc="center")
plt.legend()
plt.show()
#%%
print("权重",w)
#%%

```

