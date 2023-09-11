def removeChar(s:str,c:str)->str:
    if len(s)==1:
        if s[0]==c:
            return ""
        return s
    else:
        x=len(s)-1
        if s[x]==c:
            return removeChar(s[:x],c)
        return removeChar(s[:x],c)+s[x]
print(removeChar("Saurabh","a"))
                        
