package main

/**

 * @Author: AloneAtWar

 * @Date:   2022/9/2 12:45

 * @Note:	https://leetcode.cn/problems/basic-calculator/
			224. Basic Calculator
			224. 基本计算器
 **/

//func calculate(s string) int {
//	var stack []int
//	n:=len(s)
//	curr:=1
//	res:=0
//	for i := 0; i < n; i++ {
//		if s[i]==' '{
//			continue
//		}
//		if s[i]>='0' && s[i]<='9'{
//			thisNum:=0
//			for i<n && s[i]>='0' && s[i]<='9' {
//				thisNum=thisNum*10+int(s[i]-'0')
//				i++
//			}
//			i--
//			thisNum*=curr
//			curr=1
//			res+=thisNum
//		}else if s[i]=='-'{
//			curr=-1
//		}else if s[i]=='+'{
//			curr=1
//		}else if s[i]=='('{
//			stack=append(stack,res)
//			res=0
//			stack=append(stack,curr)
//			curr=1
//		}else if s[i]==')'{
//			res*=stack[len(stack)-1]
//			res+=stack[len(stack)-2]
//			stack=stack[:len(stack)-2]
//		}
//	}
//	return res
//}
