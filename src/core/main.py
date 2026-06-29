import os
import time
while True:
 main_terminal_shell = input("root@archiso ~ # ")
 if main_terminal_shell == "exit":
  break
 elif main_terminal_shell == "clear":
  if os.name == "nt":
   os.system("cls")
  else:
   os.system("clear")