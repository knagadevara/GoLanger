#### Go Basics

- To run a go script

        go run main.go

- To build a go script into a binary [it has to be named as main.go], ususlly the output would be the name of residing directory

        go build <path to main.go>

- Debug Conf(valid only for vscode), to be added in .bash_env/.bashrc

        alias godebug="vdlv debug --headless --listen=:2345 --log --api-version=2 -- $@"


##### Variable Decleration 

- All the variables can be declared at the package level inside var block
    - lowercase variables are private and scoped to the present package 'go-file' only.
    - uppercase are public and can be used by all the functions outside of the package if exported.
    - ':=' can only be used while declaring a new variable
        var (
            fname string
            LName string
            Age int
            dateofBirth int 
        )

- Constants can be declared on const () block

        const (
            a = iota
            b
            c
            d
        )

        const (

            var a int  = 1
            var b int8 = 2
        )

- There can be multiple blocks of 'const' and 'var'


#### Comparission

- The result of a any comparission operator is a boolean.

        num1 := 20
        num2 := 10
        answer := num1 > num2
        fmt.Println(answer) // true


#### Switch Case:

        All the case expression will be checked for a match with switch's pattern.
        If they match then statements under the case will be excecuted and no other expressions will be checked, else default statements will be executed.
        fallthrough will directly execute the statements in the next case block.


        var checker := "<pattern>"

        switch checker{
                case "zxcs":
                        statement 1
                        statement 2
                        ..
                        statement n
                case "Xyc":
                        statement 1
                        statement 2
                        ..
                        statement n
                case "zxcs":
                        fallthrough
                case "prox":
                        fallthrough
                case "inx":
                        statement 1
                        statement 2
                        ..
                        statement n
                default:
                        statement 1
                        ..
                        statement n                
        }
