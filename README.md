## Quiz Game With CSV File

**This program relies on the use of flags to make the quiz decisions such as time limit of the quiz and decision whether to shuffle dependent on the user's choice.**

## Getting Started

- The go exe file should be built first
```
go build .
```

- To make the game have more than one csv files to choose from, It could be added to the same directory, and a flag could be used to make it toggle from the numerous CSV files to choose from. If ignored, the default will be chosen. The default problems.csv which is already attached.

```
go build . && ./Quiz-Game-With-CSV-File.exe -csvFile=problems.csv
``` 

- To edit the time limit, default is 25 seconds.
```
go build . && ./Quiz-Game-With-CSV-File.exe -timelimit=[input required time limit in seconds]
```

- To make the game shuffle the arrangement of the quiz questions, default is No
```
go build . && ./Quiz-Game-With-CSV-File.exe -shuffle=Yes
```