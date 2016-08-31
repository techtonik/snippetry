#!/usr/bin/env python

import socket

listenon = "localhost:11111"
print("Listening on: " + listenon)
sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
addr = listenon.split(':')
addr = (addr[0], int(addr[1]))
sock.bind(addr)
sock.listen(1)
conn, remote = sock.accept()
remoteaddr = remote[0] + ":" + str(remote[1])
print("New connection from: %s" % remoteaddr)
