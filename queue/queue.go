package main

import "fmt"

type node struct{
	data int
	next *node
}
var head *node=nil
var tail *node=nil

func createNode(value int)(*node){
	n:=new(node)
	n.data=value
	n.next=nil
	return n
}

func addNode(value int){
	if head==nil{
		temp:=createNode(value)
		head=temp
		tail=temp
	} else{
		temp:=createNode(value)
		tail.next=temp
		tail=temp
	}
}

func delete(){
	temp:=head
	head=temp.next
	temp=nil
}

func display(){
	for temp:=head;temp!=nil; {
		fmt.Println(temp.data)
		temp=temp.next

	}
}

func main(){
	addNode(10)
	addNode(20)
	addNode(30)
	delete()
	addNode(40)
	addNode(50)
	delete()
	addNode(60)
	display()

	//fmt.Println(head.data)

}