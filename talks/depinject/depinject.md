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
class Horse {
  eat(Food f) {
    // eat food
  }
  pull(Payload p) {
    // pull payload
  }
}
```

:::
::: {.column width="50%"}

```
class Dog {
  eat(Food f) {
    // eat food
  }
  bark() {
    // bork
  }
}
```

:::
:::

## Inheritance

Abstract away common implementations

\lstset{language=Java,basicstyle={\scriptsize}}

::: {.columns}
::: {.column width="50%"}

```
class Animal {
  eat(Food f) {
    // eat food
  }
}
```

:::
::: {.column width="50%"}

```
class Horse extends Animal {
  pull(Payload p) {
    // pull payload
  }
}

class Dog extends Animal {
  bark() {
    // bork
  }
}
```

:::
:::

::: notes

- inheritance may appear to solve the problem
- code deduplication
  - both can eat food; abstract away implementation
- the issues arise when these classes are used and extended
- this example feels contrived because it is, once you think with composition,
  you can never go back

:::

## Inheritance Usage

Our medieval world:

\lstset{language=Java,basicstyle={\scriptsize}}

```
class Carriage extends Payload {
  transport(Human h) {
    // transports a human
  }
}

class HorseCarriage extends Horse {
  Food f;
  Carriage c;
  HorseCart() {
    this.f = new Apple();
    this.c = new Carriage();
  }
  move(Human h) {
    this.eat(this.f);
    this.c.transport(h);
    this.pull(this.c);
  }
}
```

## Inheritance Usage

Some time passes, and the industrial revolution occurs:

\lstset{language=Java,basicstyle={\scriptsize}}

```
class Train {
  Carriage c;
  Train() {
    this.c = new Carriage();
  }
  move(Human h) {
    this.c.transport(h);
    this.pull(this.c);
  }
}
```

## Inheritance Usage

Is it possible to build a general transportation hub?

\lstset{language=Java,basicstyle={\scriptsize}}

::: {.columns}
::: {.column width="50%"}

```
class TransportHub {
  ____ transporter;
  move(Human h) {
    this.transporter.move(h);
  }
}
```

:::
::: {.column width="50%"}

```
Human h = new Human();

HorseCarriage hc =
  new HorseCarriage();
hc.move(h);

Train t = new Train();
t.move(h);
```

:::
:::

::: notes

- main issue is that our original model did not account for the fact that we
  would need a general 'Transporter' abstraction
- the danger of inheritance is that it forces you to plan for the future

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
