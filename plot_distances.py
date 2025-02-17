import matplotlib.pyplot as plt
import numpy
import json

def sort_distances_by_key(distances):
    """
    Sort a dictionary of distances by their numeric keys in ascending order.
    
    Args:
        distances (dict): Dictionary with numeric keys and distance values
        
    Returns:
        dict: Sorted dictionary with numeric keys in ascending order
    """
    return dict(sorted(distances.items(), key=lambda x: int(x[0])))

def plot_distances(distances):
    """
    Plot the distances from the sorted dictionary.

    Args:
        distances (dict): Sorted dictionary with numeric keys and distance values
    """
    x = list(distances.keys())
    y = list(distances.values())

    plt.figure(figsize=(10, 6))
    plt.plot(x, y, 'o-', markersize=5, linewidth=1, color='blue')
    plt.title('Distances Over Generations')
    plt.xlabel('Generation')
    plt.ylabel('Distance')
    plt.grid(True)
    plt.xticks(rotation=45, ha='right', fontsize=8)
    
    plt.tight_layout()
    plt.show()


if __name__ == "__main__":
    with open('results/pr/distances.json', 'r') as f:
        data = json.load(f)
        distances = data['distances']
        
        # Sort the distances
        sorted_distances = sort_distances_by_key(distances)
        plot_distances(sorted_distances)

        