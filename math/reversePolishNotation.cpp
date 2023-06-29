#include<iostream>
#include<cmath>
#include<stack>
#include<climits>

using namespace std;
float characterToFloat(char);
int findOperator(char );
int findOperand(char);
float prnOperat(int , int , char );
float prnEval(string postfix);

int main(){
   string post = "25*23+/";
   cout << " Result= "<<prnEval(post);
}
float characterToFloat(char ch){
   int value;
   value = ch;
   return float(value-'0');
}
int findOperator(char ch){
   if(ch == '+'|| ch == '-'|| ch == '*'|| ch == '/' || ch == '^')
      return 1;//if they found an operatAor
   return -1;//if they didn't found an operator
}
int findOperand(char ch){
   if(ch >= '0' && ch <= '9')
      return 1;//character is an operand
   return -1;//not an operand
}
float prnOperat(int x, int y, char operat){
   //Perform prnOperat
   if(operat == '+')
      return y+x;
   else if(operat == '-')
      return y-x;
   else if(operat == '*')
      return y*x;
   else if(operat == '/')
      return y/x;
   else if(operat == '^')
      return pow(y,x); //find b^a
   else
      return INT_MIN; //return negative infinity
}
float prnEval(string postfix){
   int x, y;
   stack<float> prnhandler;
   string::iterator it;
   for(it=postfix.begin(); it!=postfix.end(); it++){
      //read elements and perform postfix evaluation
      if(findOperator(*it) != -1){
         x = prnhandler.top();
         prnhandler.pop();
         y = prnhandler.top();
         prnhandler.pop();
         prnhandler.push(prnOperat(x, y, *it));
      }
      else if(findOperand(*it) > 0){
         prnhandler.push(characterToFloat(*it));
      }
   }
   return prnhandler.top();
}