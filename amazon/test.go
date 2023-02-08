package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)



/*
 * Complete the 'getMaxFreqDeviation' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */
func checksameness(a,b string) bool{
    if a == b {
        return true
    }
    return false
}
func getslices(data map[string]int) (int32,int32){
    results := []int32{}
    for _,y :=range data{
        results = append(results, int32(y))
    }
    res := sorting(results)
    return res[0],res[len(data)-1]
}
func sorting( d []int32)[]int32{
    for i:=0; i<len(d); i++{
            for j:=i+1; j<len(d); j++{
                if d[i]>d[j]{
                    d[i],d[j]= d[j],d[i]
                }
            }
    }
    return d
}

func getMaxFreqDeviation(s string) int32 {
    // Write your code here
    res := make(map[string]int)
    for i:=0; i<len(s); i++{
            for j:=i+1; j<len(s); j++{
            same := checksameness(string(s[i]),string(s[j]))
            if !same {
                res[string(s[j])]= 1
            }else{
                res[string(s[j])]++
            }
        }
    }
    result1,result2 := getslices(res)
    diviation := result2 -result1
    return diviation
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := getMaxFreqDeviation(s)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
