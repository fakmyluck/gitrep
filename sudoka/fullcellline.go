func FullCellLine(grid *[9][9][11]int,CC *bool){
    var sumyx[9][9] int
  // var sumxy[9] int
  var num1,num2 int
    
    
    
    
    for y:=0 ; y<9 ; y++{
       
       for n:=0 ; n<9 ; n++{
           
        for x:=0 ; x<9 ; x++{
                
            
               sumyx[n][x]=grid[y][x][n]
            }
    }
    num1,num2=99,99
    //proveeka na sovpodenie
       for x1:=0 ; x1<9-1 ; x1++{
           if Sumrow(sumyx[x1])==2{
               
              for x2:= x1+1 ; x2< 9;x2++{
                  if sumyx[x1]==sumyx[x2]{
                   //√ naideno
                   
                   fmt.Printf("y%v %v\n%v\n\n", y,sumyx[x1],sumyx[x2])
                   
                   for i:= 0;i<9;i++{
                       if sumyx[x1][i]==1{
                           if num1==99{
                               num1=i
                           }else{
                               num2=i
                               break
                           }
                       }
                   }
                   fmt.Print("\n",grid[y][x1],'\n',grid[y][x2],'\n')
                    for i:=0;i<9;i++{
                        
                        if i!=x1 && i!=x2{
                            grid[y][x1][i]=0
                            grid[y][x2][i]=0
                            //grid[y][X][10]=6
                        }
                        
                    }
                    
                   fmt.Print("\n",grid[y][x1],'\n',grid[y][x2],'\n')
                   fmt.Printf("num1%v  num2%v  x1%v x2%v", num1,num2 ,x1,x2)
               }
           }
       }
    }
}
}
