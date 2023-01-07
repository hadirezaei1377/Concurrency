learning concurrency :

concurrency = Division of program execution methods into several parts for simultaneous processing by the CPU
The purpose of implementing this feature in programming is to increase the speed and efficiency of the program by dividing tasks between CPU cores, which shortens the time to reach the result ,Like high speed sites.

Concurrency means the simultaneity of the program execution time and parallelism means doing things in parallel.
Both of these are used when we have more than one task.
Tasks in the concurrency subset start at the same time, are processed at the same time, and end at the same time.
In concurrency, the CPU does not synchronize exactly the same, but performs a part of the first task, taking into account the time it stops in the first task (for example, where a file is to be opened, some time is lost.
Here it orders the hard drive to open this file for me and at the same time it goes to the second task to hit the time stop)
In this way, it does the whole work.

Also, times when the internet is weak and generally this stop is for a place that has nothing to do with CPU.
In this case, all the tasks are done and after the completion of each task, it informs that the task is finished.
But in parallelism, two tasks are done in parallel and at the same time. In this case, the CPU assigns task 1 to core 1 and task 2 to core 2.
In this case, the system must have the capability and prerequisites of this method, for example, we must have enough core for this work.

sequential mod : start task 1 , task 2  ... finish
concurrency mod : start task 1 , wait , task 2 , wait , task 1 , wait , task 2 finish conccunrently
parallel mod : start task 1 and task 2 in seperated cores and finish them , finish

