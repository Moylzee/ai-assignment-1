import matplotlib.pyplot as plt
import json

class City:
    def __init__(self, ID, X, Y):
        self.ID = ID
        self.X = X
        self.Y = Y

def plot_tour(cities, tour):
    x = [cities[i].X for i in tour] + [cities[tour[0]].X]
    y = [cities[i].Y for i in tour] + [cities[tour[0]].Y]

    plt.figure(figsize=(10, 6))
    plt.plot(x, y, 'o-', markersize=5, linewidth=1, color='blue')
    plt.title('Salesman Tour')
    plt.xlabel('X Coordinate')
    plt.ylabel('Y Coordinate')
    plt.grid(True)
    plt.show()

if __name__ == "__main__":
    # Load cities and best tour from JSON file
    with open('results/pr/best_tour.json', 'r') as f:
        data = json.load(f)
        cities = [City(**city) for city in data['cities']]
        best_tour = data['best_tour']

    plot_tour(cities, best_tour)