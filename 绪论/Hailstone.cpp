#include<iostream>
/*
Hailstone(42)=(42,21,64,..,1)
*/
using namespace std;
int hailstone(int n){//计算Hailstone(n)的长度
  int length=1;
  while(1<n){
    (n%2)==0?n=3*n+1:n=n/2;
    length++;
  }
  return length;
}

int main(void){
  int n;
  n=hailstone(42);
  cout<<n<<endl;
  return 0;
}
