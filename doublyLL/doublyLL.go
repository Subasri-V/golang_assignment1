package main

import "fmt"

// import "fmt"

type node struct{
	data int
	prev *node
	next *node
}

var head *node=nil
var tail *node=nil

func createNode(value int)(*node){
	n:=new(node)
	n.data=value
	n.prev=nil
	n.next=nil
	return n
}

func insertAtBegin(value int){
	if head==nil{
		temp:=createNode(value)
		head=temp
		tail=temp
	} else {
		temp:=createNode(value)
		temp.next=head
		head.prev=temp
		head=temp
	}
}

func insertAtlast(value int){
	if tail==nil{
		temp:=createNode(value)
		head=temp
		tail=temp
	} else {
		temp:=createNode(value)
		tail.next=temp
		temp.prev=tail
		tail=temp
	}
}

func insertAtPos(item int, NodeAfter int){
	if head!=nil{
		for temp:=head;temp!=nil;{
			if temp.data==NodeAfter{
				n:=createNode(item)
				n.prev=temp
				n.next=temp.next
				temp.next.prev=n
				temp.next=n
				temp=temp.next
			} else {
				temp=temp.next
			}
		}
	} else {
		fmt.Println("LL is empty")
	}

}

func deleteAtBegin(){
	temp:=head
	head=temp.next
	temp=nil
}

func deleteAtEnd(){
	if tail!=nil{
		temp:=tail
		tail=temp.prev
		tail.next=nil
		temp=nil
	}
	
}

func deleteANode(value int){
	if head!=nil{
		for temp:=head;temp!=nil; {
			if temp.data==value {
				temp.prev.next=temp.next
				temp.next.prev=temp.prev
				temp=nil
			} else {
				temp=temp.next
			}
		}
	}
	
}

func search(value int)(string){
	for temp:=head;temp!=nil; {
		if temp.data==value {
			return "Yes"
		} else {
			temp=temp.next
		}
	}
	return "No"
}

func update(old_value int,new_value int){
	for temp:=head;temp!=nil; {
		if temp.data==old_value {
			temp.data=new_value
		} else {
			temp=temp.next
		}
	}	
}

func length(){
	var ctr int = 0
	if head!=nil {
		ctr+=1
		for temp:=head;temp!=nil; {
			if temp.next==nil{
				break
			} else {
				ctr+=1
				temp=temp.next
			}
		}
	}
		
	fmt.Println(ctr)
}

// func reverse(){
// 	if head!=nil {
// 		temp:=head
// 		head=tail
// 		tail=temp

// 		for temp=tail;temp!=nil ; {
// 			ttemp:=temp.prev
// 			temp.prev=temp.next
// 			temp.next=ttemp
// 			temp=temp.next
// 		}
// 	}
// }

func display(){
	if head!=nil{
		for temp:=head;temp!=nil; {
			fmt.Println(temp.data)
			temp=temp.next
		}
	} else {
		fmt.Println("LL is empty")
	}
	
}

func main(){
	// insertAtBegin(10)
	// insertAtlast(20)
	// insertAtlast(30)
	// insertAtlast(40)
	// var x string= search(40)
	// fmt.Println(x)
	//x:=search(20)
	//update(20,200)
	
	//length()
	//reverse()
	//insertAtPos(15,10)
	display()

}