def returnStringFromString(s:str,c:str)->str:
    if len(s)<len(c):
        return s
    t=s[:len(c)]
    print(s,t)
    if t==s:
        #print(s,t)
        return returnStringFromString(s[len(c):],c)
    print(s[0])
    return s[0]+returnStringFromString(s[1:],c)
print(returnStringFromString("Saurabh Kumar Ojha","Kuamr"))
