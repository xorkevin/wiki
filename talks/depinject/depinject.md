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

  transportCargo(
    Cargo c, Destination d
  ) {
    Truck t = truckfleet.pop();
    t.carryCargo(c);
    t.drive(d);
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
  would need to be able to swap out the Engine
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

- do not extend to obtain functionality
- 'compose' to obtain new functionality
  - creating things from parts is how we think in the real world
  - if you need to saw some material, you do not care if given a hacksaw
    extended from a hand tool, or a table saw extended from a shop tool
- only constraint that matters is that the interface is fulfilled
  - focusing only on the capabilities of the object
- not explicitly written, but it can be inferred that we can rewrite Truck in
  much the same way

:::

# Dependency Injection and Inversion of Control

## Non dependency injected

Is it possible to add a train?

\lstset{language=Java,basicstyle={\scriptsize}}

::: {.columns}
::: {.column width="50%"}

```
class Hub {
  Car[] carfleet;
  Truck[] truckfleet;

  transportPassenger(...);
  transportCargo(...);

  addCar() {
    carfleet.push(
      new Car(
        new ElectricEngine()
      )
    );
  }
  addTruck() {
    truckfleet.push(
      new Truck(
        new ElectricEngine()
      )
    );
  }
}
```

:::
::: {.column width="50%"}

```
class Train {
  carryPassenger(Human h);
  carryCargo(Cargo c);
  move(Destination d);
}
```

:::
:::

::: notes

- limiting factor here is the Car and Truck constraint
- what we really care about is the ability to transport passengers and cargo

:::

## Dependency Injection

\lstset{language=Java,basicstyle={\scriptsize}}

::: {.columns}
::: {.column width="50%"}

```
interface PassengerTransporter {
  carryPassenger(Human h);
  move(Destination d);
}
interface CargoTransporter {
  carryCargo(Cargo c);
  move(Destination d);
}
```

:::
::: {.column width="50%"}

```
class Hub {
  PassengerTransporter[] passf;
  CargoTransporter[] cargof;

  transportPassenger(pass, dest) {
    t = passf.pop();
    t.carryPassenger(pass);
    t.move(dest);
  }
  transportCargo(cargo, dest) {
    t = cargof.pop();
    t.carryCargo(cargo);
    t.move(dest);
  }

  addPassT(passT) {
    passf.push(passT);
  }
  addCargoT(cargoT) {
    cargof.push(cargoT);
  }
}
```

:::
:::

::: notes

- remove Car and Truck dependency
- focus on the actual capabilities of the dependencies
- let those dependencies be provided to you
  - injection

:::

## Dependency Injection

\lstset{language=Java,basicstyle={\scriptsize}}

::: {.columns}
::: {.column width="50%"}

```
class Train {
  carryPassenger(Human h);
  carryCargo(Cargo c);
  move(Destination d);
}
```

:::
::: {.column width="50%"}

```
class Horse {
  carryPassenger(Human h);
  move(Destination d);
  eat(Food f);
}

class BFR {
  carryPassenger(Human h);
  carryCargo(Cargo c);
  move(Destination d);
}
```

:::
:::

::: notes

- dependencies can be fulfilled by anything that implements the interface
- focus on the capabilities of the dependencies, not their implementation
- easy to provide newer implementations without having to change code
- note that providing an Engine to a Car in the previous example is also an
  example of dependency injection, except that was in the context of
  composition

:::
