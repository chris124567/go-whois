#!/usr/bin/env python
import socket

def recv_all(socket):
	total_data=[]
	while True:
		data = socket.recv(2048) #recv in blocks of 2048
		if not data: break #until connection closes
		total_data.append(data)
	return ''.join(total_data)

def query_whois(tld):
	ip = '192.0.32.59' #iana server
	port = 43 #whois port
	message = 'nic.' + tld + "\r\n" #all tlds have a nic.tld domain as far as I can tell

	s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
	s.connect((ip, port))
	s.sendall(message)
	data = recv_all(s)
	s.close()
	return data

fp = open('tlds', 'r') #download this file from https://data.iana.org/TLD/tlds-alpha-by-domain.txt
line = fp.readline()
cnt = 1
while line: #go line by line
	tld = line.strip() #the name of the tld, each separated by lines in the file
	data = query_whois(tld)
	f = open("tld_whois/" + tld, "aw") #download directory
	f.write(data) #write to file
	f.close()
	print("Downloaded WHOIS data for TLD %s (%d of 1527)" % (tld, cnt))
	line = fp.readline()
	cnt += 1

fp.close()
