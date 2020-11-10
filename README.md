# jvm-in-go
It is a simple jvm implementation in go
## Chapter01
In Chapter01 I defined cmd.go to parse the command line with the help of flag and os package in go
detailed code is in ch01
+ flag package can help us to parse the command line options
- os package is used to deal with the arguments passed into our programm

## Chapter02
In Chapter02 I defined several types of class file: *.class, *.jar, *.zip, /* with Composite pattern
+ entry_composite: define a list of path which contains class files
- entry_dir: store the absolute path
+ entry_wildcard: treat the /* case
- entry_zip: deal with the jar and zip file
