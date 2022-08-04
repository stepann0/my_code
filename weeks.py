import calendar, locale
locale.setlocale(locale.LC_ALL, 'ru_RU.UTF-8')

cal = calendar.Calendar()
arr = list(enumerate(cal.yeardatescalendar(2021, 1)))
arr = [(i[0], i[1][0]) for i in arr]

weeks = {"odd": [], "even": []}
c = 1
for month_numb, month in arr:
	for week_numb, week in enumerate(month):
		if month_numb > 1 and week_numb == 0 and arr[month_numb - 1][1][-1] == week:
			continue
		elif (c+1) % 2 != 0:
			weeks["odd"].append(week)
			c+=1
		else:
			weeks["even"].append(week)
			c+=1

weeks["even"].pop(0) # because weeks["even"][0] is 53th week of 2020, but isn't belong to 2021


# writing to the file
M = enumerate(calendar.month_abbr)

with open("/home/stepan/Desktop/Dev/cal.txt", "w") as cal:
	for num, name in M:
		for w in ["odd", "even"]:
			cal.write(name +  ' ' + w + "\n")
			for i in weeks[w]:
				if i[0].month == num:
					cal.write(str(i[0].day) + "-" + str(i[-1].day) + "\n")
				elif i[0].month > num:break
			cal.write("\n")
		cal.write("\n\n")
