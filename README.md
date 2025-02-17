# AI Assignment 1 
## Brian Moyles - 21333461

<details>
<summary>Assignment Description</summary>
<br>
- In this assignment, you will implement a genetic algorithm to solve the Traveling Salesman Problem (TSP)
</details>

## How To Run
- Clone the Repo Locally 
- Navigate to the Repo (ai-assignment-1) in your Terminal
- run `go run main.go`

## How to change which file to run
- Navigate to `main.go`
- On `line 28`, change the `filename` variable to reference either:
    - berlin (berlin)
    - kroA (kr)
    - pr (pr)
- (Note: The Variable Names can be found on lines 20-22 of main.go)

## How to Visualize Plots
- After a runtime, the results will be saved in /results
- To visualize your results, run any of the python scripts

# Files 
<details>
<summary>main.go</summary>
This is the main entry point for the genetic algorithm. 
It initializes the parameters, reads the TSP file, runs the genetic algorithm, and saves the results.
</details>

<details>
<summary>utilities.go</summary>
Contains utility functions such as reading TSP files and checking if a city is in a tour.
</details>

<details>
<summary>selection.go</summary>
Implements the selection mechanism for the genetic algorithm, specifically the tournament selection method.
</details>

<details>
<summary>crossover.go</summary>
Implements the crossover mechanisms for the genetic algorithm, including Ordered Crossover (OX) and Partially Mapped Crossover (PMX).
</details>

<details>
<summary>mutations.go</summary>
Implements mutation mechanisms for the genetic algorithm, including Swap Mutation and Inversion Mutation.
</details>

<details>
<summary>plot_distances.py</summary>
Plots the distances over generations using matplotlib. It reads the distances from a JSON file and sorts them before plotting.
</details>

<details>
<summary>plot_tour.py</summary>
Plots the best tour found by the genetic algorithm using matplotlib. It reads the cities and the best tour from a JSON file.
</details>

<details>
<summary>plot_fitnesses.py</summary>
Plots the fitness values over generations using matplotlib. It reads the fitness values from a JSON file and sorts them before plotting.
</details>