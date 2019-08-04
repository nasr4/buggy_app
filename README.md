# buggy_app

At work I am mentoring a junior engineer on the value of using Linux's strace to debug application problems.

A useful use case to use it would be when an application is just hanging. Strace can show what the application is doing under the hood by revealing all of the system calls it is making to the kernel.

buggy_app.go attempts to read a file in a continuous loop. If the file doesn't exist in the current directory then an open file exception is thrown. The open file exception is being caught, but instead of presenting the user with an error I am simply continuing the loop to start the futile exercise over again.

That behavior mimics a piece of buggy software where an error is either not caught or caught in a way that doesn't log any useful information or exit the program with an error. That would cause the application to hang or not perform as expected.

Steps to reproduce:

In one terminal build the buggy app to get the executable binary "buggy_app" and then run that binary:

$ go build buggy_app.go

$ ls

buggy_app  buggy_app.go

$ ./buggy_app


In another terminal on the same system find the PID of the buggy app:

$ ps aux | grep buggy

nasr  23357 97.3  0.1 106220 13312 pts/0    Rl+  01:36   0:20 ./buggy_app

nasr  23367  0.0  0.0 105300   880 pts/1    S+   01:36   0:00 grep buggy

And then run strace on that PID:

$ strace -p 23357

...

openat(AT_FDCWD, "./config_file", O_RDONLY|O_CLOEXEC) = -1 ENOENT (No such file or directory)

openat(AT_FDCWD, "./config_file", O_RDONLY|O_CLOEXEC) = -1 ENOENT (No such file or directory)

openat(AT_FDCWD, "./config_file", O_RDONLY|O_CLOEXEC) = -1 ENOENT (No such file or directory)

As you can see above strace shows clearly the bug. The app is trying to open a file continuously that does not exist on the system.

In a production environment you won't have the buggy app's source code on the system to debug. You will just have a binary executable. Even if you did have the source code it may be too long for you to debug in an outage situation where time is money. Instead you can use strace to see what is happening under the hood quickly and easily.
