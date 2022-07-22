from multiprocessing import Process, Value

SHARED_MEM_REQUEST_COUNTER = Value('i', 0)

p = Process(target=throughput, args=[SHARED_MEM_REQUEST_COUNTER])
p.start()

def throughput(self, *args):
    current_throughput = args[0]
    previous_throughput = 0
    while True:
        time.sleep(1)
        thr = current_throughput.value - previous_throughput
        previous_throughput = current_throughput.value
        print(f"{time.time_ns()} {thr}")

p.join()
