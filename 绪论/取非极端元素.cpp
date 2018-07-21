#include <stdlib.h>
#include <iostream>
#include <vector>
//在数组中取得一个既不是最大也不是最小的数。

int chooseNormalElement(std::vector<int> &a){
  int length = a.size();
  if (3 > length) {
    std::cout<<"Error Length"<<std::endl;
    return false;
  }
  std::cout<<"数组长度为："<<a.size()<<std::endl;
  //随机选取三个数。
  int index = rand()%(length-2);
  int max;
  int min;
  //得到三个元素中的最大和最小。
  if (a[index]<a[index+1]){
    max = index+1;
    min = index;
    if(a[index+1]<a[index+2])
      max = index+2;
    else{
      if(a[index+2]<a[index])
	min = index+2;
    }
    
  }
  else{
    max = index;
    min = index+1;
    if(a[index]<a[index+2])
      max = index+2;
    else{
      if (a[index+2]<a[index+1])
      min = index+2;
    }
  }
  
  for (int i=index;i<=index+2;i++){
    if (i!=max&&i!=min){
      std::cout<<"非极端元素坐标为:"<<i<<std::endl;
      return a[i];
    }
  }
 
}

int main(void){
  std::vector<int> a{1,2,3};
  int result = chooseNormalElement(a);
  std:: cout << "元素为："<<result<<std::endl;
}
