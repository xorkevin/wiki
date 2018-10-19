---
title: Composition and Dependency Injection
author: Kevin Wang
date: 2018-10-19
---

# Composition Over Inheritance

## Inheritance

We are running a transportation company:

\lstset{language=Java,basicstyle={\scriptsize}}

::: {.columns}
::: {.column width="50%"}

```
class Car {
  Gasoline g;
  GasEngine e;

  Car() {
    this.g = new Gasoline();
    this.e = new GasEngine();
  }

  refuel(Gasoline g);
  drive(Destination d) {
    fourStrokeCombustion(this.e);
  }
  carryPassenger(Human h);
}
```

:::
::: {.column width="50%"}

```
class Truck {
  Gasoline g;
  GasEngine e;

  Truck() {
    this.g = new Gasoline();
    this.e = new GasEngine();
  }

  refuel(Gasoline g);
  drive(Destination d) {
    fourStrokeCombustion(this.e);
  }
  carryCargo(Cargo c);
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
class Vehicle {
  Gasoline g;
  GasEngine e;

  Vehicle() {
    this.g = new Gasoline();
    this.e = new GasEngine();
  }

  refuel(Gasoline g);
  drive(Destination d) {
    fourStrokeCombustion(this.e);
  }
}
```

:::
::: {.column width="50%"}

```
class Car extends Vehicle {
  carryPassenger(Human h);
}

class Truck extends Vehicle {
  carryCargo(Cargo c);
}
```

:::
:::

::: notes

- inheritance may appear to solve the problem
- code deduplication
  - both are vehicles, so they both have gas engines that need to be refueled
- the issues arise when these classes are used and extended

:::

## Inheritance Usage

Say we wanted to have a transportation hub with a fleet of vehicles:

\lstset{language=Java,basicstyle={\scriptsize}}

```
class Hub {
  Car[] carfleet;
  Truck[] truckfleet;

  transportPassenger(Human h, Destination d) {
    Car c = carfleet.pop();
    c.carryPassenger(h);
    c.drive(d);
  }

  transportCargo(Cargo c, Destination d) {
    Truck t = truckfleet.pop();
    t.carryCargo(c);
    t.drive(d);
  }
}
```

## Inheritance Usage

Some time passes, electric vehicles are now a thing:

\lstset{language=Java,basicstyle={\scriptsize}}

```
class ElectricCar {
  Battery b;
  ElectricEngine e;

  ElectricCar() {
    this.b = new Battery();
    this.e = new ElectricEngine();
  }

  charge(Battery b);
  drive(Destination d) {
    emInduction(this.e);
  }
  carryPassenger(Human h);
}
```

## Inheritance Usage

Is it possible to include this ElectricCar in the carfleet?

\lstset{language=Java,basicstyle={\scriptsize}}

::: {.columns}
::: {.column width="50%"}

```
class Hub {
  Car[] carfleet;
  Truck[] truckfleet;

  transportPassenger(
    Human h, Destination d
  ) {
    Car c = carfleet.pop();
    c.carryPassenger(h);
    c.drive(d);
  }

  transportPassenger(
    Human h, Destination d
  ) {
    Car c = carfleet.pop();
    c.carryPassenger(h);
    c.drive(d);
  }
}
```

:::
::: {.column width="50%"}

```
class ElectricCar {
  Battery b;
  ElectricEngine e;

  ElectricCar() {
    this.b = new Battery();
    this.e = new ElectricEngine();
  }

  charge(Battery b);
  drive(Destination d) {
    emInduction(this.e);
  }
  carryPassenger(Human h);
}
```

:::
:::

::: notes

- main issue is that our original model did not account for the fact that we
  would need a general 'PassengerTransporter' abstraction
- the danger of inheritance is that it forces you to plan for the future

:::

## Composition

\lstset{language=Java,basicstyle={\scriptsize}}

::: {.columns}
::: {.column width="50%"}

```
interface Engine {
  doWork();
}

class GasEngine
  implements Engine {
  Gasoline g;
  refuel(Gasoline g);
  doWork() {
    fourStrokeCombustion(this);
  }
}

class ElectricEngine
  implements Engine {
  Battery b;
  charge(Battery b);
  doWork() {
    emInduction(this);
  }
}
```

:::
::: {.column width="50%"}

```
class Car {
  Engine e;

  Car(Engine e) {
    this.e = e;
  }

  drive(Destination d) {
    this.e.doWork();
  }
  carryPassenger(Human h);
}

Car gasolineCar =
  new Car(new GasEngine());

Car electricCar =
  new Car(new ElectricEngine());
```

:::
:::

::: notes

- only constraint that matters is that the interface is fulfilled
- relaxing all the constraints about what type something inherits from
- focused only on the capabilities of the object
- not explicitly written, but it can be inferred that we can rewrite Truck in
  much the same way

:::

# Dependency Injection
