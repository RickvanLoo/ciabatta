# ciabatta
Golang Command-Line-Interface for Bread Baking


## What?
Ciabatta is (very opinionated) simple command-line-interface to assist the process of bread baking and managing recipes. It is not dependent on any external libraries, making it quite low maintanance for the author and future-proof. Ciabatta keeps track of the list of ingredients, scales them according to the amount of flour, and keeps track of the baking steps.

Write down ingredients and steps as you go and save for later use, or open a recipe, start it, and follow the steps in the kitchen!


## Why?
When baking bread and studying the art of baking bread, I found out that classic bakers tend to use 'bakers percentages'. This makes all the ingredients a factor of the amount of flour and possible to easily scale recipes to your requirements. I do not like writing recipes down on paper, so I wanted to make an application where I can scale recipes of others instantly and keep track of my own. 

This means everything is in metric and in grams: even liquids. Weighing liquids is easier and more accurate, no measuring cups: just a scale and a bowl for all ingredients. In future I might add conversion formulas for commonly used liquids as olive oil (instead of using the 1ml oil = 1 gr simplification). Cups, tablespoons, teaspoons, or any imperial unit will never make the list. 

In case of the baking process I found out that there are various steps (autolyse, fermenting, rising, resting) that all come to the general idea of letting your pre-bread sit on the counter and set a timer before you come back. This has been simplified. Hence the following bread baking 'functions' can be used in various permutations: 

* Mixing
* Resting
* Knead
* Shape
* Bake

## Features
* Only been tested on Linux, might function elsewhere.
* Add and list ingredients
* Edit recipes
* Save and open recipes
* Scale ingredients, by editing the amount of flour

## Planned
* Add processes based upon the 5 'functions'
* Run recipes according to the process
* Start a timer (with an alarm sound/notification)
* Config