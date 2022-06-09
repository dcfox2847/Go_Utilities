# Go Utilities
This is a place to store files while I work on building out individual services for the IR toolkit.
All or most of the files in here I intend to reuse for other scripts and applications that I write moving forward.
This is really just a place to hold that data for reference later. So many tasks will need to be reused as new use cases arise, so I want to group individual functions that I create into files that can easily be imported into new projects.
This will be a long term iterative process, and very likely none of the finished products will be posted to this repository.

## filePuller and dirPuller for Log and File Collection
These files will be reused often, and were created as Go natively has no easy way to access the contents of a directory in a streamlined manner.
The capabilities at creation are only to find the data and create a copied version of the data.
It does this by taking in a string for arguments for source and destination locations on a file system, referencing the pointer information stored within a variable for the files, ensuring that all permissions are set as necessary, and creating the file.
Very early in testing as of now.
