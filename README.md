(Note: Though only library package is asked, I have checked in the whole project in GIT so that the reader may know how to use this library package very well.)

Description of the project:
Here, a library package named "action_status/action" is developed. This library has 2 functions :
	func AddAction(input string) error   
		&
	func GetStats() string

AddAction function accepts a json serialized string of the form below and maintains an average time
for each action. 3 sample inputs:
	1) {"action":"jump", "time":100}
	2) {"action":"run", "time":75}
	3) {"action":"jump", "time":200}
An end user can make concurrent calls into this function.

GetStats function accepts no input and returns a serialized json array of the average time for each action that has been provided to the addAction function. Output after the 3
sample calls above would be:
[
        {
                "action": "jump",
                "avg": 150
        },
        {
                "action": "run",
                "avg": 75
        }
]
An end user can make concurrent calls into this function as well.

How to build:
After cloning the repository, from ../../ ....../action_status directory, please run the following command
	go build

How to execute:
After successful running of 'go build' operation, please run the following command from the same directory,
        action_status
 
How to use:
In main.go file there is an example how this program can be used.

Future considerations:
	1. Storing info in database can be used to make the program more persistent.
	2. Two web APIs can be built to provide more user friendliness and to run the program anywhere from the web.

 

