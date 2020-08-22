# data-processor
    
An application that processes random objects by batches of size 100 items.
The final result will be stored in a folder with the name `results`.
The link to the file with results will be printed to the STDOUT after finishing the execution of the script.

## Description of the solution

The main idea is to use the `producer-consumer` pattern.
I decided to use the plaint structure of the application and put everything in one folder because the application is really simple.
Random objects are stored in a folder `testdata`.
When a consumer has a batch of 100 items it saves these items to the result and sort by the `Seq` field. 

## How to build and run an application?

To build an application you need to execute `make build`
After that executed file will be stored in `bin` folder


To run an application you need to execute `./bin/data-processor`
After that file with the result will be stored in `results` folder and the file name will be printed to the STDOUT.

## How to run unit tests?

To run unit tests you need to execute `make test`

## Bonus question

Each time when I save a batch of 100 items I sort the result array.
We can try to use PriorityQueue and sorting will be made on the fly or we can try to use SortedSet data structure.

## Possible improvements

* We have 50 000 items, so we can create an empty array with a length of 50 000. On each iteration, we can put items in an appropriate index without using sorting.
* I covered just base case by unit tests. For a production application, I would add more unit tests.
* I would add linters to improve code quality.  







  
