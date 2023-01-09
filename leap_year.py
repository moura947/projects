#program to check if a year is a leap year

year = int(input("Enter year:"))

if year % 4 == 0:
    if year % 100 == 0:
        if year % 400 == 0:
            print("It's a leap year!")
        else:
            print("It's not a leap year.")    
    else:
        print("It's a leap year!")
else:
    print("It's not a leap year.")

