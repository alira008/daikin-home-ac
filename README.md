# Raspberry PI Daikin Home AC

I have pets at home and we are not home all the time. Especially during the summer
we don't want to leave the AC on all day.

I decided to use an old Raspberry PI I had to create a server and user interface to be able to
control my Daikin air conditioner.

I save state of the Daikin AC unit internally in a Sqlite table and encode commands that the user 
sends to the server. The server encodes 3 message frames into bits, then converts them into 
frequencies that the IR LED should turn on so that the AC understands on and off state pulse. The
output is saved onto a lircd file that can be executed by lircd on the raspberry pi so that it 
sends the signal to the Daikin AC.

## Stack
- GO, lirc, and Sqlite3 for the backend
- Sveltekit with Tailwindcss for the frontend

## References
[Daikin IR Reverse Protocol](https://github.com/blafois/Daikin-IR-Reverse)
