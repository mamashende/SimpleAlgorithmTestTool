import numpy as np

import matplotlib.pyplot as plt

def read_results(file_path):
    with open(file_path, 'r') as file:
        lines = file.readlines()
    
    data_sizes = []
    execution_times = []
    for line in lines:
        if "Data size:" in line:
            data_size = int(line.split(":")[1].strip())
            data_sizes.append(data_size)
        if "Execution time in nanoseconds:" in line:
            time_ns = int(line.split(":")[1].strip().split()[0])
            execution_times.append(time_ns)
    
    return data_sizes, execution_times

def calculate_averages(data_sizes, execution_times, group_size=20):
    avg_data_sizes = []
    avg_execution_times = []
    for i in range(0, len(execution_times), group_size):
        group_times = execution_times[i:i + group_size]
        group_sizes = data_sizes[i:i + group_size]
        if len(group_times) == group_size:
            avg_time = sum(group_times) / group_size
            avg_size = sum(group_sizes) / group_size
            avg_execution_times.append(avg_time)
            avg_data_sizes.append(avg_size)
    return avg_data_sizes, avg_execution_times

def plot_results(avg_data_sizes, avg_execution_times):
    plt.plot(avg_data_sizes, avg_execution_times, marker='o', label='Measured')

    # Adding asymptotic complexity functions
    sizes = np.array(avg_data_sizes)
    plt.plot(sizes, np.log(sizes), label='O(log n)')
    plt.plot(sizes, sizes, label='O(n)')
    plt.plot(sizes, sizes * np.log(sizes), label='O(n log n)')
    plt.plot(sizes, sizes**2, label='O(n^2)')
    plt.plot(sizes, sizes**3, label='O(n^3)')
    #plt.plot(sizes, 2**sizes, label='O(2^n)')
    plt.xlabel('Data Size')
    plt.ylabel('Average Execution Time (ns)')
    plt.title('Algorithm Execution Time vs Data Size')
    plt.legend()
    plt.grid(True)
    plt.show()

if __name__ == "__main__":
    file_path = 'result.txt'
    data_sizes, execution_times = read_results(file_path)
    avg_data_sizes, avg_execution_times = calculate_averages(data_sizes, execution_times)
    plot_results(avg_data_sizes, avg_execution_times)