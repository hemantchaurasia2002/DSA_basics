package main

import (
	"fmt"
)

//arraySize is the size of the hash table
const arraySize=7

//Hashtable hold an array
type HashTable struct{
	array [arraySize]*bucket
}

//bucket is a linkedlist in each slot of the array
type bucket struct{
	head *bucketNode
}

//bucketNode structure
type bucketNode struct{
	key string
	next *bucketNode
}

//insert will take in a key and add it to the hash table array
// func (h*HashTable) Insert(key string){
// 	index:=hash(key)
// 	h.array[index].insert(key)
// }
// //search will take in a key and return true if that key is stored in hash table
// func (h*HashTable) Insert(key string){
// 	index:=hash(key)
// }
// //delete will take in a key and delete it from the hash table
// func (h*HashTable) Insert(key string){
// 	index:=hash(key)
// }

//insert will take in a key, create a node with the key and insert the node in the bucket
func (b *bucket) insert(k string){
	if !b.search(k) {
	newNode:= &bucketNode{key: k}
	newNode.next = b.head
	b.head = newNode
	}else{
		fmt.Println(k,"already exists" )
	}
}
//search will take in a key and return true if that key is stored in hash table
func (b *bucket) search(k string) bool{
	currentNode:=b.head
	for currentNode!=nil{
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}
//delete will take in a key and delete the node from the hash table
func (b *bucket) delete(k string){
	previousNode:=b.head
	for previousNode.next!=nil{
		if previousNode.next.key == k {
			//delete
			previousNode.next = previousNode.next
		}
		previousNode = previousNode.next
	}
}

//hash
func hash(key string) int{
	sum:=0
	for _,v:= range key{
		sum+=int(v)

	}
	return sum % arraySize
}

//Init will create a bucket in each slot of the hash table
func Init() *HashTable{
	result:=&HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main(){
	// testHashTable:= Init()
	// fmt.Println(testHashTable)
	// fmt.Println(hash("RANDY"))

	testBucket := &bucket{}
	testBucket.insert("RANDY")
	testBucket.insert("RANDY")

	fmt.Println(testBucket.search("RANDY"))
	fmt.Println(testBucket.search("ERIC"))
}