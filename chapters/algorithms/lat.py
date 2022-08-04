from time import time_ns

def calculate_latency(self, client, content):
    start_time = time_ns()
    client.perform_request(content)
    end_time = time_ns()
    print(f"{end_time} {end_time - start_time}")
