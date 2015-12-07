package main

import (
        "bufio"
        "fmt"
        "os"
        "strconv"
        "strings")

type Two struct {
   in1, in2 string
}

type TwoInt struct {  // rename to useful name
    in string
    shift int
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var and_expressions = make(map[Two]string)
    var or_expressions = make(map[Two]string)
    var not_expressions = make(map[string]string)
    var shift_expressions = make(map[TwoInt]string)
    var num_and_expressions = make(map[TwoInt]string)
    var assignments = make(map[string]string)
    var inputs = make(map[string]int)
    
    for scanner.Scan() {
        var in, in1, in2, out string
        var shift int
        if strings.Contains(scanner.Text(), "AND") {
            fmt.Sscanf(scanner.Text(), "%s AND %s -> %s",&in1,&in2,&out)
            if (in1 == "1") {
                num_and_expressions[TwoInt{in2,1}] = out
            } else {
                and_expressions[Two{in1,in2}]=out
            }
        } else if strings.Contains(scanner.Text(), "OR") {
            fmt.Sscanf(scanner.Text(), "%s OR %s -> %s",&in1,&in2,&out)
            or_expressions[Two{in1,in2}]=out
        } else if strings.Contains(scanner.Text(), "NOT") {
            fmt.Sscanf(scanner.Text(), "NOT %s -> %s",&in,&out)
            not_expressions[in]=out
        } else if strings.Contains(scanner.Text(), "LSHIFT") {
            fmt.Sscanf(scanner.Text(), "%s LSHIFT %d -> %s",&in,&shift,&out)
            shift_expressions[TwoInt{in,-shift}]=out
        } else if strings.Contains(scanner.Text(), "RSHIFT") {
            fmt.Sscanf(scanner.Text(), "%s RSHIFT %d -> %s",&in,&shift,&out)
            shift_expressions[TwoInt{in,shift}] = out
        } else {
            fmt.Sscanf(scanner.Text(), "%s -> %s",&in,&out)
            input_value,err := strconv.Atoi(in)
            if err==nil {
                inputs[out] = input_value
            } else {
                assignments[in] = out
            }
            
        }
    }

    fmt.Println("not",not_expressions)
    fmt.Println("and",and_expressions)
    fmt.Println("num_and",num_and_expressions)
    fmt.Println("assignments",assignments)
    fmt.Println("or",or_expressions)
    fmt.Println("shift",shift_expressions)
    // return
    for true {
        fmt.Println(inputs)
        for inport,outport := range not_expressions {
            var known_value int
            known_value, found := inputs[inport]
            if found {
                inputs[outport] = 65535-known_value
                delete(not_expressions, inport)
            }
        }

        for shift_operator,outport := range shift_expressions {
            var known_value int
            known_value, found := inputs[shift_operator.in]
            if found {
                if (shift_operator.shift > 0) {
                    inputs[outport] = known_value >> uint(shift_operator.shift)
                } else {
                    inputs[outport] = known_value << uint(-shift_operator.shift)
                }
                
                delete(shift_expressions, shift_operator)
            }
        }

        for and_operator,outport := range and_expressions {
            var known_value1, known_value2 int
            known_value1, found1 := inputs[and_operator.in1]
            known_value2, found2 := inputs[and_operator.in2]
            if found1 && found2 {
                inputs[outport] = known_value1 & known_value2
                delete(and_expressions, and_operator)
            }
        }

        for num_and_operator,outport := range num_and_expressions {
            var known_value int
            known_value, found := inputs[num_and_operator.in]
            if found {
                inputs[outport] = known_value & 1
                delete(num_and_expressions, num_and_operator)
            }
        }

        for input,outport := range assignments {
            var known_value int
            known_value, found := inputs[input]
            if found {
                inputs[outport] = known_value
                delete(assignments, input)
            }
        }

        for or_operator,outport := range or_expressions {
            var known_value1, known_value2 int
            known_value1, found1 := inputs[or_operator.in1]
            known_value2, found2 := inputs[or_operator.in2]
            if found1 && found2 {
                inputs[outport] = known_value1 | known_value2
                delete(or_expressions, or_operator)
            }
        }

        var found_a bool
        _,found_a = inputs["a"]
        if((len(not_expressions) == 0 &&
            len(shift_expressions) == 0 &&
            len(and_expressions) == 0 &&
            len(or_expressions) == 0 &&
            len(num_and_expressions) == 0 &&
            len(assignments) == 0)) {
            fmt.Println("found a ", inputs["a"])
            return
        }
    }
}