#       找环
### 无向图找环   
1.DFS+parent node   
2.Union Find    
3.拓扑也可以，无向图改有向图，完全不建议使用    

### 有向图找环
1.使用DFS+双array标记， visited set 和 recursion array   (用于获取整个环的信息)     
2. DFS三色法拓扑
3. BFS拓扑