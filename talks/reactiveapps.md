---
title: Reactive Apps
author: Kevin Wang
date: 2018-02-21
---

# A brief history of the web

::: notes
- motivation is to give a bit of context for client side web apps
- why has the industry settled on this methodology
- while writing this, young age of the modern web was put into perspective
:::

## Generation 1

Server rendered web pages

MVC becomes the de facto architecture

- 1995 - PHP
- 2005 - Ruby on Rails
- 2005 - Django

::: notes
- server responsible for rendering every page
- tight coupling between the backend and ui
  - monolithic
- MVC becomes the de facto architecture
:::

## Generation 2

Client side apps

::: {.columns}
::: {.column width="50%"}
#### Web server frameworks
- 2010 - Flask
- 2010 - Express.js
:::
::: {.column width="50%"}
#### Client side web frameworks
- 2010 - Backbone.js
- 2011 - AngularJS
- 2011 - Ember.js
:::
:::

::: notes
- web browsers became more powerful
  - we begin to think about the browser as a platform itself
  - js became ubiquitous
- work from the server was offloaded to the client as a result
- server side dynamic was changing
  - microservices, approx 2005 - Borg, 2010 - Mesos, 2012 - Eureka / Hystrix
- required the view to be separate from the server side logic
- idea of a SPA was created as entire web apps could be created from a single
  page
  - MVC architecture on the front end
  - client side routing
:::

## Generation 2 - Problems

- Traditional MVC does not scale on the client side
- State management is complex
  - the model and view may be desynchronized

::: notes
- MVC client side apps are difficult to maintain
  - do not scale
  - cannot use the same approach as with completely server rendered apps
  - not just a single output of html to the browser
- in particular state management becomes overly complex
  - need to ensure that the view reflects the model
:::

## Generation 2 - Problems

![mvc data flow](assets/reactiveapps/mvcdataflow.png)

::: notes
- a certain change in the model could have many side effects
- AngularJS was notorious for being non-performant, because of dirty checking
  to keep the model and view in sync
:::

# Generation 3

## React

- Kickstarted the componentization craze
  - Focused on reusability
- View should be a function of app state
  - Virtual DOM

::: notes
- no longer thinking in terms of templating or pages
  - breaking down an app into reusable components
- key principle is view is a pure function of the application state
  - implemented using a virtual dom
  - solves the issue that the view always needs to reflect the current model
:::

## React - Virtual DOM

![virtual dom](assets/reactiveapps/virtualdom.png)

::: notes
- virtual dom is a representation of the current dom in js
  - when a change is made, the vdom is updated
  - the vdom is diffed against its previous version
  - all the changed nodes get updated in the dom
- gives react a performance advantage
  - expensive dom manipulations are minimized
:::

## Redux

![redux data flow](assets/reactiveapps/reduxdataflow.png)

# Generation 4
