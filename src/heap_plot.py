import sys
import matplotlib.pyplot as plt
import mpld3


if len(sys.argv) < 2:
    print("Please provide path to log file")
    sys.exit(1)

log_file_path = sys.argv[1]

heap_allocs = []
heap_released = []
heap_idle = []
heap_sys = []
stack_sys = []
sys_values = []
counter = []

with open(log_file_path, 'r') as log_file:
    for i, line in enumerate(log_file):
        if "HeapAlloc" in line:
            alloc_str = line.split("HeapAlloc = ")[1]
            heap_alloc = int(alloc_str.split(" MB")[0])
            
            heap_allocs.append(heap_alloc)

        elif "src/corso.go:77	Sys =" in line:
            alloc_str = line.split("Sys = ")[1]
            sys_value = int(alloc_str.split(" MB")[0])
        
            sys_values.append(sys_value)

        elif "HeapReleased" in line:
            alloc_str = line.split("HeapReleased = ")[1]
            rel = int(alloc_str.split(" MB")[0])
            
            heap_released.append(rel)

        elif "HeapIdle" in line:
            alloc_str = line.split("HeapIdle = ")[1]
            idle = int(alloc_str.split(" MB")[0])
            
            heap_idle.append(idle)
        
        elif "StackSys" in line:
            alloc_str = line.split("StackSys = ")[1]
            st = int(alloc_str.split(" MB")[0])
            
            stack_sys.append(st)

        elif "HeapSys" in line:
            alloc_str = line.split("HeapSys = ")[1]
            hs = int(alloc_str.split(" MB")[0])
            
            heap_sys.append(hs)

        
# Plot the data
counter = list(range(1, len(heap_allocs) + 1))
plt.plot(counter, heap_allocs, marker='.', label='HeapAlloc')
plt.plot(counter, heap_released, marker='.', label='HeapReleased')
plt.plot(counter, heap_idle, marker='.', label='HeapIdle')
#plt.plot(counter, stack_sys, marker='.', label='StackSys')
plt.plot(counter, sys_values, marker='.', label='Sys')
plt.plot(counter, heap_sys, marker='.', label='HeapSys')
plt.xlabel('Seconds')
plt.ylabel('Memory (MB)')
plt.title('HeapAlloc and Sys over time')
plt.grid(True)
plt.legend()
plt.savefig('graph.png')

# interactive_plot = mpld3.fig_to_html(plt.figure())
# mpld3.save_html(plt.figure(),"index.html")

plt.show()