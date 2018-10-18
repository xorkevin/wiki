---
title: Composition and Dependency Injection
author: Kevin Wang
date: 2018-10-19
---

# Composition Over Inheritance

## Inheritance

We are running a medieval farm:

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
  feed(Food f) {
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
- this example feels contrived because it is, once you think with composition,
  you can never go back

:::

## Inheritance Usage

Our medieval world:

\lstset{language=Java,basicstyle={\scriptsize}}

::: {.columns}
::: {.column width="50%"}

```
class Cart extends Payload {
  move(Animal a) {
    // use animal to pull cart
    a.pull(this);
  }
}

class AnimalCart {
  Animal a;
  Cart c;
  Food f;
  AnimalCart(Animal a) {
    this.a = a;
    this.c = new Cart();
    this.f = new Apple();
  }
  move() {
    this.a.feed(this.f);
    this.c.move(this.a);
  }
}
```

:::
::: {.column width="50%"}

```
AnimalCart(new Horse()).move();

AnimalCart(new Cow()).move();
```

:::
:::

## Inheritance Usage

Some time passes, and the industrial revolution occurs:

\lstset{language=Java,basicstyle={\scriptsize}}

::: {.columns}
::: {.column width="50%"}

```
class Train {
  Train() {
    // initialize train
  }
  pull(Payload p) {
    // train pulls payload
  }
  refuel(Coal c) {
    // give train coal
  }
}
```

:::
::: {.column width="50%"}

```
// cannot do
AnimalCart(new Train()).move();
```

:::
:::

::: notes

- main issue is that our original model did not account for the fact that we
  would need a general 'Engine' abstraction

:::

## Composition

\lstset{language=Java,basicstyle={\scriptsize}}

::: {.columns}
::: {.column width="50%"}

```
interface Engine {
  pull(Payload payload);
  refuel();
}

class Cart extends Payload {
  move(Engine e) {
    // use engine to pull cart
    e.pull(this);
  }
}

class DrivenCart {
  Engine e;
  Cart c;
  DrivenCart(Engine e) {
    this.e = e;
    this.c = new Cart();
  }
  move() {
    this.e.refuel();
    this.c.move(this.e);
  }
}
```

:::
::: {.column width="50%"}

```
class Horse implements Engine {
  Food f;
  Horse() {
    this.f = new Apple();
  }
  pull(Payload payload) {
    // pull payload
  }
  refuel() {
    // consume this.f
  }
}

class Train implements Engine {
  Fuel f;
  Train() {
    this.f = new Coal();
  }
  pull(Payload payload) {
    // pull payload
  }
  refuel() {
    // consume this.f
  }
}
```

:::
:::

## Composition Usage

\lstset{language=Java,basicstyle={\scriptsize}}

```
// equally valid
DrivenCart(new Horse()).move();
DrivenCart(new Train()).move();
```

::: notes

- only constraint that matters is that the interface is fulfilled
- relaxing all the constraints about what type something inherits from
- focused only on the capabilities of the object

:::

# Dependency Injection
