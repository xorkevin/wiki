---
title: Composition and Dependency Injection
author: Kevin Wang
date: 2018-10-19
---

# Composition Over Inheritance

## Inheritance

\lstset{language=Java,basicstyle={\scriptsize}}

::: {.columns}
::: {.column width="50%"}

```
class Animal {
  Animal(numLegs) {
    // initialize legs
  }
  pull(Payload p) {
    // use legs to pull payload
  }
  eat(Food f) {
    // consume food
  }
}
```

:::
::: {.column width="50%"}

```
class Horse extends Animal {
  Horse() {
    super(4);
  }
}

class Cow extends Animal {
  Cow() {
    super(4);
  }
  moo() {
    // print moo
  }
}
```

:::
:::

::: notes

- inheritance may appear to solve the problem
- code deduplication
  - initialize legs, walking, consuming food are shared implementation details
- the issues arise when these classes are used

:::

## Inheritance Usage

\lstset{language=Java,basicstyle={\scriptsize}}

::: {.columns}
::: {.column width="50%"}

Our medieval world:

```
class Payload {}

class Cart extends Payload {
  move(Animal a) {
    // use animal to pull cart
    a.pull(this);
  }
}

Cart cart = new Cart();
Horse horse = new Horse();
cart.move(horse);
```

:::
::: {.column width="50%"}

Some time passes, and the industrial revolution occurs:

```
class Train {
  Train() {
    // initialize train
  }
  pull(Payload p) {
    // train pulls payload
  }
  fuel(Fuel f) {
    // give train coal
  }
}

Cart cart = new Cart();
Train train = new Train();

// cannot do:
cart.move(train);
```

:::
:::

# Dependency Injection
