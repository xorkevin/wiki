# Neural Networks

## Gradients

\begin{align}
y &= W \cdot x \\
\frac{\partial z}{\partial W} &= \frac{\partial z}{\partial y} \frac{\partial y}{\partial W} \\
&= \frac{\partial z}{\partial y} \cdot X^{T} \\
\frac{\partial z}{\partial x} &= \frac{\partial z}{\partial y} \frac{\partial y}{\partial x} \\
&= W^{T} \cdot \frac{\partial z}{\partial y}
\end{align}
