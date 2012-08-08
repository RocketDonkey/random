# Google Hurdle Doodle Doer?
# If you suck at running hurdles I feel bad for you son,
# I got 99 problems but I can run through hurdles.
import win32com.client
import time

def CrushIt():
  for i in xrange(0, 1000):
    shell.SendKeys('{LEFT}')
    shell.SendKeys('{RIGHT}')
    i += 1

# Make the shell
shell = win32com.client.Dispatch("WScript.Shell")

# Activate the Doodle (not brittle at all)
shell.AppActivate('Google - Google Chrome')

# Important - you have to CLICK the doodle while this piece of trash sleeps
time.sleep(0.1)
CrushIt()
