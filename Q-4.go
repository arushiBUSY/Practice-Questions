package main
import (
	"fmt"
	"reflect"
)
func merge(arr1,arr2 interface{})(interface{},error){
	var m []interface{}
	if arr1==nil{

		return arr2,nil

	}else if arr2==nil{

		return arr1,nil

	}else{

		arr2_collect:=reflect.ValueOf(arr2)

		
		

		switch arr2_collect.Kind(){

		case reflect.Slice:

			if reflect.TypeOf(arr2).Elem().Kind()==reflect.Int{
				arr:=arr2_collect.Interface().([]int)
				
				for _,val:=range arr{
					m=append(m,val)
				}

			}else{

				arr:=arr2_collect.Interface().([]string)
				for _,val:=range arr{
					m=append(m,val)
				}

			}
			
		case reflect.Int:

			arr:=arr2_collect.Interface().(int)
			m=append(m,arr)

		case reflect.String:

			arr:=arr2_collect.Interface().(string)
			m=append(m,arr)

		}

		arr1_collect:=reflect.ValueOf(arr1)

		switch arr1_collect.Kind(){

		case reflect.Slice:

			if reflect.TypeOf(arr1).Elem().Kind()==reflect.Int{
				arr:=arr1_collect.Interface().([]int)
				
				for _,val:=range arr{
					m=append(m,val)
				}

			}else{

				arr:=arr1_collect.Interface().([]string)
				for _,val:=range arr{
					m=append(m,val)
				}
			}
			
		case reflect.Int:

			arr:=arr1_collect.Interface().(int)
			m=append(m,arr)

		case reflect.String:

			arr:=arr1_collect.Interface().(string)
			m=append(m,arr)

		}
	}
	var ans interface{}
	ans=m
	return ans,nil
}
func main(){
	var arr1 interface{}=[]int{1,2,3,4}
	var arr2 interface{}=[]string{"arushi","hey"}
	merged,err:=merge(arr1,arr2)
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println(merged)
	}

}
