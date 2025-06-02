import socket
import threading

def handle_data(conn):
    with conn:
        while True:
            data = conn.recv(512)
            if not data:
                break
            print(data.decode(), end='')

def main():
    host = '0.0.0.0'
    port = 8080
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
        s.bind((host, port))
        s.listen()
        print(f"Listening on {host}:{port}")
        while True:
            conn, addr = s.accept()
            t = threading.Thread(target=handle_data, args=(conn,))
            t.start()

if __name__ == "__main__":
    main()
