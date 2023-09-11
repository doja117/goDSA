def removeString(s:str,c:str)->str:
    if len(s)<=len(c):
        if s==c:
            return ""
        return s
    
    if s[:len(c)]==c:
         return removeString(s[len(c)+1:],c)
    else:
        return s[0]+removeString(s[1:],c)


print(removeString("Saurabh Kuamr Ojha","Kumar"))
