// package 
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



/*
 * Complete the 'maximumQuality' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY packets
 *  2. INTEGER channels
 */
func getEvenAverage(a,b int32)int32{
    return (a+b)/2
}
func sorting(packets []int32) []int32{
    for i:=0; i < len(packets); i++{
        for j:=i+1; j < len(packets); j++{
            if packets[i]>packets[j]{
                packets[i],packets[j]=packets[j],packets[i]
            }
        }
    }
    return packets
}
func maximumQuality(packets []int32, channels int32) int64 {
    // Write your code here
    lens := len(packets)%2
    packets = sorting(packets)
    middle := lens/2
    results := 0
    for channels < 0 {
        if 
    }
    

}